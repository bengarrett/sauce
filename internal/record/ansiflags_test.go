package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func Test_lsBit_String(t *testing.T) {
	tests := []struct {
		name string
		ls   record.LsBit
		want string
	}{
		{"empty", "", record.ErrInvalid.Error()},
		{"00", "00", record.NoPref},
		{"8px", "01", "select 8 pixel font"},
		{"9px", "10", "select 9 pixel font"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.String(); got != tt.want {
				t.Errorf("LsBit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
