package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestVector_String(t *testing.T) {
	tests := []struct {
		name string
		v    layout.Vector
		want string
	}{
		{"out of range", 999, ""},
		{"first", layout.Dxf, "AutoDesk CAD vector graphic"},
		{"last", layout.Kinetix, "3D Studio vector graphic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("Vector.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
