package zarr

import (
	"fmt"
	"encoding/json"
	"github.com/mumax/3/util"
	"github.com/mumax/3/httpfs"
	"github.com/mumax/3/data"
	"time"
)

type Zmeta struct {
	dx		float64
	dy 		float64
	dz 		float64
	Nx  	int
	Ny    	int
	Nz      int
	Tx      float64
	Ty 		float64
	Tz 		float64
	PBC 	[3]int
	Start_time string
	End_time string
	Total_time string
}

func (a *Zmeta) Save(fname string) {
	zattrs, err := httpfs.Create(fname)
	util.FatalErr(err)
	defer zattrs.Close()
	metadata, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	_,err = zattrs.Write([]byte(metadata))
	util.FatalErr(err)
}

func (a *Zmeta) EndSave(fname string, start_time time.Time){
	EndTime := time.Now()
	a.End_time = EndTime.Format(time.UnixDate)
	a.Total_time = fmt.Sprintf("%s",EndTime.Sub(start_time))
	a.Save(fname) // we save once at the end
}

func (a *Zmeta) StartSave(fname string, m data.Mesh){
	D := m.CellSize()
	S := m.Size()
	P := m.PBC()
	a.dx = D[0]
	a.dy = D[1]
	a.dz = D[2]
	a.Nx = S[0]
	a.Ny = S[1]
	a.Nz = S[2]
	a.Tx = float64(a.Nx) * a.dx
	a.Ty = float64(a.Ny) * a.dy
	a.Tz = float64(a.Nz) * a.dz
	a.PBC = [3]int{P[0], P[1], P[2]}
	a.Start_time = time.Now().Format(time.UnixDate)
	a.Save(fname) // we save once at the end
}