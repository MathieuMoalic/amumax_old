package zarr

import (
	"fmt"
	"time"

	"github.com/MathieuMoalic/amumax/data"
	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
)

func SaveMetaStart(fname string, m data.Mesh, start_time time.Time) {
	SaveMeta(fname, m, start_time.Format(time.UnixDate), "", "")
}
func SaveMetaEnd(fname string, m data.Mesh, start_time time.Time) {
	SaveMeta(fname, m, start_time.Format(time.UnixDate), time.Now().Format(time.UnixDate), fmt.Sprintf("%s", time.Now().Sub(start_time)))
}

func SaveMeta(fname string, m data.Mesh, start_time string, end_time string, diff_time string) {
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
	metadata := fmt.Sprintf(zattrs_template, D[0], D[1], D[2], S[0], S[1], S[2], float64(S[0])*D[0], float64(S[1])*D[1], float64(S[2])*D[2], P[0], P[1], P[2], start_time, end_time, diff_time)
	zattrs.Write([]byte(metadata))
}
