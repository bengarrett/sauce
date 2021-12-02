package record

import "testing"

func TestNone_String(t *testing.T) {
	tests := []struct {
		name string
		n    None
		want string
	}{
		{"out of range", 999, ""},
		{"first and last", 0, "Undefined filetype"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("None.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
