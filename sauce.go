// Package sauce parses SAUCE (Standard Architecture for Universal Comment Extensions) metadata.
// http://www.acid.org/info/sauce/sauce.htm
package sauce

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/bengarrett/sauce/internal/layout"
)

// Record is the container for SAUCE data.
type Record struct {
	ID       string         `json:"id" xml:"id,attr"`
	Version  string         `json:"version" xml:"version,attr"`
	Title    string         `json:"title" xml:"title"`
	Author   string         `json:"author" xml:"author"`
	Group    string         `json:"group" xml:"group"`
	Date     layout.Dates   `json:"date" xml:"date"`
	FileSize layout.Sizes   `json:"filesize" xml:"filesize"`
	Data     layout.Datas   `json:"dataType"  xml:"data_type"`
	File     layout.Files   `json:"fileType" xml:"file_type"`
	Info     layout.Infos   `json:"typeInfo"  xml:"type_info"`
	Desc     string         `json:"-" xml:"-"`
	Comnt    layout.Comment `json:"comments" xml:"comments"`
}

// Contains returns true if a valid SAUCE record is contained within the bytes.
func Contains(b []byte) bool {
	const missing int = -1
	return layout.Scan(b) != missing
}

// Decode and extract the SAUCE data contained within the bytes.
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

// Index ...
func Index(b []byte) int {
	return layout.Scan(b)
}

func NewRecord(r io.Reader) (*Record, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	d := Decode(b)
	return &d, nil
}

func JSON(r *Record) ([]byte, error) {
	return json.Marshal(r)
}

func JSONIndent(r *Record, indent string) ([]byte, error) {
	return json.MarshalIndent(r, "", indent)
}
