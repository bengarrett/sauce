package layout_test

import (
	"reflect"
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

const (
	null  = ""
	zero  = "00"
	one   = "01"
	two   = "10"
	three = "11" // SAUCE v00.5: Not currently a valid value.
)

func TestFlags_parse(t *testing.T) {
	var (
		invalid = layout.ErrInvalid.Error()
		blink   = layout.BBit("0").String()
		noBlink = layout.BBit("1").String()
		noPref  = layout.Unsupported
		stretch = layout.ArBit(one).String()
		square  = layout.ArBit(two).String()
		px8     = layout.LsBit(one).String()
		px9     = layout.LsBit(two).String()
	)
	tests := []struct {
		name       string
		f          uint8 // range: 0 through 255.
		wantB      string
		wantLS     string
		wantAR     string
		wantString string
	}{
		{"zero", 0, blink, noPref, noPref, ""},
		{"one", 1, blink, noPref, stretch, "blink mode, stretch pixels"},
		{"two", 2, blink, noPref, square, "blink mode, square pixels"},
		{"three", 3, blink, noPref, invalid, "blink mode, invalid value"},
		{"four", 4, blink, px8, noPref, "blink mode, select 8 pixel font"},
		{"five", 5, blink, px8, stretch, "blink mode, select 8 pixel font, stretch pixels"},
		{"no blink", 99, noBlink, px9, noPref, "non-blink mode, select 9 pixel font"},
		{"max", 255, noBlink, invalid, invalid, "non-blink mode, invalid value, invalid value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := layout.Flags(tt.f).Parse(); !reflect.DeepEqual(got.B.Info, tt.wantB) {
				t.Errorf("Flags.parse() = %v, want %v", got.B.Info, tt.wantB)
			}
			if got := layout.Flags(tt.f).Parse(); !reflect.DeepEqual(got.LS.Info, tt.wantLS) {
				t.Errorf("Flags.parse() = %v, want %v", got.LS.Info, tt.wantLS)
			}
			if got := layout.Flags(tt.f).Parse(); !reflect.DeepEqual(got.AR.Info, tt.wantAR) {
				t.Errorf("Flags.parse() = %v, want %v", got.AR.Info, tt.wantAR)
			}
			if got := layout.Flags(tt.f).Parse(); got.String() != tt.wantString {
				t.Errorf("Flags.String() = %q, want %q", got.String(), tt.wantB)
			}
		})
	}
}

func Test_LsBit_String(t *testing.T) {
	tests := []struct {
		name string
		ls   layout.LsBit
		want string
	}{
		{"empty", null, layout.ErrInvalid.Error()},
		{"00", zero, layout.Unsupported},
		{"8px", one, "select 8 pixel font"},
		{"9px", two, "select 9 pixel font"},
		{"err", three, layout.ErrInvalid.Error()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.String(); got != tt.want {
				t.Errorf("LsBit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ArBit_String(t *testing.T) {
	tests := []struct {
		name string
		ar   layout.ArBit
		want string
	}{
		{"empty", null, layout.ErrInvalid.Error()},
		{"00", zero, layout.Unsupported},
		{"8px", one, "stretch pixels"},
		{"9px", two, "square pixels"},
		{"err", three, layout.ErrInvalid.Error()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ar.String(); got != tt.want {
				t.Errorf("arBit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBBit_String(t *testing.T) {
	tests := []struct {
		name string
		b    layout.BBit
		want string
	}{
		{"empty", null, layout.ErrInvalid.Error()},
		{"8px", "0", "blink mode"},
		{"9px", "1", "non-blink mode"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("BBit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
