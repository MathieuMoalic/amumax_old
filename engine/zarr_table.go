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

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

type zTableTimeS struct {
	io     httpfs.WriteCloseFlusher
	buffer []byte
}

func (t *zTableTimeS) WriteToBuffer() {
	t.buffer = append(t.buffer, float64ToByte(Time)...)
}
func (t *zTableTimeS) Flush() {
	t.io.Write(t.buffer)
	var b []byte
	t.buffer = b
	t.io.Flush()
	zTableSavezay(OD()+"table/t/.zarray", 1)
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
		t.writers[i].buffer = append(t.writers[i].buffer, float64ToByte(v)...)
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
		output[i] = Float64frombytes(rawdata[count*8 : (count+1)*8])
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
	zTableSavezay(fmt.Sprintf(OD()+"table/%s/.zarray", t.name), t.q.NComp())
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func AutoFlush() {
	for {
		if zTableIsInit {
			// flush the time table
			zTableTime.Flush()
			// and all the other tables next
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

func zTableInit() {
	MakeZgroup("table")
	zTableIsInit = true
	err := httpfs.Mkdir(OD() + "table/t")
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/t/0")
	util.FatalErr(err)
	var b []byte
	zTableTime = zTableTimeS{f, b}

}

func zTableAdd(q Quantity) {
	if !zTableIsInit {
		zTableInit()
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

func zTableSavezay(path string, ncomp int) {
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
