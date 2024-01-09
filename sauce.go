// Package sauce is a Go module that parses SAUCE (Standard Architecture for
// Universal Comment Extensions) metadata.
//
// See http://www.acid.org/info/sauce/sauce.htm.
//
// The Standard Architecture for Universal Comment Extensions or SAUCE as it is
// more commonly known, is an architecture or protocol for attaching meta data
// or comments to files. Mainly intended for ANSI art files, SAUCE has always
// had provisions for many different file types.
package sauce

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/bengarrett/sauce/internal/layout"
)

const (
	ID        = "SAUCE"
	Version   = "00"
	SauceDate = "20060102" // Date format is CCYYMMDD (century, year, month, day).
)

// Contains reports whether a valid SAUCE record is within b.
func Contains(b []byte) bool {
	const missing int = -1
	return layout.Index(b) != missing
}

// Index returns the index of the instance of the SAUCE ID in b,
// or -1 if it is not present in b.
func Index(b []byte) int {
	return layout.Index(b)
}

// Trim returns b without any SAUCE metadata.
func Trim(b []byte) []byte {
	const none = -1
	si := Index(b)
	if si == none {
		return b
	}
	// the optional comnt index always prefixes the sauce index
	sr := Decode(b)
	if ci := sr.Comnt.Index; ci > none {
		if ci > len(b) {
			return nil
		}
		return b[:ci]
	}
	if si > len(b) {
		return nil
	}
	return b[:si]
}

// Record is the SAUCE data structure that corresponds with the SAUCE Layout fields.
type Record struct {
	// ID is the SAUCE identification, it must equal "SAUCE".
	ID string `json:"id" xml:"id,attr"`
	// Version is the SAUCE version, it must equal "00".
	Version string `json:"version" xml:"version,attr"`
	// Title of the file.
	Title string `json:"title" xml:"title"`
	// Author is nick, name or handle of the creator of the file.
	Author string `json:"author" xml:"author"`
	// Group is the name of the group or company the creator is employed by.
	Group string `json:"group" xml:"group"`
	// Date the file was created.
	Date layout.Dates `json:"date" xml:"date"`
	// FileSize of the original file, not including any appended SAUCE data.
	FileSize layout.Sizes `json:"filesize" xml:"filesize"`
	// Data type used by the file.
	Data layout.Datas `json:"dataType"  xml:"data_type"`
	// File type of file.
	File layout.Files `json:"fileType" xml:"file_type"`
	// Info contains type dependant numeric informations.
	Info layout.Infos `json:"typeInfo"  xml:"type_info"`
	// Desc is a humanized description of the file.
	Desc string `json:"-" xml:"-"`
	// Comnt contains comments or notes from the creator.
	Comnt layout.Comment `json:"comments" xml:"comments"`
}

// Decode the SAUCE data contained within b.
func Decode(b []byte) Record {
	const empty = "\x00\x00"
	d := layout.Data(b).Extract()
	if string(d.Version[:]) == empty {
		return Record{}
	}
	return Record{
		ID:       d.ID.String(),
		Version:  d.Version.String(),
		Title:    strings.TrimSpace(d.Title.String()),
		Author:   strings.TrimSpace(d.Author.String()),
		Group:    strings.TrimSpace(d.Group.String()),
		Date:     d.Dates(),
		FileSize: d.Sizes(),
		Data:     d.DataType(),
		File:     d.FileType(),
		Info:     d.InfoType(),
		Desc:     d.Description(),
		Comnt:    d.CommentBlock(),
	}
}

// NewRecord creates a new SAUCE record from r.
func NewRecord(r io.Reader) (*Record, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("new sauce record: %w", err)
	}
	d := Decode(b)
	return &d, nil
}

// JSON returns the JSON encoding of the r SAUCE record.
func (r *Record) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("record as json: %w", err)
	}
	return b, nil
}

// JSONIndent is like JSON but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with one
// or more copies of indent according to the indentation nesting.
func (r *Record) JSONIndent(indent string) ([]byte, error) {
	b, err := json.MarshalIndent(r, "", indent)
	if err != nil {
		return nil, fmt.Errorf("record as json indent: %w", err)
	}
	return b, nil
}

// Valid reports the completeness of the r SAUCE record.
func (r *Record) Valid() bool {
	return r.ID == ID && r.Version == Version
}

// XML returns the XML encoding of the r SAUCE record.
func (r *Record) XML() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("record as xml: %w", err)
	}
	return b, nil
}

// XMLIndent is like XML but applies Indent to format the output.
// Each XML element in the output will begin on a new line beginning with one
// or more copies of indent according to the indentation nesting.
func (r *Record) XMLIndent(indent string) ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", indent)
	if err != nil {
		return nil, fmt.Errorf("record as xml indent: %w", err)
	}
	return b, nil
}
