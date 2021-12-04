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
	ID       string         `json:"id" xml:"id,attr"`           // ID is the SAUCE identification, it should equal "SAUCE".
	Version  string         `json:"version" xml:"version,attr"` // Version is the SAUCE version, it should be "00".
	Title    string         `json:"title" xml:"title"`          // Title of the file.
	Author   string         `json:"author" xml:"author"`        // The nick, name or handle of the creator of the file.
	Group    string         `json:"group" xml:"group"`          // The name of the group or company the creator is employed by.
	Date     layout.Dates   `json:"date" xml:"date"`            // The date the file was created.
	FileSize layout.Sizes   `json:"filesize" xml:"filesize"`    // The original file size not including the SAUCE information.
	Data     layout.Datas   `json:"dataType"  xml:"data_type"`  // Type of data.
	File     layout.Files   `json:"fileType" xml:"file_type"`   // Type of file.
	Info     layout.Infos   `json:"typeInfo"  xml:"type_info"`  // Type dependant numeric informations.
	Desc     string         `json:"-" xml:"-"`                  // Humanized description of the file.
	Comnt    layout.Comment `json:"comments" xml:"comments"`    // Comments or notes from the creator.
}

// Decode the SAUCE data contained within b.
func Decode(b []byte) Record {
	const empty = "\x00\x00"
	d := layout.Data(b).Extract()
	if string(d.Version[:]) == empty {
		return Record{}
	}
	return Record{
		ID:       d.Id.String(),
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
		return nil, err
	}
	d := Decode(b)
	return &d, nil
}

// JSON returns the JSON encoding of the r SAUCE record.
func (r *Record) JSON() ([]byte, error) {
	return json.Marshal(r)
}

// JSONIndent is like JSON but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with one
// or more copies of indent according to the indentation nesting.
func (r *Record) JSONIndent(indent string) ([]byte, error) {
	return json.MarshalIndent(r, "", indent)
}

// Valid reports the completeness of the r SAUCE record.
func (r *Record) Valid() bool {
	return r.ID == ID && r.Version == Version
}

// XML returns the XML encoding of the r SAUCE record.
func (r *Record) XML() ([]byte, error) {
	return xml.Marshal(r)
}

// XMLIndent is like XML but applies Indent to format the output.
// Each XML element in the output will begin on a new line beginning with one
// or more copies of indent according to the indentation nesting.
func (r *Record) XMLIndent(indent string) ([]byte, error) {
	return xml.MarshalIndent(r, "", indent)
}
