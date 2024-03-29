package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestExecutable_String(t *testing.T) {
	tests := []struct {
		name string
		e    layout.Executable
		want string
	}{
		{"out of range", 999, ""},
		{"first and last", 0, "Executable program file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Executable.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
