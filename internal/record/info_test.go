package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func Test_data_InfoType(t *testing.T) {
	type fields struct {
		datatype record.DataType
		filetype record.FileType
		tinfo1   record.TInfo1
		tinfo2   record.TInfo2
		tinfo3   record.TInfo3
		tFlags   record.TFlags
	}
	ascii := fields{datatype: [1]byte{1}, filetype: [1]byte{0}}
	htm := fields{datatype: [1]byte{1}, filetype: [1]byte{6}}
	rip := fields{datatype: [1]byte{1}, filetype: [1]byte{3}}
	pcx := fields{datatype: [1]byte{2}, filetype: [1]byte{1}}
	dxf := fields{datatype: [1]byte{3}, filetype: [1]byte{0}}
	samp16 := fields{datatype: [1]byte{4}, filetype: [1]byte{19}}
	bintxt := fields{datatype: [1]byte{5}}
	xbin := fields{datatype: [1]byte{6}}
	lzh := fields{datatype: [1]byte{7}, filetype: [1]byte{2}}
	exe := fields{datatype: [1]byte{8}}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, ""},
		{"ascii", ascii, "character width"},
		{"html", htm, ""},
		{"pcx", pcx, "pixel width"},
		{"rip script", rip, "pixel width"},
		{"vector", dxf, ""},
		{"smp16s", samp16, "sample rate"},
		{"binary text", bintxt, ""},
		{"xbin", xbin, "character width"},
		{"lzh", lzh, ""},
		{"exe", exe, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
				Tinfo1:   tt.fields.tinfo1,
				Tinfo2:   tt.fields.tinfo2,
				Tinfo3:   tt.fields.tinfo3,
				TFlags:   tt.fields.tFlags,
			}
			if got := d.InfoType(); got.Info1.Info != tt.want {
				t.Errorf("got.Info1.Inf = %q, want %q", got.Info1.Info, tt.want)
				t.Errorf("data.InfoType() = %v", got)
			}
		})
	}
}
