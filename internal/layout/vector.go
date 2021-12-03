// A vector graphic file.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package layout

// Vector graphic files.
type Vector uint

const (
	Dxf Vector = iota
	Dwg
	Wpvg
	Kinetix
)

func (v Vector) String() string {
	if v > Kinetix {
		return ""
	}
	return [...]string{
		"AutoDesk CAD vector graphic",
		"AutoDesk CAD vector graphic",
		"WordPerfect vector graphic",
		"3D Studio vector graphic",
	}[v]
}
