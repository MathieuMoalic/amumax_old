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
	DeclFunc("TableSave", ZarrTableSave, "Save the data table right now (appends one line).")
	go AutoFlush()
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

type ZarrTable struct {
	name    string
	q       Quantity
	writers []httpfs.WriteCloseFlusher
	step    int
	ncomp   int
}

var ZarrTables []ZarrTable
var ZarrTableInits []string

func (t *ZarrTable) WriteToBuffer() {
	vec := AverageOf(t.q)
	for i, v := range vec {
		// fprint(t.writers[i], float64ToByte(v))
		t.writers[i].Write(float64ToByte(v))
	}
	t.step += 1
	ZarrTableSaveZarray(fmt.Sprintf(OD()+"%s/.zarray", t.name), t.step+1, t.ncomp)
}
func (t *ZarrTable) Flush() {
	for _, writer := range t.writers {
		writer.Flush()
	}
}

func getTable(q Quantity) ZarrTable {
	for _, v := range ZarrTables {
		if v.q == q {
			return v
		}
	}
	// if not in the list of tables, then create it
	err := httpfs.Mkdir(OD() + NameOf(q))
	util.FatalErr(err)
	// one file = one writer for each comp
	var fs []httpfs.WriteCloseFlusher
	for comp := 0; comp < q.NComp(); comp++ {
		f, err := httpfs.Create(OD() + NameOf(q) + "/" + fmt.Sprint(comp) + ".0")
		util.FatalErr(err)
		fs = append(fs, f)
	}
	z := ZarrTable{NameOf(q), q, fs, 0, q.NComp()}
	ZarrTables = append(ZarrTables, z)
	return z
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
	t := getTable(q)
	t.WriteToBuffer()
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
