package zarr

import (
	// "bufio"
	"encoding/binary"
    // "fmt"
    "math"
	"github.com/mumax/3/data"
	// "github.com/mumax/3/util"
	"github.com/mumax/3/httpfs"
	"io/ioutil"
	"path"
	"encoding/json"
	"github.com/DataDog/zstd" 
	// "reflect"
)

func Float32frombytes(bytes []byte) float32 {
    bits := binary.LittleEndian.Uint32(bytes)
    float := math.Float32frombits(bits)
    return float
}

type Zarray struct {
	Chunks     [5]int
	Compressor string
	Dtype      string
	FillValue  string
	Filters    string
	Order      string
	Shape      string
	ZarrFormat string
}

func Read(fname string) (s *data.Slice, err error) {
	basedir := path.Dir(fname)
	content, _ := ioutil.ReadFile(basedir+"/.zarray")
	var zarray Zarray
	json.Unmarshal([]byte(content), &zarray)

	// fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	// fmt.Println(zarray)
	// fmt.Println(zarray.Chunks)
	// fmt.Println(zarray.Chunks[3])
	// fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	sizez := zarray.Chunks[1]
	sizey := zarray.Chunks[2]
	sizex := zarray.Chunks[3]
	sizec := zarray.Chunks[4]

	var size = [3]int{sizex,sizey,sizez}
	array := data.NewSlice(sizec, size)
	tensors := array.Tensors()
	ncomp := array.NComp()
	io_reader,err := httpfs.Open(fname)
	if err != nil {
		panic(err)
	}
	compressedData, err := ioutil.ReadAll(io_reader)
	if err != nil {
		panic(err)
	} 
	data,err := zstd.Decompress(nil, compressedData)
	if err != nil {
		panic(err)
	}
	count := 0
	for iz := 0; iz < size[2]; iz++ {
		for iy := 0; iy < size[1]; iy++ {
			for ix := 0; ix < size[1]; ix++ {
				for c := 0; c < ncomp; c++ {
					tensors[c][iz][iy][ix] = Float32frombytes(data[count*4:(count+1)*4])
					count++
				}
			}
		}
	}
	
	return array,nil
}
