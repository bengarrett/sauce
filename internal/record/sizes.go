package record

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/bengarrett/sauce/humanize"
	"golang.org/x/text/language"
)

// Sizes of the file data in multiples.
type Sizes struct {
	Bytes   uint16 `json:"bytes" xml:"bytes"`
	Decimal string `json:"decimal" xml:"decimal,attr"`
	Binary  string `json:"binary" xml:"binary,attr"`
}

func (d *Data) Sizes() Sizes {
	value := unsignedBinary4(d.Filesize)
	en := language.English
	return Sizes{
		Bytes:   value,
		Decimal: humanize.Decimal(int64(value), en),
		Binary:  humanize.Binary(int64(value), en),
	}
}
func unsignedBinary4(b [4]byte) (value uint16) {
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println("unsigned 4 byte, LE binary failed:", err)
	}
	return value
}
