package layout

import (
	"bytes"
)

const (
	// cmtCapacity is the initial capacity for comment line slices.
	cmtCapacity = 16
)

// A SAUCE comment block is an optional, variable sized structure that holds up to 255 lines
// of additional information, each line 64 characters wide. There are as many comment lines
// as is mentioned in the Comments field of the SAUCE record.
// If the Comments field is set to 0, there should not be a comment block at all.
// See http://www.acid.org/info/sauce/sauce.htm

type Comnt struct {
	Index  int      // index is the calculated starting position of the comment block
	Length int      // length is the calculated length of the comment block
	Count  Comments // count are the reported number of lines in the SAUCE comment block
	Lines  []byte   // lines of text
}

// Comment contains the optional SAUCE comment block.
// A SAUCE comment block is an optional, variable sized structure that holds
// up to 255 lines of additional information, each line 64 characters wide.
type Comment struct {
	ID      string   `json:"id" xml:"id,attr"`       // id is the SAUCE comment block identification, this should be "COMNT"
	Count   int      `json:"count" xml:"count,attr"` // count are the reported number of lines in the SAUCE comment block
	Index   int      `json:"-" xml:"-"`              // index are the calculated starting position of the comment block
	Comment []string `json:"lines" xml:"line"`       // comment value, each comment line should be comprised of 64 characters
}

// CommentBlock parses the optional SAUCE comment block.
func (d *Layout) CommentBlock() Comment {
	hasLineBreak := bytes.ContainsAny(d.Comnt.Lines, "\n\r")
	var c Comment
	c.ID = ComntID
	c.Count = int(UnsignedBinary1(d.Comnt.Count))
	c.Index = -1
	if d.Comnt.Index > 0 {
		c.Index = d.Comnt.Index - len(ComntID)
	}
	if hasLineBreak {
		// comments with line breaks are technically invalid but they exist in the wild.
		// https://github.com/16colo-rs/16c/issues/67
		c.Comment = CommentByBreak(d.Comnt.Lines)
		return c
	}
	c.Comment = CommentByLine(d.Comnt.Lines)
	return c
}

// CommentByBreak parses the SAUCE comment by line break characters.
func CommentByBreak(b []byte) []string {
	if len(b) == 0 {
		return []string{}
	}

	lines := make([]string, 0, cmtCapacity)
	start := 0

	for i := 0; i < len(b); i++ {
		if b[i] == '\n' || b[i] == '\r' {
			// Handle line ending
			if i > start {
				lines = append(lines, string(b[start:i]))
			}
			start = i + 1

			// Handle \r\n sequences
			if i+1 < len(b) && b[i] == '\r' && b[i+1] == '\n' {
				i++
				start = i + 1
			}
		}
	}
	if start < len(b) {
		lines = append(lines, string(b[start:]))
	}

	return lines
}

// CommentByLine parses the SAUCE comment by lines of 64 characters.
func CommentByLine(b []byte) []string {
	if len(b) == 0 {
		return []string{}
	}
	expectedLines := len(b) / ComntLineSize
	if len(b)%ComntLineSize != 0 {
		expectedLines++
	}
	lines := make([]string, 0, expectedLines)
	for i := 0; i < len(b); i += ComntLineSize {
		end := min(i+ComntLineSize, len(b))
		lines = append(lines, string(b[i:end]))
	}
	return lines
}
