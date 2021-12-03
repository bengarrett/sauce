// Package sauce parses SAUCE (Standard Architecture for Universal Comment Extensions) metadata.
// http://www.acid.org/info/sauce/sauce.htm
package sauce

import (
	"strings"

	"github.com/bengarrett/sauce/internal/record"
)

// Record is the container for SAUCE data.
type Record struct {
	ID       string         `json:"id" xml:"id,attr"`
	Version  string         `json:"version" xml:"version,attr"`
	Title    string         `json:"title" xml:"title"`
	Author   string         `json:"author" xml:"author"`
	Group    string         `json:"group" xml:"group"`
	Date     record.Dates   `json:"date" xml:"date"`
	FileSize record.Sizes   `json:"filesize" xml:"filesize"`
	Data     record.Datas   `json:"dataType"  xml:"data_type"`
	File     record.Files   `json:"fileType" xml:"file_type"`
	Info     record.Infos   `json:"typeInfo"  xml:"type_info"`
	Desc     string         `json:"-" xml:"-"`
	Comnt    record.Comment `json:"comments" xml:"comments"`
}

// Parse and extract the record data. // todo: rename to marshal
func Parse(data ...byte) Record {
	const empty = "\x00\x00"
	r := record.Data(data)
	d := r.Extract()
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
