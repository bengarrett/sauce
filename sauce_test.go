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

func TestDecode(t *testing.T) {
	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("Decode() %v error: %v", example, err)
	}
	got := sauce.Decode(raw)
	const wantT = "Sauce title"
	if got.Title != wantT {
		t.Errorf("Decode().Title = %q, want %q", got.Title, wantT)
	}
	const wantA = "Sauce author"
	if got.Author != wantA {
		t.Errorf("Decode().Author = %q, want %q", got.Author, wantA)
	}
}
