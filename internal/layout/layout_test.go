package layout_test

import (
	"embed"
	"log"
	"strings"
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

var (
	//go:embed static/*
	static embed.FS
)

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

func TestScan(t *testing.T) {
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
			if gotIndex := layout.Scan(tt.b...); gotIndex != tt.wantIndex {
				t.Errorf("Scan() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestId_String(t *testing.T) {
	const s = "SAUCE"
	if got := exampleData().Id.String(); got != s {
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
