package engine

import (
	"time"

	"github.com/MathieuMoalic/amumax/httpfs"
	"github.com/MathieuMoalic/amumax/util"
	"github.com/MathieuMoalic/amumax/zarr"
)

func init() {
	DeclFunc("TableSave", ZTableSave, "Save the data table right now.")
	DeclFunc("TableAdd", ZTableAdd, "Save the data table periodically.")
	DeclFunc("TableAutoSave", ZTableAutoSave, "Save the data table periodically.")
	ZTables = ZTablesStruct{Step: -1, AutoSavePeriod: 0.0, FlushInterval: 5 * time.Second}
}

var ZTables ZTablesStruct

type ZTablesStruct struct {
	Tables         []ZTable `json:"Tables"`
	Qs             []Quantity
	AutoSavePeriod float64       `json:"AutoSavePeriod"`
	AutoSaveStart  float64       `json:"AutoSaveStart"`
	Step           int           `json:"Step"`
	FlushInterval  time.Duration `json:"FlushInterval"`
}

type ZTable struct {
	Name   string    `json:"Name"`
	Data   []float64 `json:"Data"`
	buffer []byte
	io     httpfs.WriteCloseFlusher
}

func (ts *ZTablesStruct) WriteToBuffer() {
	buf := []float64{}
	// always save the current time
	buf = append(buf, Time)
	// for each quantity we append each component to the buffer
	for _, q := range ts.Qs {
		buf = append(buf, AverageOf(q)...)
	}
	// size of buf should be same as size of []Ztable
	for i, b := range buf {
		ts.Tables[i].buffer = append(ts.Tables[i].buffer, zarr.Float64ToByte(b)...)
		ts.Tables[i].Data = append(ts.Tables[i].Data, b)
	}
}
func (ts *ZTablesStruct) Flush() {
	for i := range ts.Tables {
		ts.Tables[i].io.Write(ts.Tables[i].buffer)
		ts.Tables[i].buffer = []byte{}
		ts.Tables[i].io.Flush()
		zarr.SaveFileTableZarray(OD()+"table/"+ts.Tables[i].Name+"/.zarray", ts.Step)
	}
}
func (ts *ZTablesStruct) NeedSave() bool {
	return ts.AutoSavePeriod != 0 && (Time-ts.AutoSaveStart)-float64(ts.Step)*ts.AutoSavePeriod >= ts.AutoSavePeriod
}

func TableInit() {
	zarr.MakeZgroup("table", OD(), &zGroups)
	err := httpfs.Mkdir(OD() + "table/t")
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/t/0")
	util.FatalErr(err)
	ZTables.Tables = append(ZTables.Tables, ZTable{"t", []float64{}, []byte{}, f})
	go AutoFlush()

}

func AutoFlush() {
	for {
		ZTables.Flush()
		time.Sleep(ZTables.FlushInterval)
	}
}

func ZTableSave() {
	ZTables.Step += 1
	ZTables.WriteToBuffer()
}

func CreateTable(name string) ZTable {
	err := httpfs.Mkdir(OD() + "table/" + name)
	util.FatalErr(err)
	f, err := httpfs.Create(OD() + "table/" + name + "/0")
	util.FatalErr(err)
	return ZTable{Name: name, Data: []float64{}, buffer: []byte{}, io: f}
}

func ZTableAdd(q Quantity) {
	if len(ZTables.Tables) == 0 {
		TableInit()
	}
	if ZTables.Step != -1 {
		util.Fatal("Add Table Quantity BEFORE you save the table for the first time")
	}
	ZTables.Qs = append(ZTables.Qs, q)
	if q.NComp() == 1 {
		ZTables.Tables = append(ZTables.Tables, CreateTable(NameOf(q)))
	} else {
		suffixes := []string{"x", "y", "z"}
		for comp := 0; comp < q.NComp(); comp++ {
			ZTables.Tables = append(ZTables.Tables, CreateTable(NameOf(q)+suffixes[comp]))
		}
	}
}

func ZTableAutoSave(period float64) {
	ZTables.AutoSaveStart = Time
	ZTables.AutoSavePeriod = period
}
