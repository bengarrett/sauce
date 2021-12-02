package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func TestCharacter_String(t *testing.T) {
	tests := []struct {
		name string
		c    record.CharacterBase
		want string
	}{
		{"first", record.Ascii, "ASCII text"},
		{"last", record.TundraDraw, "TundraDraw color text"},
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
		c    record.CharacterBase
	}{
		{"first", record.Ascii},
		{"last", record.TundraDraw},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Desc(); got == "" {
				t.Errorf("Character.Desc() = %q", got)
			}
		})
	}
}

func Test_data_description(t *testing.T) {
	type fields struct {
		datatype record.DataType
		filetype record.FileType
	}
	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{"none", fields{[1]byte{0}, [1]byte{0}}, ""},
		{"pc board", fields{[1]byte{1}, [1]byte{4}}, record.PCBoard.Desc()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &record.Data{
				Datatype: tt.fields.datatype,
				Filetype: tt.fields.filetype,
			}
			if gotS := d.Description(); gotS != tt.wantS {
				t.Errorf("data.Description() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
