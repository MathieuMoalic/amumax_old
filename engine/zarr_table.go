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
	DeclFunc("TableAdd", ZarrTableAdd, "Save the data table periodically.")
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
	name    string
	q       Quantity
	writers []*Writer
}

var ZarrTables []ZarrTable
var ZarrTableAutoSavePeriod float64 = 0.0
var ZarrTableAutoSaveStart float64 = 0.0
var ZarrTableAutoSaveStep int = -1
var ZarrTableFlushInterval time.Duration = 5 * time.Second
var ZarrTableIsInit bool
var ZarrTableTime Writer

func (t *ZarrTable) WriteToBuffer() {
	for i, v := range AverageOf(t.q) {
		t.writers[i].buffer = append(t.writers[i].buffer, float64ToByte(v)...)
	}
}

func (t *ZarrTable) Read() []float64 {
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

func (t *ZarrTable) Flush() {
	for _, writer := range t.writers {
		writer.io.Write(writer.buffer)
		writer.io.Flush()
		var b []byte
		writer.buffer = b
	}
	ZarrTableSaveZarray(fmt.Sprintf(OD()+"table/%s/.zarray", t.name), t.q.NComp())
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func AutoFlush() {
	for {
		if ZarrTableIsInit {
			// flush the time table
			ZarrTableTime.io.Write(ZarrTableTime.buffer)
			ZarrTableTime.io.Flush()
			// and all the other tables next
			for i := range ZarrTables {
				ZarrTables[i].Flush()
			}
		}
		time.Sleep(ZarrTableFlushInterval)
	}
}

func ZarrTableSave() {
	ZarrTableAutoSaveStep += 1
	println(Time)
	for _, v := range float64ToByte(Time) {
		print(v)
	}
	println("")
	ZarrTableTime.buffer = append(ZarrTableTime.buffer, float64ToByte(Time)...)
	for i := range ZarrTables {
		ZarrTables[i].WriteToBuffer()
	}
}

func ZarrTableInit() {
	MakeZgroup("table")
	ZarrTableIsInit = true
	err := httpfs.Mkdir(OD() + "table/t")
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/t/0")
	util.FatalErr(err)
	var b []byte
	ZarrTableTime = Writer{f, b}

}

func ZarrTableAdd(q Quantity) {
	if !ZarrTableIsInit {
		ZarrTableInit()
	}
	for i := range ZarrTables {
		if ZarrTables[i].name == NameOf(q) {
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
	z := ZarrTable{NameOf(q), q, writers}
	ZarrTables = append(ZarrTables, z)
}

func ZarrTableAutoSave(period float64) {
	ZarrTableAutoSaveStart = Time
	ZarrTableAutoSavePeriod = period
}

func ZarrTableSaveZarray(path string, ncomp int) {
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
	metadata := fmt.Sprintf(zarray_template, ZarrTableAutoSaveStep+1, ncomp, ZarrTableAutoSaveStep+1)
	fzarray.Write([]byte(metadata))
}
