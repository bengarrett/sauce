// ANSiFlags allow an author of ANSi and similar files to provide a clue to a viewer / editor
// how to render the image. The 8 bits in the ANSiFlags contain the following information:
// 0 	0 	0 	A 	R 	L 	S 	B
// B: Non-blink mode (iCE Color).
// LS: Letter-spacing (a.k.a. 8/9 pixel font selection).
// AR: Aspect Ratio.
// See http://www.acid.org/info/sauce/sauce.htm#ANSiFlags.
package layout

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalid = errors.New("invalid value")

// Unsupported is a legacy value used by Letter-spacing and Aspect Ratio.
// It acts as an unsupported placeholder for SAUCE versions prior to v00.5 from Nov 2013.
const Unsupported = "no preference"

// Flags is the SAUCE Flags field.
type Flags uint8

// ANSiFlags allow an author of ANSi and similar files to provide a clue to a viewer/editor how to render the image.
type ANSIFlags struct {
	Decimal         Flags      `json:"decimal" xml:"decimal,attr"`         // Flags value as an unsigned integer.
	Binary          string     `json:"binary" xml:"binary,attr"`           // Flags value in binary notation.
	B               ANSIFlagB  `json:"nonBlinkMode" xml:"non_blink_mode"`  // Non-blink mode (iCE Color) value.
	LS              ANSIFlagLS `json:"letterSpacing" xml:"letter_spacing"` // Letter-spacing value.
	AR              ANSIFlagAR `json:"aspectRatio" xml:"aspect_ratio"`     // Aspect Ratio value.
	Interpretations string     `json:"-" xml:"-"`                          // Humanized descriptions of the ANSIFlags bits.
}

func (a *ANSIFlags) String() string {
	if a.Decimal == 0 {
		return ""
	}
	b, ls, ar := a.B.Info, a.LS.Info, a.AR.Info
	l := []string{}
	if b != Unsupported {
		l = append(l, b)
	}
	if ls != Unsupported {
		l = append(l, ls)
	}
	if ar != Unsupported {
		l = append(l, ar)
	}
	if strings.TrimSpace(strings.Join(l, "")) == "" {
		return ""
	}
	return strings.Join(l, ", ")
}

// ANSIFlagB is the interpretation of the SAUCE Flags non-blink mode binary bit.
type ANSIFlagB struct {
	Flag BBit   `json:"flag" xml:"flag"`                          // Non-blink mode (iCE Color) toggle.
	Info string `json:"interpretation" xml:"interpretation,attr"` // A description of the toggle.
}

func (f Flags) Parse() ANSIFlags {
	const binary5Bits, minLen = "%05b", 5
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
		B:       ANSIFlagB{Flag: BBit(b), Info: BBit(b).String()},
		LS:      ANSIFlagLS{Flag: LsBit(ls), Info: LsBit(ls).String()},
		AR:      ANSIFlagAR{Flag: ArBit(ar), Info: ArBit(ar).String()},
	}
}

// ANSIFlagLS is the interpretation of the SAUCE Flags letter spacing binary bits.
type ANSIFlagLS struct {
	Flag LsBit  `json:"flag" xml:"flag"`                          // Letter-spacing value.
	Info string `json:"interpretation" xml:"interpretation,attr"` // A description of the value.
}

type LsBit string // Letter-spacing two bit value.

func (ls LsBit) String() string {
	const none, eight, nine = "00", "01", "10"
	switch ls {
	case none:
		return Unsupported
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
	Flag ArBit  `json:"flag" xml:"flag"`                          // Aspect Ratio value.
	Info string `json:"interpretation" xml:"interpretation,attr"` // A description of the value.
}

type ArBit string // Aspect Ratio two bit value.

func (ar ArBit) String() string {
	const none, strect, square = "00", "01", "10"
	switch ar {
	case none:
		return Unsupported
	case strect:
		return "stretch pixels"
	case square:
		return "square pixels"
	default:
		return ErrInvalid.Error()
	}
}

type BBit string // Non-blink mode (iCE Color) bit value.

func (b BBit) String() string {
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
