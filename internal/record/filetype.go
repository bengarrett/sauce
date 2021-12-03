// Type of file.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package record

import "errors"

var ErrFileType = errors.New("unknown filetype")

// Files, both the SAUCE FileType value and name.
type Files struct {
	Type TypeOfFile `json:"type" xml:"type"`
	Name string     `json:"name" xml:"name"`
}

// TypeOfFile is the type of file (SAUCE FileType).
type TypeOfFile uint

func (d *Data) FileType() Files {
	data, file :=
		UnsignedBinary1(d.Datatype),
		UnsignedBinary1(d.Filetype)
	switch TypeOfData(data) {
	case Nones:
		n := None(file)
		return Files{TypeOfFile(n), n.String()}
	case Characters:
		c := Character(file)
		return Files{TypeOfFile(c), c.String()}
	case Bitmaps:
		b := Bitmap(file)
		return Files{TypeOfFile(b), b.String()}
	case Vectors:
		v := Vector(file)
		return Files{TypeOfFile(v), v.String()}
	case Audios:
		a := Audio(file)
		return Files{TypeOfFile(a), a.String()}
	case BinaryTexts:
		b := BinaryText(file)
		return Files{TypeOfFile(b), b.String()}
	case XBins:
		x := XBin(file)
		return Files{TypeOfFile(x), x.String()}
	case Archives:
		a := Archive(file)
		return Files{TypeOfFile(a), a.String()}
	case Executables:
		e := Executable(file)
		return Files{TypeOfFile(e), e.String()}
	default:
		var empty TypeOfFile
		return Files{empty, ErrFileType.Error()}
	}
}
