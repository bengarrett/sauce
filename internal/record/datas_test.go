package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func TestDataType_String(t *testing.T) {
	tests := []struct {
		name string
		d    record.TypeOfData
		want string
	}{
		{"none", record.Nones, "undefined"},
		{"executable", record.Executables, "executable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DataType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
