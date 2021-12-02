package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func TestVector_String(t *testing.T) {
	tests := []struct {
		name string
		v    record.Vector
		want string
	}{
		{"first", record.Dxf, "AutoDesk CAD vector graphic"},
		{"last", record.Kinetix, "3D Studio vector graphic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("Vector.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
