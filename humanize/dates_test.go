package humanize_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bengarrett/sauce/humanize"
)

func midnight2020() time.Time {
	return time.Date(2020, 1, 1, 24, 0, 0, 0, time.UTC)
}

func newYear2020() time.Time {
	return time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
}

func TestFormat(t *testing.T) {
	midnight, nyd := midnight2020(), newYear2020()
	type args struct {
		format humanize.Layout
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// errors
		{"nyd empty", args{"", nyd}, "2020 Jan 1 12:00"},
		{"nyd invalid", args{"???", nyd}, ""},
		{"error", args{"xxx", midnight}, ""},
		// dates
		{"nyd DMY", args{humanize.DMY, nyd}, "1 Jan 2020"},
		{"nyd YMD", args{humanize.YMD, nyd}, "2020 Jan 1"},
		{"nyd MDY", args{humanize.MDY, nyd}, "Jan 1 2020"},
		// times,
		{"nyd 12", args{humanize.H12, nyd}, "12:00 pm"},
		{"nyd 24", args{humanize.H24, nyd}, "12:00"},
		{"midnight 12", args{humanize.H12, midnight}, "12:00 am"},
		{"midnight 24", args{humanize.H24, midnight}, "00:00"},
		// dates+times
		{"dmy 12", args{humanize.DMY12, nyd}, "1 Jan 2020 12:00 pm"},
		{"ymd 12", args{humanize.YMD12, nyd}, "2020 Jan 1 12:00 pm"},
		{"mdy 12", args{humanize.MDY12, nyd}, "Jan 1 2020 12:00 pm"},
		{"dym 24", args{humanize.DMY24, nyd}, "1 Jan 2020 12:00"},
		{"ymd 24", args{humanize.YMD24, midnight}, "2020 Jan 2 00:00"},
		{"mdy 24", args{humanize.MDY24, midnight}, "Jan 2 2020 00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.format.Format(tt.args.t); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleLayout_Format() {
	t := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)

	fmt.Println(humanize.DMY.Format(t))
	fmt.Println(humanize.YMD24.Format(t))
	// Output: 1 Jan 2020
	// 2020 Jan 1 12:00
}
