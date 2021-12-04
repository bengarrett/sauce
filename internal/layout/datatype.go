// Type of data. SAUCE supports 8 different types and an undefined value.
// See http://www.acid.org/info/sauce/sauce.htm
package layout

// Datas is the SAUCE DataType value and name.
type Datas struct {
	Type TypeOfData `json:"type" xml:"type"` // A unsigned integer DataType.
	Name string     `json:"name" xml:"name"` // Name of the DataType.
}

// TypeOfData is the SAUCE DataType.
type TypeOfData uint

const (
	Nones       TypeOfData = iota // Undefined filetype.
	Characters                    // A character based file.
	Bitmaps                       // Bitmap graphic and animation files.
	Vectors                       // A vector graphic file.
	Audios                        // An audio file.
	BinaryTexts                   // This is a raw memory copy of a text mode screen. Also known as a .BIN file.
	XBins                         // An XBin or eXtended BIN file.
	Archives                      // An archive file.
	Executables                   // A executable file.
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
