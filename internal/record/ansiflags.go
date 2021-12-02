package record

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalid = errors.New("invalid value") // todo: handle this correctly
)

const NoPref = "no preference" // todo: make into an error that can be handled using errors.Is()

// Flags is the SAUCE Flags field.
type Flags uint8

// ANSIFlags are the interpretation of the SAUCE Flags field.
type ANSIFlags struct {
	Decimal         Flags      `json:"decimal" xml:"decimal,attr"`
	Binary          string     `json:"binary" xml:"binary,attr"`
	B               ANSIFlagB  `json:"nonBlinkMode" xml:"non_blink_mode"`
	LS              ANSIFlagLS `json:"letterSpacing" xml:"letter_spacing"`
	AR              ANSIFlagAR `json:"aspectRatio" xml:"aspect_ratio"`
	Interpretations string     `json:"-" xml:"-"`
}

func (a *ANSIFlags) String() string {
	if a.Decimal == 0 {
		return ""
	}
	b, ls, ar := a.B.Info, a.LS.Info, a.AR.Info
	l := []string{}
	if b != NoPref {
		l = append(l, b)
	}
	if ls != NoPref {
		l = append(l, ls)
	}
	if ar != NoPref {
		l = append(l, ar)
	}
	if strings.TrimSpace(strings.Join(l, "")) == "" {
		return ""
	}
	return strings.Join(l, ", ")
}

// ANSIFlagB is the interpretation of the SAUCE Flags non-blink mode binary bit.
type ANSIFlagB struct {
	Flag bBit   `json:"flag" xml:"flag"`
	Info string `json:"interpretation" xml:"interpretation,attr"`
}

func (f Flags) parse() ANSIFlags {
	const binary5Bits, minLen = "%05b", 6
	bin := fmt.Sprintf(binary5Bits, f)
	r := []rune(bin)
	if len(r) < minLen {
		return ANSIFlags{
			Decimal: f,
			Binary:  bin,
		}
	}
	b, ls, ar := string(r[0]), string(r[1:3]), string(r[3:5])
	return ANSIFlags{
		Decimal: f,
		Binary:  bin,
		B:       ANSIFlagB{Flag: bBit(b), Info: bBit(b).String()},
		LS:      ANSIFlagLS{Flag: LsBit(ls), Info: LsBit(ls).String()},
		AR:      ANSIFlagAR{Flag: arBit(ar), Info: arBit(ar).String()},
	}
}

// ANSIFlagLS is the interpretation of the SAUCE Flags letter spacing binary bits.
type ANSIFlagLS struct {
	Flag LsBit  `json:"flag" xml:"flag"`
	Info string `json:"interpretation" xml:"interpretation,attr"`
}

type LsBit string

func (ls LsBit) String() string {
	const none, eight, nine = "00", "01", "10"
	switch ls {
	case none:
		return NoPref
	case eight:
		return "select 8 pixel font"
	case nine:
		return "select 9 pixel font"
	default:
		return ErrInvalid.Error()
	}
}

// ANSIFlagAR is the interpretation of the SAUCE Flags aspect ratio binary bits.
type ANSIFlagAR struct {
	Flag arBit  `json:"flag" xml:"flag"`
	Info string `json:"interpretation" xml:"interpretation,attr"`
}

type arBit string

func (ar arBit) String() string {
	const none, strect, square = "00", "01", "10"
	switch ar {
	case none:
		return NoPref
	case strect:
		return "stretch pixels"
	case square:
		return "square pixels"
	default:
		return ErrInvalid.Error()
	}
}

type bBit string

func (b bBit) String() string {
	const blink, non = "0", "1"
	switch b {
	case blink:
		return "blink mode"
	case non:
		return "non-blink mode"
	default:
		return ErrInvalid.Error()
	}
}
