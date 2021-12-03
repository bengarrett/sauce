package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestArchive_String(t *testing.T) {
	tests := []struct {
		name string
		a    layout.Archive
		want string
	}{
		{"out of range", 999, ""},
		{"(first) zip", layout.Zip, "ZIP compressed archive"},
		{"(last) squeeze", layout.Sqz, "Squeeze It compressed archive"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Archive.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
