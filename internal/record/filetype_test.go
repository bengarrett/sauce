package record_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func Test_data_FileType(t *testing.T) {
	type fields struct {
		datatype record.DataType
		filetype record.FileType
	}
	empty := record.Files{record.TypeOfFile(record.Nones),
		record.Undefined.String()}
	outofrange := fields{[1]byte{255}, [1]byte{255}}
	none := fields{[1]byte{0}, [1]byte{0}}
	chars := fields{[1]byte{1}, [1]byte{0}}
	bmp := fields{[1]byte{2}, [1]byte{0}}
	vec := fields{[1]byte{3}, [1]byte{0}}
	aud := fields{[1]byte{4}, [1]byte{0}}
	btxt := fields{[1]byte{5}, [1]byte{0}}
	xb := fields{[1]byte{6}, [1]byte{0}}
	arc := fields{[1]byte{7}, [1]byte{0}}
	exe := fields{[1]byte{8}, [1]byte{0}}
	tests := []struct {
		name   string
		fields fields
		want   record.Files
	}{
		{"empty", fields{}, empty},
		{"out of range", outofrange, record.Files{
			record.TypeOfFile(record.Nones), record.ErrFileType.Error()}},
		{"nones", none, empty},
		{"characters", chars, record.Files{
			record.TypeOfFile(record.Ascii), record.Ascii.String()}},
		{"bitmaps", bmp, record.Files{
			record.TypeOfFile(record.Gif), record.Gif.String()}},
		{"vectors", vec, record.Files{
			record.TypeOfFile(record.Dxf), record.Dxf.String()}},
		{"audios", aud, record.Files{
			record.TypeOfFile(record.Mod), record.Mod.String()}},
		{"binarytexts", btxt, record.Files{
			record.TypeOfFile(record.BinaryScreenImage), record.BinaryScreenImage.String()}},
		{"xbins", xb, record.Files{
			record.TypeOfFile(record.ExtendedBin), record.ExtendedBin.String()}},
		{"archives", arc, record.Files{
			record.TypeOfFile(record.Zip), record.Zip.String()}},
		{"executable", exe, record.Files{
			record.TypeOfFile(record.Exe), record.Exe.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Layout{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
			}
			if got := d.FileType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Layout.FileType() = %v, want %v", got, tt.want)
			}
		})
	}
}
