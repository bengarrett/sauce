// This is a raw memory copy of a text mode screen. Also known as a .BIN file.
// This is essentially a collection of character and attribute pairs.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package record

// BinaryText is a raw memory copy of a text mode screen.
type BinaryText uint

const (
	BinaryScreenImage BinaryText = iota
)

func (b BinaryText) String() string {
	if b > BinaryScreenImage {
		return ""
	}
	return [...]string{
		"Binary text or a .BIN file",
	}[b]
}
