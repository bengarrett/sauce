package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestBitmap_String(t *testing.T) {
	tests := []struct {
		name string
		b    layout.Bitmap
		want string
	}{
		{"out of range", 999, ""},
		{"first", layout.Gif, "GIF image"},
		{"last", layout.Avi, "AVI video"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Bitmap.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
