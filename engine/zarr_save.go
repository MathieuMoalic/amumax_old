package engine

// Bookkeeping for auto-saving quantities at given intervals.

import (
	"fmt"
	"unsafe"

	"github.com/DataDog/zstd"

	"github.com/MathieuMoalic/amumax/cuda"
	"github.com/MathieuMoalic/amumax/data"
	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
)

func init() {
	DeclFunc("AutoSaveAs", ZarrAutoSaveAs, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("AutoSave", ZarrAutoSave, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("SaveAs", ZarrSaveAs, "Save space-dependent quantity as the zarr standard.")
	DeclFunc("Save", ZarrSave, "Save space-dependent quantity as the zarr standard.")
}

var zarr_autonum = make(map[string]int)
var zarr_output = make(map[Quantity]*zarr_autosave) // when to save quantities
var ZarrGroups []string

type zarr_autosave struct {
	period float64 // How often to save
	start  float64 // Starting point
	count  int     // Number of times it has been autosaved
	name   string
	save   func(Quantity, string) // called to do the actual save
}

// returns true when the time is right to save.
func (a *zarr_autosave) needSave() bool {
	t := Time - a.start
	return a.period != 0 && t-float64(a.count)*a.period >= a.period
}
func MakeZgroup(name string) {
	exists := false
	for _, v := range ZarrGroups {
		if name == v {
			exists = true
			ZarrGroups = append(ZarrGroups, name)
		}
	}
	if !exists {
		err := httpfs.Mkdir(OD() + name)
		util.FatalErr(err)
		InitZgroup(name + "/")
		ZarrGroups = append(ZarrGroups, name)
	}

}

func InitZgroup(name string) {
	zgroup, err := httpfs.Create(OD() + name + ".zgroup")
	util.FatalErr(err)
	defer zgroup.Close()
	zgroup.Write([]byte("{\"zarr_format\": 2}"))
}

// Register quant to be auto-saved every period.
// period == 0 stops autosaving.
func ZarrAutoSave(q Quantity, period float64) {
	ZarrAutoSaveAs(q, NameOf(q), period)
}

func ZarrAutoSaveAs(q Quantity, fname string, period float64) {
	if period == 0 {
		delete(zarr_output, q)
	} else {
		zarr_output[q] = &zarr_autosave{period, Time, -1, fname, ZarrSaveAs} // init count to -1 allows save at t=0
	}
}

func ZarrSaveAs(q Quantity, fname string) {
	httpfs.Mkdir(OD() + fname)

	buffer := ValueOf(q)
	defer cuda.Recycle(buffer)
	data := buffer.HostCopy() // must be copy (async io)
	t := zarr_autonum[fname]  // no desync this way
	queOutput(func() { ZarrSyncSave(data, fname, t) })
	zarr_autonum[fname]++
}

// Save once, with auto file name
func ZarrSave(q Quantity) {
	ZarrSaveAs(q, NameOf(q))
}

func ZarrSaveZarray(path string, size [3]int, ncomp int, time int) {
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
	fzarray, err := httpfs.Create(path)
	util.FatalErr(err)
	defer fzarray.Close()
	metadata := fmt.Sprintf(zarray_template, size[Z], size[Y], size[X], ncomp, time+1, size[Z], size[Y], size[X], ncomp)
	fzarray.Write([]byte(metadata))
}

// synchronous save
func ZarrSyncSave(array *data.Slice, qname string, time int) {
	f, err := httpfs.Create(fmt.Sprintf(OD()+"%s/%d.0.0.0.0", qname, time))
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
					for k := 0; k < 4; k++ {
						bdata = append(bdata, bytes[k])
					}
				}
			}
		}
	}
	// CompressLevel(dst, src []byte, level int) // alternative with compress levels
	compressedData, err := zstd.Compress(nil, bdata)
	util.FatalErr(err)
	f.Write(compressedData)

	//.zarray file
	ZarrSaveZarray(fmt.Sprintf(OD()+"%s/.zarray", qname), size, ncomp, time)
}
