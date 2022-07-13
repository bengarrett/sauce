// Package sauce to handle the reading and parsing of embedded SAUCE metadata.
package sauce_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/bengarrett/sauce"
	"github.com/bengarrett/sauce/internal/layout"
)

const example = "static/sauce.txt"

func TestTrim(t *testing.T) {
	none := []byte("This is a string without any SAUCE.")
	if got := sauce.Trim(none); !reflect.DeepEqual(got, none) {
		t.Errorf("Trim() = %q, want %q", got, none)
	}

	fake := none
	fake = append(fake, layout.SauceSeek...)
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

func TestJSON(t *testing.T) {
	const id, ver = "SAUCE", "00"
	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("JSON() %v error: %v", example, err)
		return
	}
	rec := sauce.Decode(raw)
	// test json
	got, err := rec.JSON()
	if err != nil {
		t.Errorf("JSON() %v error: %v", example, err)
		return
	}
	var res sauce.Record
	if err := json.Unmarshal([]byte(got), &res); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}
	if res.ID != id {
		t.Errorf("Unmarshal ID got: %v, want %v", res.ID, id)
	}
	if res.Version != ver {
		t.Errorf("Unmarshal Version got: %v, want %v", res.Version, ver)
	}
	// test json tab indent
	got, err = rec.JSONIndent("\t")
	if err != nil {
		t.Errorf("JSONIndent() %v error: %v", example, err)
		return
	}
	res = sauce.Record{}
	if err := json.Unmarshal([]byte(got), &res); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}
	if res.ID != id {
		t.Errorf("Unmarshal ID got: %v, want %v", res.ID, id)
	}
	if res.Version != ver {
		t.Errorf("Unmarshal Version got: %v, want %v", res.Version, ver)
	}
}

func TestXML(t *testing.T) {
	const id, ver = "SAUCE", "00"
	raw, err := static.ReadFile(example)
	if err != nil {
		t.Errorf("JSON() %v error: %v", example, err)
		return
	}
	rec := sauce.Decode(raw)
	got, err := rec.XMLIndent("  ")
	if err != nil {
		t.Errorf("XML() %v error: %v", example, err)
		return
	}
	var res sauce.Record
	if err := xml.Unmarshal([]byte(got), &res); err != nil {
		t.Errorf("Unmarshal error: %v", err)
		s := strings.Split(string(got), "\n")
		for i, row := range s {
			fmt.Fprintf(os.Stderr, "%d. %s\n", i+1, row)
		}
		return
	}
	if res.ID != id {
		t.Errorf("Unmarshal ID got: %v, want %v", res.ID, id)
	}
	if res.Version != ver {
		t.Errorf("Unmarshal Version got: %v, want %v", res.Version, ver)
	}
}
