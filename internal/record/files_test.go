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
	tests := []struct {
		name   string
		fields fields
		want   record.Files
	}{
		{"none", fields{[1]byte{0}, [1]byte{0}},
			record.Files{record.TypeOfFile(record.Nones), record.Nones.String()}},
		{"audio", fields{[1]byte{4}, [1]byte{0}},
			record.Files{record.TypeOfFile(record.Mod), record.Mod.String()}},
		{"executable", fields{[1]byte{8}, [1]byte{0}},
			record.Files{record.TypeOfFile(record.Executables), record.Executables.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
			}
			if got := d.FileType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("data.FileType() = %v, want %v", got, tt.want)
			}
		})
	}
}
