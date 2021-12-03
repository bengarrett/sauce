package record

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
	Record []byte // Src is the input data.

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

func (r Record) Extract() Layout {
	i := Scan(r...)
	if i == -1 {
		return Layout{}
	}
	d := Layout{
		Id:       r.id(i),
		Version:  r.version(i),
		Title:    r.title(i),
		Author:   r.author(i),
		Group:    r.group(i),
		Date:     r.date(i),
		Filesize: r.fileSize(i),
		Datatype: r.dataType(i),
		Filetype: r.fileType(i),
		Tinfo1:   r.TInfo1(i),
		Tinfo2:   r.TInfo2(i),
		Tinfo3:   r.TInfo3(i),
		Tinfo4:   r.TInfo4(i),
		TFlags:   r.tFlags(i),
		TInfoS:   r.TInfoS(i),
	}
	d.Comments = r.comments(i)
	d.Comnt = r.Comnt(d.Comments, i)
	return d
}

func (r Record) author(i int) Author {
	var a Author
	const (
		start = 42
		end   = start + len(a)
	)
	for j, c := range r[start+i : end+i] {
		a[j] = c
	}
	return a
}

func (r Record) comments(i int) Comments {
	return Comments{r[i+104]}
}

func (r Record) Comnt(count Comments, sauceIndex int) Comnt {
	var block = Comnt{
		Count: count,
	}
	if int(UnsignedBinary1(count)) == 0 {
		return block
	}
	id, l := []byte(ComntID), len(r)
	var backwardsLoop = func(i int) int {
		return l - 1 - i
	}
	// search for the id sequence in b
	for i := range r {
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
		if r[i-1] != id[4] {
			continue // T
		}
		if r[i-2] != id[3] {
			continue // N
		}
		if r[i-3] != id[2] {
			continue // M
		}
		if r[i-4] != id[1] {
			continue // O
		}
		if r[i-5] != id[0] {
			continue // C
		}
		block.Index = i
		block.Length = sauceIndex - block.Index
		block.Lines = r[i : i+block.Length]
		return block
	}
	return block
}

func (r Record) dataType(i int) DataType {
	return DataType{r[i+94]}
}

func (r Record) date(i int) Date {
	var d Date
	const (
		start = 82
		end   = start + len(d)
	)
	for j, c := range r[start+i : end+i] {
		d[j] = c
	}
	return d
}

func (r Record) fileSize(i int) FileSize {
	return FileSize{r[i+90], r[i+91], r[i+92], r[i+93]}
}

func (r Record) fileType(i int) FileType {
	return FileType{r[i+95]}
}

func (r Record) group(i int) Group {
	var g Group
	const (
		start = 62
		end   = start + len(g)
	)
	for j, c := range r[start+i : end+i] {
		g[j] = c
	}
	return g
}

func (r Record) id(i int) Id {
	return Id{r[i+0], r[i+1], r[i+2], r[i+3], r[i+4]}
}

func (r Record) tFlags(i int) TFlags {
	return TFlags{r[i+105]}
}

func (r Record) title(i int) Title {
	var t Title
	const (
		start = 7
		end   = start + len(t)
	)
	for j, c := range r[start+i : end+i] {
		t[j] = c
	}
	return t
}

func (r Record) TInfo1(i int) TInfo1 {
	return TInfo1{r[i+96], r[i+97]}
}

func (r Record) TInfo2(i int) TInfo2 {
	return TInfo2{r[i+98], r[i+99]}
}

func (r Record) TInfo3(i int) TInfo3 {
	return TInfo3{r[i+100], r[i+101]}
}

func (r Record) TInfo4(i int) TInfo4 {
	return TInfo4{r[i+102], r[i+103]}
}

func (r Record) TInfoS(i int) TInfoS {
	var s TInfoS
	const (
		start = 106
		end   = start + len(s)
	)
	for j, c := range r[start+i : end+i] {
		if c == 0 {
			continue
		}
		s[j] = c
	}
	return s
}

func (r Record) version(i int) Version {
	return Version{r[i+5], r[i+6]}
}

func UnsignedBinary1(b [1]byte) (value uint8) {
	buf := bytes.NewReader(b[:])
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		log.Println("unsigned 1 byte, LE binary failed:", err)
	}
	return value
}
