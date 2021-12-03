package humanize_test

import (
	"fmt"
	"testing"

	"github.com/bengarrett/sauce/humanize"
	"golang.org/x/text/language"
)

func Test_binary_decimal(t *testing.T) {
	us := language.AmericanEnglish
	type args struct {
		b int64
		t language.Tag
	}
	tests := []struct {
		name    string
		args    args
		wantBin string
		wantDec string
	}{
		{"0", args{int64(0), us}, "0", "0"},
		{"1", args{int64(1), us}, "1B", "1B"},
		{"1.5K", args{int64(1500), us}, "1.5 KiB", "1.5 kB"},
		{"2.5M", args{int64(2.5e+6), us}, "2.38 MiB", "2.50 MB"},
		{"3.1G", args{int64(3.1e+9), us}, "2.89 GiB", "3.10 GB"},
		{"9.9T", args{int64(9.9e+12), us}, "9.00 TiB", "9.90 TB"},
		{"upper limit", args{int64(1.5e+18), us}, "1,332.27 PiB", "1,500.00 PB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := humanize.Binary(tt.args.b, tt.args.t); got != tt.wantBin {
				t.Errorf("Binary() = %v, want %v", got, tt.wantBin)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := humanize.Decimal(tt.args.b, tt.args.t); got != tt.wantDec {
				t.Errorf("decimal() = %v, want %v", got, tt.wantDec)
			}
		})
	}
}

func ExampleBinary() {
	const size = int64(15724)
	us, de := language.AmericanEnglish, language.German
	fmt.Printf("Filesize is %s.\n", humanize.Binary(size, us))
	fmt.Printf("Dateigröße ist %s.", humanize.Binary(size, de))
	// Output: Filesize is 15.4 KiB.
	// Dateigröße ist 15,4 KiB.
}

func ExampleDecimal() {
	const size = int64(15724)
	us, de := language.AmericanEnglish, language.German
	fmt.Printf("Filesize is %s.\n", humanize.Decimal(size, us))
	fmt.Printf("Dateigröße ist %s.", humanize.Decimal(size, de))
	// Output: Filesize is 15.7 kB.
	// Dateigröße ist 15,7 kB.
}
