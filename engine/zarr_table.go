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
}

var zTables = make(map[string]*zTable)
var zTableAutoSavePeriod float64 = 0.0
var zTableAutoSaveStart float64 = 0.0
var zTableStep int = -1
var zTableFlushInterval time.Duration = 5 * time.Second
var zTableTime zTableTimeS

type zTableTimeS struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
	data   []float64
}

func (t *zTableTimeS) WriteToBuffer() {
	t.buffer = append(t.buffer, zarr.Float64ToByte(Time)...)
	t.data = append(t.data, Time)
}
func (t *zTableTimeS) Flush() {
	t.io.Write(t.buffer)
	t.buffer = []byte{}
	t.io.Flush()
	zarr.SaveFileTableZarray(OD()+"table/t/.zarray", 1, zTableStep)
}

type Writer struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
}

type zTable struct {
	name    string
	q       Quantity
	writers []*Writer
	data    [][]float64
}

func (t *zTable) WriteToBuffer() {
	for i, v := range AverageOf(t.q) {
		t.writers[i].buffer = append(t.writers[i].buffer, zarr.Float64ToByte(v)...)
		t.data[i] = append(t.data[i], v)
	}
}

func (t *zTable) Flush() {
	for _, writer := range t.writers {
		writer.io.Write(writer.buffer)
		writer.io.Flush()
		writer.buffer = []byte{}
	}
	zarr.SaveFileTableZarray(fmt.Sprintf(OD()+"table/%s/.zarray", t.name), t.q.NComp(), zTableStep)
}

func TableInit() {
	zarr.MakeZgroup("table", OD(), &zGroups)
	err := httpfs.Mkdir(OD() + "table/t")
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/t/0")
	util.FatalErr(err)
	zTableTime = zTableTimeS{f, []byte{}, []float64{}}
	go AutoFlush()

}

func AutoFlush() {
	for {
		zTableTime.Flush()
		for i := range zTables {
			zTables[i].Flush()
		}
		time.Sleep(zTableFlushInterval)
	}
}

func zTableSave() {
	zTableStep += 1
	zTableTime.WriteToBuffer()
	for _, t := range zTables {
		t.WriteToBuffer()
	}
}

func zTableAdd(q Quantity) {
	if len(zTables) == 0 {
		TableInit()
	}
	if _, exists := zTables[NameOf(q)]; exists {
		util.Println(NameOf(q) + " was already added to the table")
		return
	}
	if zTableStep != -1 {
		util.Fatal("Add Table Quantity BEFORE you save the table for the first time")
	}

	err := httpfs.Mkdir(OD() + "table/" + NameOf(q))
	util.FatalErr(err)
	// one file = one writer = one component
	var writers []*Writer
	var data [][]float64
	for comp := 0; comp < q.NComp(); comp++ {
		f, err := httpfs.Create(OD() + "table/" + NameOf(q) + "/" + fmt.Sprint(comp) + ".0")
		util.FatalErr(err)
		writers = append(writers, &Writer{f, []byte{}})
		data = append(data, []float64{})
	}
	fmt.Println(">>>>>>>", data)
	zTables[NameOf(q)] = &zTable{NameOf(q), q, writers, data}
}

func zTableAutoSave(period float64) {
	zTableAutoSaveStart = Time
	zTableAutoSavePeriod = period
}
