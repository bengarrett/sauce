// Package sauce is a Go module that parses [SAUCE] metadata.
//
// # What is SAUCE?
//
// The Standard Architecture for Universal Comment Extensions is an architecture
// and protocol for attaching meta data and comments to files. While intended
// for [ANSI art files], SAUCE has always had provisions for many different file types.
//
// # Why SAUCE?
//
// From the original [SAUCE] specification:
//
// In the early 1990s there was a growing popularity in ANSI artwork. The ANSI art groups regularly released the works of their members over a certain period. Some of the bigger groups also included specialised viewers in each ‘artpack’. One of the problems with these artpacks was a lack of standardized way to provide meta data to the art, such as the title of the artwork, the author, the group, ... Some of the specialised viewers provided such information for a specific artpack either by encoding it as part of the executable, or by having some sort of database or list. However every viewer did it their own way. This meant you either had to use the viewer included with the artpack, or had to make do without the extra info. SAUCE was created to address that need. So if you wanted to, you could use your prefered viewer to view the art in a certain artpack, or even store the art files you liked in a separate folder while retaining the meta data.
//
// The goal was simple, but the way to get there certainly was not. Logistically, we wanted as many art groups as possible to support it. Technically, we wanted a system that was easy to implement and – if at all possible – manage to provide this meta data while still being compatible with all the existing software such as ANSI viewers, and Bulletin Board Software.
//
// [SAUCE]: http://www.acid.org/info/sauce/sauce.htm
// [ANSI art files]: https://16colo.rs
package sauce

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/bengarrett/sauce/internal/layout"
)

// SAUCE identifier and version.
const (
	ID      = "SAUCE"
	Version = "00" // the version is always 00
)

// Date format layout.
const Date = "20060102"

// EOF is the end-of-file marker, otherwise known as SUB, the substitute character.
const EOF byte = 26

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

// Trim returns b without any SAUCE metadata and the optional end-of-file marker.
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
		// trim the eof marker
		if b[ci-1] == EOF && ci > 2 {
			return b[:ci-2]
		}
		return b[:ci]
	}
	if si > len(b) {
		return nil
	}
	// trim the eof marker
	if b[si-1] == EOF && si > 2 {
		return b[:si-2]
	}
	return b[:si]
}

// Record is the SAUCE data structure that corresponds with the SAUCE Layout fields.
type Record struct {
	ID       string         `json:"id" xml:"id,attr"`           // SAUCE identification.
	Version  string         `json:"version" xml:"version,attr"` // version must equal "00".
	Title    string         `json:"title" xml:"title"`          // title of the file.
	Author   string         `json:"author" xml:"author"`        // author of the file.
	Group    string         `json:"group" xml:"group"`          // author employer or membership.
	Date     layout.Dates   `json:"date" xml:"date"`            // date of creation or release.
	FileSize layout.Sizes   `json:"filesize" xml:"filesize"`    // size of file in bytes without SAUCE.
	Data     layout.Datas   `json:"dataType" xml:"data_type"`   // data type of file.
	File     layout.Files   `json:"fileType" xml:"file_type"`   // file type of file.
	Info     layout.Infos   `json:"typeInfo" xml:"type_info"`   // file type dependant information.
	Desc     string         `json:"-" xml:"-"`                  // description of the file.
	Comnt    layout.Comment `json:"comments" xml:"comments"`    // comment block or notes.
}

// Decode the SAUCE data contained within b.
func Decode(b []byte) Record {
	const empty = "\x00\x00"
	d := layout.Data(b).Extract()
	if string(d.Version[:]) == empty {
		return Record{
			ID:      "",
			Version: "",
			Title:   "",
			Author:  "",
			Group:   "",
			Date: layout.Dates{
				Value: "",
				Time:  time.Time{},
				Epoch: 0,
			},
			FileSize: layout.Sizes{
				Bytes:   0,
				Binary:  "",
				Decimal: "",
			},
			Data: layout.Datas{
				Type: d.DataType().Type,
				Name: "",
			},
			File: layout.Files{
				Type: d.FileType().Type,
				Name: "",
			},
			Info: layout.Infos{
				Info1: layout.Info{},
				Info2: layout.Info{},
				Info3: layout.Info{},
				Flags: layout.ANSIFlags{},
			},
			Desc: "",
			Comnt: layout.Comment{
				ID:      "",
				Count:   0,
				Index:   -1,
				Comment: []string{},
			},
		}
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

// Read and return the SAUCE record in r.
func Read(r io.Reader) (*Record, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read sauce record: %w", err)
	}
	d := Decode(b)
	return &d, nil
}

// NewRecord is deprecated, use [Read].
func NewRecord(r io.Reader) (*Record, error) {
	return Read(r)
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
