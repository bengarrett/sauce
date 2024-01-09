package layout_test

import (
	"embed"
	"log"
	"strings"
	"testing"

	"github.com/bengarrett/sauce"
	"github.com/bengarrett/sauce/internal/layout"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

//go:embed static/*
var static embed.FS

func raw() []byte {
	b, err := static.ReadFile(example)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func exampleData() layout.Layout {
	return layout.Data(raw()).Extract()
}

func TestIndex(t *testing.T) {
	const hi = "Hello world!"
	fake := make([]byte, 0, len(hi)+len(layout.SauceSeek))
	fake = append(fake, []byte(hi)...)
	fake = append(fake, []byte(layout.SauceSeek)...)
	okay := fake
	okay = append(okay, []byte(strings.Repeat("?", 150))...)
	tests := []struct {
		name      string
		b         []byte
		wantIndex int
	}{
		{"empty", nil, -1},
		{"none", []byte(hi), -1},
		{"falsepos", fake, -1},
		{"pretend", okay, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := layout.Index(tt.b); gotIndex != tt.wantIndex {
				t.Errorf("Index() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestId_String(t *testing.T) {
	const s = "SAUCE"
	if got := exampleData().ID.String(); got != s {
		t.Errorf("Id.String() = %q, want %q", got, s)
	}
	const v = "00"
	if got := exampleData().Version.String(); got != v {
		t.Errorf("Version.String() = %q, want %q", got, v)
	}
	const st = "Sauce title                        "
	if got := exampleData().Title.String(); got != st {
		t.Errorf("Title.String() = %q, want %q", got, st)
	}
	const sa = "Sauce author        "
	if got := exampleData().Author.String(); got != sa {
		t.Errorf("Author.String() = %q, want %q", got, sa)
	}
	const d = "20161126"
	if got := exampleData().Date.String(); got != d {
		t.Errorf("Date.String() = %q, want %q", got, d)
	}
	const vga = "IBM VGA"
	if got := exampleData().TInfoS.String(); got != vga {
		t.Errorf("TInfoS.String() = %q, want %q", got, vga)
	}
}

func Test_Incomplete(t *testing.T) {
	b, err := static.ReadFile("static/error.txt")
	if err != nil {
		t.Error(err)
	}
	r := sauce.Decode(b)
	const want = "20161126"
	if r.Date.Value != want {
		t.Errorf("Date.Value = %q, want %q", r.Date.Value, want)
	}
	if r.Info.Font != "" {
		t.Errorf("Font = %q, want %q", r.Info.Font, "")
	}
}

func Test_Empty(t *testing.T) {
	b, err := static.ReadFile("static/empty.txt")
	if err != nil {
		t.Error(err)
	}
	r := sauce.Decode(b)
	if r.Date.Value != "" {
		t.Errorf("Date.Value = %q, want %q", r.Date.Value, "")
	}
	if r.Info.Font != "" {
		t.Errorf("Font = %q, want %q", r.Info.Font, "")
	}
}
