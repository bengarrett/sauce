package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestDataType_String(t *testing.T) {
	tests := []struct {
		name string
		d    layout.TypeOfData
		want string
	}{
		{"out of range", 999, ""},
		{"none", layout.Nones, "undefined"},
		{"executable", layout.Executables, "executable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DataType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
