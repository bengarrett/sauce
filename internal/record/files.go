package record

// Files, both the SAUCE FileType value and name.
type Files struct {
	Type TypeOfFile `json:"type" xml:"type"`
	Name string     `json:"name" xml:"name"`
}

// FileType is the type of file (SAUCE FileType).
type TypeOfFile uint

func (d *Data) FileType() Files {
	data, file := unsignedBinary1(d.Datatype), unsignedBinary1(d.Filetype)
	switch TypeOfData(data) {
	case None:
		return Files{TypeOfFile(None), None.String()}
	case Characters:
		c := CharacterBase(file)
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
		return Files{TypeOfFile(BinaryTexts), BinaryTexts.String()}
	case XBins:
		return Files{TypeOfFile(XBins), XBins.String()}
	case Archives:
		a := Archive(file)
		return Files{TypeOfFile(a), a.String()}
	case Executables:
		return Files{TypeOfFile(Executables), Executables.String()}
	default:
		return Files{TypeOfFile(0), "error"}
	}
}
