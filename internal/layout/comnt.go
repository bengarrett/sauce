// A SAUCE comment block is an optional, variable sized structure that holds up to 255 lines
// of additional information, each line 64 characters wide. There are as many comment lines
// as is mentioned in the Comments field of the SAUCE record.
// If the Comments field is set to 0, there should not be a comment block at all.
// See http://www.acid.org/info/sauce/sauce.htm
package layout

import (
	"bufio"
	"bytes"
	"strings"
)

type Comnt struct {
	Index  int
	Length int
	Count  Comments
	Lines  []byte
}

// Comment contain the optional SAUCE comment block.
type Comment struct {
	ID      string   `json:"id" xml:"id,attr"`
	Count   int      `json:"count" xml:"count,attr"`
	Index   int      `json:"-" xml:"-"`
	Comment []string `json:"lines" xml:"line"`
}

// CommentBlock parses the optional SAUCE comment block.
func (d *Layout) CommentBlock() (c Comment) {
	breakCount := len(strings.Split(string(d.Comnt.Lines), "\n"))
	c.ID = ComntID
	c.Count = int(UnsignedBinary1(d.Comnt.Count))
	c.Index = -1
	if d.Comnt.Index > 0 {
		c.Index = d.Comnt.Index - len(ComntID)
	}
	if breakCount > 0 {
		// comments with line breaks are technically invalid but they exist in the wild.
		// https://github.com/16colo-rs/16c/issues/67
		c.Comment = CommentByBreak(d.Comnt.Lines)
		return c
	}
	c.Comment = CommentByLine(d.Comnt.Lines)
	return c
}

// CommentByBreak parses the SAUCE comment by line break characters.
func CommentByBreak(b []byte) (lines []string) {
	r := bytes.NewReader(b)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// CommentByLine parses the SAUCE comment by lines of 64 characters.
func CommentByLine(b []byte) (lines []string) {
	s, l := "", 0
	resetLine := func() {
		s, l = "", 0
	}
	for _, c := range b {
		l++
		s += string(c)
		if l == ComntLineSize {
			lines = append(lines, s)
			resetLine()
		}
	}
	return lines
}
