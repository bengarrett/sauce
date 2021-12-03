// Type of data. SAUCE supports 8 different types and an undefined value.
// See http://www.acid.org/info/sauce/sauce.htm
package layout

// Datas is both the SAUCE DataType value and name.
type Datas struct {
	Type TypeOfData `json:"type" xml:"type"`
	Name string     `json:"name" xml:"name"`
}

// TypeOfData is the data type (SAUCE DataType).
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

func (d *Layout) DataType() Datas {
	dt := UnsignedBinary1(d.Datatype)
	return Datas{
		Type: TypeOfData(dt),
		Name: TypeOfData(dt).String(),
	}
}
