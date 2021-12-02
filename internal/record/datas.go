package record

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//TypesOfData
// Datas both the SAUCE DataType value and name.
type Datas struct {
	Type TypeOfData `json:"type" xml:"type"`
	Name string     `json:"name" xml:"name"`
}

// DataType is the data type (SAUCE DataType).
type TypeOfData uint

const (
	Nones TypeOfData = iota
	Characters
	Bitmaps
	Vectors
	Audios
	BinaryTexts
	XBins
	Archives
	Executables
)

func (d TypeOfData) String() string {
	if d > Executables {
		return ""
	}
	return [...]string{
		"undefined",
		"text or character stream",
		"bitmap graphic or animation",
		"vector graphic",
		"audio or music",
		"binary text",
		"extended binary text",
		"archive",
		"executable",
	}[d]
}

func (d *Data) DataType() Datas {
	dt := unsignedBinary1(d.Datatype)
	return Datas{
		Type: TypeOfData(dt),
		Name: TypeOfData(dt).String(),
	}
}

func unsignedBinary1(b [1]byte) (value uint8) {
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println("unsigned 1 byte, LE binary failed:", err)
	}
	return value
}
