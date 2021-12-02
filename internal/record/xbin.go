// An XBin or eXtended BIN file.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package record

// XBin or extended binary text files.
type XBin uint

const (
	ExtendedBin XBin = iota
)

func (x XBin) String() string {
	if x > ExtendedBin {
		return ""
	}
	return [...]string{
		"Extended binary text or a XBin file",
	}[x]
}
