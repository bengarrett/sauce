package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func TestArchive_String(t *testing.T) {
	tests := []struct {
		name string
		a    record.Archive
		want string
	}{
		{"zip", record.Zip, "ZIP compressed archive"},
		{"squeeze", record.Sqz, "Squeeze It compressed archive"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Archive.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
