package layout

// Type of data. SAUCE supports 8 different types and an undefined value.
// See http://www.acid.org/info/sauce/sauce.htm

// Datas is the SAUCE DataType value and name.
type Datas struct {
	Type TypeOfData `json:"type" xml:"type"` // typeofdata is the unsigned data type
	Name string     `json:"name" xml:"name"` // name of the data type
}

// TypeOfData is the SAUCE DataType.
type TypeOfData uint

const (
	Nones       TypeOfData = iota // undefined filetype
	Characters                    // characters and plain text based files
	Bitmaps                       // bitmap, graphic and animation files
	Vectors                       // vector graphic files
	Audios                        // audio and sound files
	BinaryTexts                   // binary texts that are raw memory copies of a text mode screen, also known as a 'bin' file
	XBins                         // xbin or extended bin file
	Archives                      // archived file such as a zip package
	Executables                   // executable file that is an application or program launcher
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
