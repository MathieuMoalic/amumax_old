package engine

// Bookkeeping for auto-saving quantities at given intervals.

import (
	"fmt"
	"os"
	"io"
	"unsafe"

	"github.com/DataDog/zstd" 
	
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/httpfs"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

func init() {
	DeclFunc("AutoSaveZarr", AutoSaveZarr, "Auto save space-dependent quantity every period (s) as the zarr standard.")
}

var zarray_template = `{
    "chunks": [1,%d,%d,%d,%d],
    "compressor": {"id": "zstd","level": 1},
    "dtype": "<f4",
    "fill_value": 0.0,
    "filters": null,
    "order": "C",
    "shape": [%d,%d,%d,%d,%d],
    "zarr_format": 2
}`

var zarr_autonum = make(map[string]int) 

// Register quant to be auto-saved every period.
// period == 0 stops autosaving.
func AutoSaveZarr(q Quantity, period float64) {
	zgroup, _ := os.Create(OD() + ".zgroup")
	defer zgroup.Close()
	io.WriteString(zgroup, "{\"zarr_format\": 2}")

	// init the dataset with a folder
	os.MkdirAll(OD() + NameOf(q), 0700)

	if period == 0 {
		delete(output, q)
	} else {
		output[q] = &autosave{period, Time, -1, zarr_save} // init count to -1 allows save at t=0
	}
}

// Save once, with auto file name
func zarr_save(q Quantity) {
	qname := NameOf(q)
	
	buffer := ValueOf(q) // TODO: check and optimize for Buffer()
	defer cuda.Recycle(buffer)
	data := buffer.HostCopy() // must be copy (async io)
	queOutput(func() { zarr_sync_save(data,qname) })
	zarr_autonum[qname]++
}

// synchronous save
func zarr_sync_save(array *data.Slice, qname string) {
	f, err := httpfs.Create(fmt.Sprintf(OD()+"%s/%d.0.0.0.0", qname, zarr_autonum[qname]-1))
	util.FatalErr(err)
	defer f.Close()

	data := array.Tensors()
	size := array.Size()
	
	var bdata []byte
	var bytes []byte
	
	ncomp := array.NComp()
	for iz := 0; iz < size[Z]; iz++ {
		for iy := 0; iy < size[Y]; iy++ {
			for ix := 0; ix < size[X]; ix++ {
				for c := 0; c < ncomp; c++ {
					bytes = (*[4]byte)(unsafe.Pointer(&data[c][iz][iy][ix]))[:]
					for k := 0; k < 4; k++{
						bdata = append(bdata,bytes[k])
					}
				}
			}
		}
	}
	compressedData,_ := zstd.Compress(nil, bdata)
	f.Write(compressedData)

	//.zarray file
	fzarray, err := httpfs.Create(fmt.Sprintf(OD()+"%s/.zarray", qname))
	util.FatalErr(err)
	defer fzarray.Close()
	metadata := fmt.Sprintf(zarray_template,size[Z],size[Y],size[X],ncomp,zarr_autonum[qname],size[Z],size[Y],size[X],ncomp)
	fzarray.Write([]byte(metadata))
}