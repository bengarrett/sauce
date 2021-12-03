package layout

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	ComntID       string = "COMNT"                // SAUCE comment block identification. This should be equal to "COMNT".
	SauceID       string = "SAUCE"                // SAUCE identification, this should be equal to "SAUCE".
	SauceVersion  string = "00"                   // SAUCE version number, should be "00".
	SauceSeek     string = SauceID + SauceVersion // SAUCE sequence to seek.
	ComntLineSize int    = 64                     // A SAUCE comment line is 64 characters wide.
	ComntMaxLines int    = 255                    // A SAUCE comment is variable sized structure that holds up to 255 lines of additional information.
)

type (
	Data     []byte // Data is the input data.
	Id       [5]byte
	Version  [2]byte
	Title    [35]byte
	Author   [20]byte
	Group    [20]byte
	Date     [8]byte
	FileSize [4]byte
	DataType [1]byte
	FileType [1]byte
	TInfo1   [2]byte
	TInfo2   [2]byte
	TInfo3   [2]byte
	TInfo4   [2]byte
	Comments [1]byte
	TFlags   [1]byte
	TInfoS   [22]byte
)

type Layout struct {
	Id       Id
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

// Scan returns the position of the SAUCE00 ID or -1 if no ID exists.
// TODO: rename Seek? Index?
func Scan(b ...byte) (index int) {
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

func (b Id) String() string {
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
	i := Scan(d...)
	if i == -1 {
		return Layout{}
	}
	l := Layout{
		Id:       d.id(i),
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

func (d Data) author(i int) Author {
	var a Author
	const (
		start = 42
		end   = start + len(a)
	)
	for j, c := range d[start+i : end+i] {
		a[j] = c
	}
	return a
}

func (d Data) comments(i int) Comments {
	return Comments{d[i+104]}
}

func (d Data) Comnt(count Comments, sauceIndex int) Comnt {
	var block = Comnt{
		Count: count,
	}
	if int(UnsignedBinary1(count)) == 0 {
		return block
	}
	id, l := []byte(ComntID), len(d)
	var backwardsLoop = func(i int) int {
		return l - 1 - i
	}
	// search for the id sequence in b
	for i := range d {
		if i > ComntLineSize*ComntMaxLines {
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

func (d Data) dataType(i int) DataType {
	return DataType{d[i+94]}
}

func (d Data) date(i int) Date {
	var dt Date
	const (
		start = 82
		end   = start + len(dt)
	)
	for j, c := range d[start+i : end+i] {
		dt[j] = c
	}
	return dt
}

func (d Data) fileSize(i int) FileSize {
	return FileSize{d[i+90], d[i+91], d[i+92], d[i+93]}
}

func (d Data) fileType(i int) FileType {
	return FileType{d[i+95]}
}

func (d Data) group(i int) Group {
	var g Group
	const (
		start = 62
		end   = start + len(g)
	)
	for j, c := range d[start+i : end+i] {
		g[j] = c
	}
	return g
}

func (d Data) id(i int) Id {
	return Id{d[i+0], d[i+1], d[i+2], d[i+3], d[i+4]}
}

func (d Data) tFlags(i int) TFlags {
	return TFlags{d[i+105]}
}

func (d Data) title(i int) Title {
	var t Title
	const (
		start = 7
		end   = start + len(t)
	)
	for j, c := range d[start+i : end+i] {
		t[j] = c
	}
	return t
}

func (d Data) TInfo1(i int) TInfo1 {
	return TInfo1{d[i+96], d[i+97]}
}

func (d Data) TInfo2(i int) TInfo2 {
	return TInfo2{d[i+98], d[i+99]}
}

func (d Data) TInfo3(i int) TInfo3 {
	return TInfo3{d[i+100], d[i+101]}
}

func (d Data) TInfo4(i int) TInfo4 {
	return TInfo4{d[i+102], d[i+103]}
}

func (d Data) TInfoS(i int) TInfoS {
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

func (d Data) version(i int) Version {
	return Version{d[i+5], d[i+6]}
}

func UnsignedBinary1(b [1]byte) (value uint8) {
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		log.Println("unsigned 1 byte, LE binary failed:", err)
	}
	return value
}
