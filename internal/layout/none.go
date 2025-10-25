package layout

// Undefined filetype.
// You could use this to add SAUCE to a custom or proprietary file,
// without giving it any particular meaning or interpretation.
// See http://www.acid.org/info/sauce/sauce.htm#FileType

type None uint

const (
	Undefined None = iota
)

func (n None) String() string {
	if n > Undefined {
		return ""
	}
	return [...]string{
		"Undefined filetype",
	}[n]
}
