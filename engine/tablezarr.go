package engine

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
)

func init() {
	DeclFunc("TableSave", ZarrTableSave, "Save the data table right now.")
	DeclFunc("TableAutoSave", ZarrTableAutoSave, "Save the data table periodically.")
	go AutoFlush()
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

type Writer struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
}

type ZarrTable struct {
	name     string
	q        Quantity
	writers  []*Writer
	step     int
	ncomp    int
	autosave bool
	period   float64
	start    float64
}

var ZarrTables []ZarrTable

func (t *ZarrTable) WriteToBuffer() {
	// fmt.Println("Printing to buffer")
	for i, v := range AverageOf(t.q) {
		t.writers[i].buffer = append(t.writers[i].buffer, float64ToByte(v)...)
	}
	t.step += 1
	t.Flush()
	// fmt.Println("tstep:", t.step)
}
func (t *ZarrTable) needSave() bool {
	return t.period != 0 && (Time-t.start)-float64(t.step)*t.period >= t.period && t.autosave
}
func (t *ZarrTable) Flush() {
	for _, writer := range t.writers {
		writer.io.Write(writer.buffer)
		writer.io.Flush()
		var b []byte
		writer.buffer = b
	}
	ZarrTableSaveZarray(fmt.Sprintf(OD()+"table/%s/.zarray", t.name), t.step+1, t.ncomp)
}

// return existing table else create new one
func getTable(q Quantity, autosave bool, period float64) *ZarrTable {
	for i, v := range ZarrTables {
		if v.name == NameOf(q) {
			return &ZarrTables[i]
		}
	}
	// if not in the list of tables, then create it
	MakeZgroup("table")
	err := httpfs.Mkdir(OD() + "table/" + NameOf(q))
	util.FatalErr(err)
	// one file = one writer for each comp
	var writers []*Writer
	for comp := 0; comp < q.NComp(); comp++ {
		f, err := httpfs.Create(OD() + "table/" + NameOf(q) + "/" + fmt.Sprint(comp) + ".0")
		util.FatalErr(err)
		var b []byte
		writers = append(writers, &Writer{f, b})
	}
	z := ZarrTable{NameOf(q), q, writers, -1, q.NComp(), autosave, period, Time}
	ZarrTables = append(ZarrTables, z)
	return &z
}

func AutoFlush() {
	for {
		for _, v := range ZarrTables {
			v.Flush()
		}
		time.Sleep(5 * time.Second)
	}
}

func ZarrTableSave(q Quantity) {
	t := getTable(q, false, 0.0)
	t.WriteToBuffer()
}
func ZarrTableAutoSave(q Quantity, period float64) {
	getTable(q, true, period)
	// t.WriteToBuffer()
}

func ZarrTableSaveZarray(path string, nbtime int, ncomp int) {
	var zarray_template = `{
	"chunks": [1,%d],
	"compressor": null,
	"dtype": "<f8",
	"fill_value": 0.0,
	"filters": null,
	"order": "C",
	"shape": [%d,%d],
	"zarr_format": 2
}`
	fzarray, err := httpfs.Create(path)
	util.FatalErr(err)
	defer fzarray.Close()
	metadata := fmt.Sprintf(zarray_template, nbtime, ncomp, nbtime)
	fzarray.Write([]byte(metadata))
}
