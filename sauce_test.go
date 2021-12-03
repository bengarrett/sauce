// Package sauce to handle the reading and parsing of embedded SAUCE metadata.
package sauce_test

import (
	"log"
	"testing"

	"github.com/bengarrett/sauce"
	"github.com/bengarrett/sauce/internal/record"
)

const (
	commentResult = "Any comments go here.                                           "
	example       = "static/sauce.txt"
)

func raw() []byte {
	b, err := static.ReadFile(example)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func exampleData() record.Layout {
	return record.Record(raw()).Extract()
}

func sauceIndex() int {
	return record.Scan(raw()...)
}

func TestParse(t *testing.T) {
	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("Parse() %v error: %v", example, err)
	}
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{"empty", []byte(""), ""},
		{"example", raw, "Sauce title"},
	}
	// TODO: {"example", record.Record(raw()), sauceIndex(), "Sauce author        "},
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sauce.Parse(tt.data...); got.Title != tt.want {
				t.Errorf("Parse() = %v, want %v", got.Title, tt.want)
			}
		})
	}
}
