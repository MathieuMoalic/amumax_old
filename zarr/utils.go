package zarr

import (
	"fmt"
	"time"

	"github.com/MathieuMoalic/amumax/data"
	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
)

// type ZTime struct {
// 	name   string
// 	buffer []float64
// 	writer httpfs.WriteCloseFlusher
// 	init   bool
// }
type Zattrs struct {
	Buffer []float64
}

func SaveMetaStart(fname string, m data.Mesh, dt float64, start_time time.Time) {
	SaveMeta(fname, m, dt, start_time.Format(time.UnixDate), "", "")
}
func SaveMetaEnd(fname string, m data.Mesh, dt float64, start_time time.Time) {
	SaveMeta(fname, m, dt, start_time.Format(time.UnixDate), time.Now().Format(time.UnixDate), fmt.Sprint(time.Since(start_time)))
}

func SaveMeta(fname string, m data.Mesh, dt float64, start_time string, end_time string, diff_time string) {
	var zattrs_template = `{
	"dt": %e,
	"dx": %e,
	"dy": %e,
	"dz": %e,
	"Nx": %d,
	"Ny": %d,
	"Nz": %d,
	"Tx": %e,
	"Ty": %e,
	"Tz": %e,
	"PBC": [
		%d,
		%d,
		%d
	],
	"Start_time": "%s",
	"End_time": "%s",
	"Total_time": "%s"
}`
	//.zattrs file
	zattrs, err := httpfs.Create(fname)
	util.FatalErr(err)
	defer zattrs.Close()
	D := m.CellSize()
	S := m.Size()
	P := m.PBC()
	metadata := fmt.Sprintf(zattrs_template, dt, D[0], D[1], D[2], S[0], S[1], S[2], float64(S[0])*D[0], float64(S[1])*D[1], float64(S[2])*D[2], P[0], P[1], P[2], start_time, end_time, diff_time)
	zattrs.Write([]byte(metadata))
}

func MakeZgroup(name string, od string, zGroups *[]string) {
	exists := false
	for _, v := range *zGroups {
		if name == v {
			exists = true
			*zGroups = append(*zGroups, name)
		}
	}
	if !exists {
		err := httpfs.Mkdir(od + name)
		util.FatalErr(err)
		InitZgroup(od + name + "/")
		*zGroups = append(*zGroups, name)
	}

}

func InitZgroup(path string) {
	zgroup, err := httpfs.Create(path + ".zgroup")
	util.FatalErr(err)
	defer zgroup.Close()
	zgroup.Write([]byte("{\"zarr_format\": 2}"))
}

func SaveFileZarray(path string, size [3]int, ncomp int, time int) {
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
	metadata := fmt.Sprintf(zarray_template, size[2], size[1], size[0], ncomp, time+1, size[2], size[1], size[0], ncomp)
	fzarray.Write([]byte(metadata))
}

func SaveFileTableZarray(path string, ncomp int, zTableAutoSaveStep int) {
	var zarray_template = `{
	"chunks": [%s%d],
	"compressor": null,
	"dtype": "<f8",
	"fill_value": 0.0,
	"filters": null,
	"order": "C",
	"shape": [%s%d],
	"zarr_format": 2
}`
	var s1 string
	var s2 string
	if ncomp == 1 {
		s1 = ""
		s2 = ""
	} else {
		s1 = "1,"
		s2 = fmt.Sprintf("%d,", ncomp)
	}
	fzarray, err := httpfs.Create(path)
	util.FatalErr(err)
	defer fzarray.Close()
	metadata := fmt.Sprintf(zarray_template, s1, zTableAutoSaveStep+1, s2, zTableAutoSaveStep+1)
	fzarray.Write([]byte(metadata))
}
