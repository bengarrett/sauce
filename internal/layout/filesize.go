package layout

import (
	"encoding/binary"

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
	//nolint:gosec
	return uint16(binary.LittleEndian.Uint32(b[:]))
}
