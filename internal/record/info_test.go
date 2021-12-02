package record_test

import (
	"reflect"
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, ""},
		{"ascii", fields{datatype: [1]byte{1}, filetype: [1]byte{0}}, "character width"},
		{"rip script", fields{datatype: [1]byte{1}, filetype: [1]byte{3}}, "pixel width"},
		{"smp16s", fields{datatype: [1]byte{4}, filetype: [1]byte{19}}, "sample rate"},
		{"binary text", fields{datatype: [1]byte{5}}, ""},
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
			if got := d.InfoType(); !reflect.DeepEqual(got.Info1.Info, tt.want) {
				t.Errorf("data.InfoType() = %v, want %v", got, tt.want)
			}
		})
	}
}
