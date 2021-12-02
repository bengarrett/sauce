package record_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func Test_Data_Sizes(t *testing.T) {
	tests := []struct {
		name     string
		filesize record.FileSize
		want     record.Sizes
	}{
		{"none", record.FileSize([4]byte{}), record.Sizes{0, "0", "0"}},
		{"1 byte", record.FileSize([4]byte{1}), record.Sizes{1, "1B", "1B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Filesize: tt.filesize,
			}
			if got := d.Sizes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Sizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
