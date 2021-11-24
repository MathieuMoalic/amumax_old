package engine

// Bookkeeping for auto-saving quantities at given intervals.

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/DataDog/zstd" 
	
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/httpfs"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

func init() {
	DeclFunc("ZarrAutoSave", ZarrAutoSave, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("ZarrSaveAs", ZarrSaveOnceAs, "Save space-dependent quantity as the zarr standard.")
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
var zattrs_template = `{
    "dx": %e,
    "dy": %e,
    "dz": %e,
    "Nx": %d,
    "Ny": %d,
    "Nz": %d,
    "Tx": %e,
    "Ty": %e,
    "Tz": %e,
    "PBC": [%d,%d,%d]
}`

var zarr_autonum = make(map[string]int) 

// Register quant to be auto-saved every period.
// period == 0 stops autosaving.
func ZarrAutoSave(q Quantity, period float64) {
	// this block shouldn't be here but somehwere in init but w/e
	zgroup, err := httpfs.Create(OD() + ".zgroup")
	util.FatalErr(err)
	defer zgroup.Close()
	zgroup.Write([]byte("{\"zarr_format\": 2}"))

	// init the dataset with a folder
	os.MkdirAll(OD() + NameOf(q), 0700)
	// dir, err := httpfs.Create(OD() + NameOf(q)+"/")
	// util.FatalErr(err)
	// defer dir.Close()

	SaveAttrs(q)

	if period == 0 {
		delete(output, q)
		} else {
			output[q] = &autosave{period, Time, -1, ZarrSave} // init count to -1 allows save at t=0
		}
}

func ZarrSaveOnceAs(q Quantity, fname string){
	os.MkdirAll(OD() + fname, 0700)
	buffer := ValueOf(q)
	defer cuda.Recycle(buffer)
	data := buffer.HostCopy() // must be copy (async io)
	queOutput(func() { ZarrSyncSave(data,fname,-1) })
}

func SaveAttrs(q Quantity) {
	// metadata stored in .zattrs
	zattrs, err := httpfs.Create(OD() +"/.zattrs")
	util.FatalErr(err)
	defer zattrs.Close()
	cellsize := MeshOf(q).CellSize()
	gridSize := MeshOf(q).Size()
	pbc := MeshOf(q).PBC()
	metadata := fmt.Sprintf(zattrs_template,
		cellsize[X],
		cellsize[Y],
		cellsize[Z],
		gridSize[X],
		gridSize[Y],
		gridSize[Z],
		cellsize[X] * float64(gridSize[X]),
		cellsize[Y] * float64(gridSize[Y]),
		cellsize[Z] * float64(gridSize[Z]),
		pbc[0],
		pbc[1],
		pbc[2],
	)
	zattrs.Write([]byte(metadata))

}
// Save once, with auto file name
func ZarrSave(q Quantity) {
	qname := NameOf(q)
	
	buffer := ValueOf(q)
	defer cuda.Recycle(buffer)
	data := buffer.HostCopy() // must be copy (async io)
	t := zarr_autonum[qname] // this way we hardcopy t and pass it to the save func
	queOutput(func() { ZarrSyncSave(data,qname,t) })
	zarr_autonum[qname]++
}

// synchronous save
func ZarrSyncSave(array *data.Slice, qname string, time int) {
	var ftime int
	if time == -1 {
		ftime = 0
	} else {
		ftime = time
	}
	f, err := httpfs.Create(fmt.Sprintf(OD()+"%s/%d.0.0.0.0", qname, ftime))
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
	// CompressLevel(dst, src []byte, level int) // alternative with compress levels
	compressedData,err := zstd.Compress(nil, bdata)
	util.FatalErr(err)
	f.Write(compressedData)

	//.zarray file
	fzarray, err := httpfs.Create(fmt.Sprintf(OD()+"%s/.zarray", qname))
	util.FatalErr(err)
	defer fzarray.Close()

	if time == -1 {
		ftime = 1
	} else {
		ftime = time
	}
	metadata := fmt.Sprintf(zarray_template,size[Z],size[Y],size[X],ncomp,ftime,size[Z],size[Y],size[X],ncomp)
	fzarray.Write([]byte(metadata))
}
