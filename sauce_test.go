// Package sauce to handle the reading and parsing of embedded SAUCE metadata.
package sauce_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bengarrett/sauce"
	"github.com/bengarrett/sauce/internal/layout"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

func TestTrim(t *testing.T) {
	none := []byte("This is a string without any SAUCE.")
	if got := sauce.Trim(none); !reflect.DeepEqual(got, none) {
		t.Errorf("Trim() = %q, want %q", got, none)
	}

	fake := append(none, layout.SauceSeek...)
	fake = append(fake, []byte(strings.Repeat("?", 128))...)
	if got := sauce.Trim(fake); !reflect.DeepEqual(got, none) {
		t.Errorf("Trim() = %q, want %q", got, none)
	}

	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("Trim() %v error: %v", example, err)
	}
	const wantL = 1251
	if got := sauce.Trim(raw); len(got) != wantL {
		t.Errorf("Trim() length = %d, want %d", len(got), wantL)
		t.Errorf("Trim() = %q", got)
	}
}

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
