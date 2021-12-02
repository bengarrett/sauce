package record

// BinaryText is a raw memory copy of a text mode screen.
type BinaryText uint

func (b BinaryText) String() string {
	return "Binary text or a .BIN file"
}
