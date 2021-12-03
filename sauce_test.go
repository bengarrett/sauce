// Package sauce to handle the reading and parsing of embedded SAUCE metadata.
package sauce_test

import (
	"testing"

	"github.com/bengarrett/sauce"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

func TestParse(t *testing.T) {
	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("Decode() %v error: %v", example, err)
	}
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{"empty", []byte(""), ""},
		{"example", raw, "Sauce title"},
	}
	// TODO: {"example", layout.Data(raw()), sauceIndex(), "Sauce author        "},
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sauce.Decode(tt.data); got.Title != tt.want {
				t.Errorf("Decode() = %v, want %v", got.Title, tt.want)
			}
		})
	}
}
