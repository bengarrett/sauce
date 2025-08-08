package layout

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/bengarrett/sauce/humanize"
	"golang.org/x/text/language"
)

// Sizes is the original file size in multiple formats.
type Sizes struct {
	Bytes   uint16 `json:"bytes" xml:"bytes"`          // bytes as an integer
	Decimal string `json:"decimal" xml:"decimal,attr"` // decimal is a base 10 value
	Binary  string `json:"binary" xml:"binary,attr"`   // binary is a base 2 value
}

func (d *Layout) Sizes() Sizes {
	value := UnsignedBinary4(d.Filesize)
	en := language.English
	return Sizes{
		Bytes:   value,
		Decimal: humanize.Decimal(int64(value), en),
		Binary:  humanize.Binary(int64(value), en),
	}
}

func UnsignedBinary4(b [4]byte) uint16 {
	var data uint16
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &data)
	if err != nil {
		log.Println("unsigned 4 byte, LE binary failed:", err)
	}
	return data
}
