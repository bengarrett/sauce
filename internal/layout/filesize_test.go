package layout_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func Test_Data_Sizes(t *testing.T) {
	tests := []struct {
		name     string
		filesize layout.FileSize
		want     layout.Sizes
	}{
		{"none", layout.FileSize([4]byte{}), layout.Sizes{0, "0", "0"}},
		{"1 byte", layout.FileSize([4]byte{1}), layout.Sizes{1, "1B", "1B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &layout.Layout{
				Filesize: tt.filesize,
			}
			if got := d.Sizes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Layout.Sizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
