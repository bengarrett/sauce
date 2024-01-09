package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestXBin_String(t *testing.T) {
	tests := []struct {
		name string
		x    layout.XBin
		want string
	}{
		{"out of range", 999, ""},
		{"first and last", 0, "Extended binary text or a XBin file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.String(); got != tt.want {
				t.Errorf("XBin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
