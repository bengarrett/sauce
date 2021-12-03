package layout

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Infos includes the SAUCE fields dependant on both DataType and FileType.
type Infos struct {
	Info1 Info      `json:"1" xml:"1"`
	Info2 Info      `json:"2" xml:"2"`
	Info3 Info      `json:"3" xml:"3"`
	Flags ANSIFlags `json:"flags" xml:"flags"`
	Font  string    `json:"fontName" xml:"fontname"`
}

// Info includes the SAUCE TInfo value and meaning.
type Info struct {
	Value uint16 `json:"value" xml:"value"`
	Info  string `json:"info" xml:"info,attr"`
}

const (
	chrw = "character width"
	nol  = "number of lines"
	pxw  = "pixel width"
)

func (d *Layout) InfoType() Infos {
	dt, ft :=
		UnsignedBinary1(d.Datatype),
		UnsignedBinary1(d.Filetype)
	t1, t2, t3 :=
		UnsignedBinary2(d.Tinfo1),
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
	case Ascii, Ansi, AnsiMation, PCBoard, Avatar, TundraDraw:
		ti.Info1.Info = chrw
		ti.Info2.Info = nol
	case RipScript:
		ti.Info1.Info = pxw
		ti.Info2.Info = "character screen height"
		ti.Info3.Info = "number of colors"
	case Html, Source:
		return
	}
}

func UnsignedBinary2(b [2]byte) (value uint16) {
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		log.Println("unsigned 2 bytes, LE binary failed:", err)
	}
	return value
}