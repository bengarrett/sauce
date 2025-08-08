package layout

import (
	"bytes"
	"encoding/binary"
	"log"
)

//nolint:lll
const (
	ComntID       string = "COMNT"                // comntid is the comment block identification that must be "COMNT"
	SauceID       string = "SAUCE"                // sauceid is the SAUCE identification that must be "SAUCE"
	SauceVersion  string = "00"                   // sauce version number that is always "00"
	SauceSeek     string = SauceID + SauceVersion // sauceseek is the sauce id and version value to lookup
	ComntLineSize int    = 64                     // comntlinesize is the fix length in bytes of an individual comment line
	ComntMaxLines int    = 255                    // comntmaxlines is the maximum permitted number of lines for a block of comments
)

type (
	Data     []byte   // data is a copy of the input data
	ID       [5]byte  // id is the sauce identifier
	Version  [2]byte  // version number
	Title    [35]byte // title of the file
	Author   [20]byte // author is nickname or handle of creator of the file
	Group    [20]byte // group or company name the author worked for
	Date     [8]byte  // date the file was created
	FileSize [4]byte  // file size of original file size without the sauce data
	DataType [1]byte  // data type is the type of file, such as an raster image
	FileType [1]byte  // file type is the technical format of the file, such as a GIF
	TInfo1   [2]byte  // tinfo1 is type dependant numeric information field 1
	TInfo2   [2]byte  // tinfo2 is type dependant numeric information field 2
	TInfo3   [2]byte  // tinfo3 is type dependant numeric information field 3
	TInfo4   [2]byte  // tinfo4 is type dependant numeric information field 4
	Comments [1]byte  // comments are the number of lines in the extra SAUCE comment block
	TFlags   [1]byte  // tflags are the type dependant flags
	TInfoS   [22]byte // tinfos are the type dependant string information field
)

type Layout struct {
	ID       ID
	Version  Version
	Title    Title
	Author   Author
	Group    Group
	Date     Date
	Filesize FileSize
	Datatype DataType
	Filetype FileType
	Tinfo1   TInfo1
	Tinfo2   TInfo2
	Tinfo3   TInfo3
	Tinfo4   TInfo4
	Comments Comments
	TFlags   TFlags
	TInfoS   TInfoS
	Comnt    Comnt
}

// Index returns the position of the SAUCE00 ID or -1 if no ID exists.
func Index(b []byte) int {
	const sauceSize, maximum = 128, 512
	id, l := []byte(SauceSeek), len(b)
	backwardsLoop := func(i int) int {
		return l - 1 - i
	}
	// search for the id sequence in b
	const indexEnd = 6
	for i := range b {
		if i > maximum {
			break
		}
		i = backwardsLoop(i)
		if len(b) < sauceSize {
			break
		}
		// do matching in reverse
		if b[i] != id[indexEnd] {
			continue // 0
		}
		if b[i-1] != id[5] {
			continue // 0
		}
		if b[i-2] != id[4] {
			continue // E
		}
		if b[i-3] != id[3] {
			continue // C
		}
		if b[i-4] != id[2] {
			continue // U
		}
		if b[i-5] != id[1] {
			continue // A
		}
		if b[i-indexEnd] != id[0] {
			continue // S
		}
		return i - indexEnd
	}
	return -1
}

func (b ID) String() string {
	return string(b[:])
}

func (b Version) String() string {
	return string(b[:])
}

func (b Title) String() string {
	return string(b[:])
}

func (b Author) String() string {
	return string(b[:])
}

func (b Group) String() string {
	return string(b[:])
}

func (b Date) String() string {
	return string(b[:])
}

func (t TInfoS) String() string {
	const nul = 0
	s := ""
	for _, b := range t {
		if b == nul {
			continue
		}
		s += string(b)
	}
	return s
}

func (d Data) Extract() Layout {
	i := Index(d)
	if i == -1 {
		return Layout{
			Comnt: Comnt{
				Index:  -1,
				Length: 0,
				Count:  [1]byte{0},
				Lines:  []byte{},
			},
		}
	}
	l := Layout{
		ID:       d.id(i),
		Version:  d.version(i),
		Title:    d.title(i),
		Author:   d.author(i),
		Group:    d.group(i),
		Date:     d.date(i),
		Filesize: d.fileSize(i),
		Datatype: d.dataType(i),
		Filetype: d.fileType(i),
		Tinfo1:   d.TInfo1(i),
		Tinfo2:   d.TInfo2(i),
		Tinfo3:   d.TInfo3(i),
		Tinfo4:   d.TInfo4(i),
		TFlags:   d.tFlags(i),
		TInfoS:   d.TInfoS(i),
	}
	l.Comments = d.comments(i)
	l.Comnt = d.Comnt(l.Comments, i)
	return l
}

func (d Data) Comnt(count Comments, sauceIndex int) Comnt {
	block := Comnt{
		Count: count,
		Lines: []byte{},
	}
	if int(UnsignedBinary1(count)) == 0 {
		return block
	}
	id, l := []byte(ComntID), len(d)
	backwardsLoop := func(i int) int {
		return l - 1 - i
	}
	// search for the id sequence in b
	const maximum = ComntLineSize * ComntMaxLines
	for i := range d {
		if i > maximum {
			break
		}
		i = backwardsLoop(i)
		if i < ComntLineSize {
			break
		}
		if i >= sauceIndex {
			continue
		}
		// do matching in reverse
		if d[i-1] != id[4] {
			continue // T
		}
		if d[i-2] != id[3] {
			continue // N
		}
		if d[i-3] != id[2] {
			continue // M
		}
		if d[i-4] != id[1] {
			continue // O
		}
		if d[i-5] != id[0] {
			continue // C
		}
		block.Index = i
		block.Length = sauceIndex - block.Index
		block.Lines = d[i : i+block.Length]
		return block
	}
	return block
}

func (d Data) TInfo1(i int) TInfo1 {
	if len(d) < i+96 {
		return TInfo1{}
	}
	return TInfo1{d[i+96], d[i+97]}
}

func (d Data) TInfo2(i int) TInfo2 {
	if len(d) < i+98 {
		return TInfo2{}
	}
	return TInfo2{d[i+98], d[i+99]}
}

func (d Data) TInfo3(i int) TInfo3 {
	if len(d) < i+100 {
		return TInfo3{}
	}
	return TInfo3{d[i+100], d[i+101]}
}

func (d Data) TInfo4(i int) TInfo4 {
	if len(d) < i+102 {
		return TInfo4{}
	}
	return TInfo4{d[i+102], d[i+103]}
}

func (d Data) TInfoS(i int) TInfoS {
	if len(d) < i+105 {
		return TInfoS{}
	}
	var s TInfoS
	const (
		start = 106
		end   = start + len(s)
	)
	for j, c := range d[start+i : end+i] {
		if c == 0 {
			continue
		}
		s[j] = c
	}
	return s
}

// UnsignedBinary1 returns the unsigned 1 byte integer from b
// using little-endian byte order.
func UnsignedBinary1(b [1]byte) uint8 {
	var data uint8
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &data)
	if err != nil {
		log.Println("unsigned 1 byte, LE binary failed:", err)
	}
	return data
}

func (d Data) author(i int) Author {
	if len(d) < i+41 {
		return Author{}
	}
	const start = 42
	var a Author
	// source answer:
	// https://stackoverflow.com/questions/30285680/how-to-convert-slice-to-fixed-size-array
	copy(a[:], d[start+i:])
	return a
}

func (d Data) comments(i int) Comments {
	if len(d) < i+103 {
		return Comments{}
	}
	return Comments{d[i+104]}
}

func (d Data) dataType(i int) DataType {
	if len(d) < i+93 {
		return DataType{}
	}
	return DataType{d[i+94]}
}

func (d Data) date(i int) Date {
	if len(d) < i+81 {
		return Date{}
	}
	const start = 82
	var dt Date
	copy(dt[:], d[start+i:])
	return dt
}

func (d Data) fileSize(i int) FileSize {
	if len(d) < i+92 {
		return FileSize{}
	}
	b0 := d[i+90]
	b1 := d[i+91]
	b2 := d[i+92]
	b3 := d[i+93]
	return FileSize{b0, b1, b2, b3}
}

func (d Data) fileType(i int) FileType {
	if len(d) < i+94 {
		return FileType{}
	}
	return FileType{d[i+95]}
}

func (d Data) group(i int) Group {
	if len(d) < i+61 {
		return Group{}
	}
	const start = 62
	var g Group
	copy(g[:], d[start+i:])
	return g
}

func (d Data) id(i int) ID {
	return ID{d[i+0], d[i+1], d[i+2], d[i+3], d[i+4]}
}

func (d Data) tFlags(i int) TFlags {
	if len(d) < i+104 {
		return TFlags{}
	}
	return TFlags{d[i+105]}
}

func (d Data) title(i int) Title {
	const start = 7
	var t Title
	copy(t[:], d[start+i:])
	return t
}

func (d Data) version(i int) Version {
	return Version{d[i+5], d[i+6]}
}
