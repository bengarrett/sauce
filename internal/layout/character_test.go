package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestCharacter_String(t *testing.T) {
	tests := []struct {
		name string
		c    layout.Character
		want string
	}{
		{"out of range", 999, ""},
		{"first", layout.Ascii, "ASCII text"},
		{"last", layout.TundraDraw, "TundraDraw color text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Character.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_Desc(t *testing.T) {
	tests := []struct {
		name string
		c    layout.Character
		want string
	}{
		{"out of range", 999, ""},
		{"first", layout.Ascii,
			"ASCII text file with no formatting codes or color codes."},
		{"last", layout.TundraDraw,
			"TundraDraw files, like ANSI, but with a custom palette."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Desc(); got != tt.want {
				t.Errorf("Character.Desc() = %q", got)
			}
		})
	}
}

func Test_data_description(t *testing.T) {
	type fields struct {
		datatype layout.DataType
		filetype layout.FileType
	}
	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{"out of range", fields{[1]byte{255}, [1]byte{255}}, ""},
		{"none", fields{[1]byte{0}, [1]byte{0}}, ""},
		{"pc board", fields{[1]byte{1}, [1]byte{4}}, layout.PCBoard.Desc()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &layout.Layout{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
			}
			if gotS := d.Description(); gotS != tt.wantS {
				t.Errorf("data.Description() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
