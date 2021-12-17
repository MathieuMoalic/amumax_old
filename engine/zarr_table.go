package engine

import (
	"fmt"
	"time"

	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
	"github.com/MathieuMoalic/amumax/zarr"
)

func init() {
	DeclFunc("TableSave", zTableSave, "Save the data table right now.")
	DeclFunc("TableAdd", zTableAdd, "Save the data table periodically.")
	DeclFunc("TableAutoSave", zTableAutoSave, "Save the data table periodically.")
	go AutoFlush()
}

var zTables []zTable
var zTableAutoSavePeriod float64 = 0.0
var zTableAutoSaveStart float64 = 0.0
var zTableAutoSaveStep int = -1
var zTableFlushInterval time.Duration = 5 * time.Second
var zTableIsInit bool
var zTableTime zTableTimeS

type zTableTimeS struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
}

func (t *zTableTimeS) WriteToBuffer() {
	t.buffer = append(t.buffer, zarr.Float64ToByte(Time)...)
}
func (t *zTableTimeS) Flush() {
	t.io.Write(t.buffer)
	var b []byte
	t.buffer = b
	t.io.Flush()
	zarr.SaveFileTableZarray(OD()+"table/t/.zarray", 1, zTableAutoSaveStep)
}

type Writer struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
}

type zTable struct {
	name    string
	q       Quantity
	writers []*Writer
}

func (t *zTable) WriteToBuffer() {
	for i, v := range AverageOf(t.q) {
		t.writers[i].buffer = append(t.writers[i].buffer, zarr.Float64ToByte(v)...)
	}
}

func (t *zTable) Read() []float64 {
	rawdata, err := httpfs.Read(fmt.Sprintf(`%vtable/%s/0.0`, OD(), t.name))
	if err != nil {
		fmt.Println("<<<<<< Read error")
		return nil
	}
	var output []float64
	count := 0
	for i := 0; i < len(rawdata); i++ {
		output[i] = zarr.Float64frombytes(rawdata[count*8 : (count+1)*8])
		count++
	}
	fmt.Println(">>>>>>> table read:", output)
	return output
}

func (t *zTable) Flush() {
	for _, writer := range t.writers {
		writer.io.Write(writer.buffer)
		writer.io.Flush()
		var b []byte
		writer.buffer = b
	}
	zarr.SaveFileTableZarray(fmt.Sprintf(OD()+"table/%s/.zarray", t.name), t.q.NComp(), zTableAutoSaveStep)
}

func TableInit() {
	zarr.MakeZgroup("table", OD(), &zGroups)
	zTableIsInit = true
	err := httpfs.Mkdir(OD() + "table/t")
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/t/0")
	util.FatalErr(err)
	var b []byte
	zTableTime = zTableTimeS{f, b}

}

func AutoFlush() {
	for {
		if zTableIsInit {
			zTableTime.Flush()
			for i := range zTables {
				zTables[i].Flush()
			}
		}
		time.Sleep(zTableFlushInterval)
	}
}

func zTableSave() {
	zTableAutoSaveStep += 1
	zTableTime.WriteToBuffer()
	for i := range zTables {
		zTables[i].WriteToBuffer()
	}
}

func zTableAdd(q Quantity) {
	if !zTableIsInit {
		TableInit()
	}
	for i := range zTables {
		if zTables[i].name == NameOf(q) {
			util.Println(NameOf(q) + " was already added to the table")
			return
		}
	}

	err := httpfs.Mkdir(OD() + "table/" + NameOf(q))
	util.FatalErr(err)
	// one file = one writer = one component
	var writers []*Writer
	for comp := 0; comp < q.NComp(); comp++ {
		f, err := httpfs.Create(OD() + "table/" + NameOf(q) + "/" + fmt.Sprint(comp) + ".0")
		util.FatalErr(err)
		var b []byte
		writers = append(writers, &Writer{f, b})
	}
	z := zTable{NameOf(q), q, writers}
	zTables = append(zTables, z)
}

func zTableAutoSave(period float64) {
	zTableAutoSaveStart = Time
	zTableAutoSavePeriod = period
}
