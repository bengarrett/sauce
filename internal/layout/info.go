package layout

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Infos includes the SAUCE fields dependant on both DataType and FileType.
type Infos struct {
	// Info1 dependant numeric information field 1.
	Info1 Info `json:"1" xml:"type1"`
	// Info2 dependant numeric information field 2.
	Info2 Info `json:"2" xml:"type2"`
	// Info3 dependant numeric information field 3.
	Info3 Info `json:"3" xml:"type3"`
	// Flags are file type dependant flags.
	Flags ANSIFlags `json:"flags" xml:"flags"`
	// Font field allows an author to provide a clue to the viewer/editor which font to use to render the image.
	Font string `json:"fontName" xml:"fontname"`
}

// Info is the type for the SAUCE TInfo1, TInfo2, TInfo3 and TInfo4 fields.
type Info struct {
	// Value of the field.
	Value uint16 `json:"value" xml:"value"`
	// Info is a description of the value.
	Info string `json:"info" xml:"type,attr"`
}

const (
	chrw = "character width"
	nol  = "number of lines"
	pxw  = "pixel width"
)

func (d *Layout) InfoType() Infos {
	dt, ft := UnsignedBinary1(d.Datatype),
		UnsignedBinary1(d.Filetype)
	t1, t2, t3 := UnsignedBinary2(d.Tinfo1),
		UnsignedBinary2(d.Tinfo2),
		UnsignedBinary2(d.Tinfo3)
	flag := Flags(UnsignedBinary1(d.TFlags))
	ti := Infos{
		Info{t1, ""},
		Info{t2, ""},
		Info{t3, ""},
		flag.Parse(),
		d.TInfoS.String(),
	}
	switch TypeOfData(dt) {
	case Nones:
		return ti // golangci-lint deadcode placeholder
	case Characters:
		ti.character(ft)
		return ti
	case Bitmaps:
		switch Bitmap(ft) {
		case Gif, Pcx, Lbm, Tga, Fli, Flc, Bmp, Gl, Dl, Wpg, Png, Jpg, Mpg, Avi:
			ti.Info1.Info = pxw
			ti.Info2.Info = "pixel height"
			ti.Info3.Info = "pixel depth"
		}
	case Vectors:
		switch Vector(ft) {
		case Dxf, Dwg, Wpvg, Kinetix:
			return ti
		}
	case Audios:
		ti.audio(ft)
		return ti
	case BinaryTexts:
		return ti
	case XBins:
		ti.Info1.Info = chrw
		ti.Info2.Info = nol
	case Archives:
		switch Archive(ft) {
		case Zip, Arj, Lzh, Arc, Tar, Zoo, Rar, Uc2, Pak, Sqz:
			return ti
		}
	case Executables:
		return ti
	}
	return ti
}

func (ti *Infos) character(ft uint8) {
	switch Character(ft) {
	case ASCII, Ansi, AnsiMation, PCBoard, Avatar, TundraDraw:
		ti.Info1.Info = chrw
		ti.Info2.Info = nol
	case RipScript:
		ti.Info1.Info = pxw
		ti.Info2.Info = "character screen height"
		ti.Info3.Info = "number of colors"
	case HTML, Source:
		return
	}
}

func UnsignedBinary2(b [2]byte) uint16 {
	var data uint16
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &data)
	if err != nil {
		log.Println("unsigned 2 bytes, LE binary failed:", err)
	}
	return data
}
