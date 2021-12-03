package layout_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func Test_data_FileType(t *testing.T) {
	type fields struct {
		datatype layout.DataType
		filetype layout.FileType
	}
	empty := layout.Files{layout.TypeOfFile(layout.Nones),
		layout.Undefined.String()}
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
		want   layout.Files
	}{
		{"empty", fields{}, empty},
		{"out of range", outofrange, layout.Files{
			layout.TypeOfFile(layout.Nones), layout.ErrFileType.Error()}},
		{"nones", none, empty},
		{"characters", chars, layout.Files{
			layout.TypeOfFile(layout.Ascii), layout.Ascii.String()}},
		{"bitmaps", bmp, layout.Files{
			layout.TypeOfFile(layout.Gif), layout.Gif.String()}},
		{"vectors", vec, layout.Files{
			layout.TypeOfFile(layout.Dxf), layout.Dxf.String()}},
		{"audios", aud, layout.Files{
			layout.TypeOfFile(layout.Mod), layout.Mod.String()}},
		{"binarytexts", btxt, layout.Files{
			layout.TypeOfFile(layout.BinaryScreenImage), layout.BinaryScreenImage.String()}},
		{"xbins", xb, layout.Files{
			layout.TypeOfFile(layout.ExtendedBin), layout.ExtendedBin.String()}},
		{"archives", arc, layout.Files{
			layout.TypeOfFile(layout.Zip), layout.Zip.String()}},
		{"executable", exe, layout.Files{
			layout.TypeOfFile(layout.Exe), layout.Exe.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &layout.Layout{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
			}
			if got := d.FileType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Layout.FileType() = %v, want %v", got, tt.want)
			}
		})
	}
}
