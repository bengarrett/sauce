// This is a raw memory copy of a text mode screen. Also known as a .BIN file.
// This is essentially a collection of character and attribute pairs.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestBinaryText_String(t *testing.T) {
	tests := []struct {
		name string
		b    layout.BinaryText
		want string
	}{
		{"out of range", 999, ""},
		{"first and last", 0, "Binary text or a .BIN file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("BinaryText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
