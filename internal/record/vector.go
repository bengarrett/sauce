package record

// Vector graphic files.
type Vector uint

const (
	Dxf Vector = iota
	Dwg
	Wpvg
	Kinetix
)

func (v Vector) String() string {
	return [...]string{
		"AutoDesk CAD vector graphic",
		"AutoDesk CAD vector graphic",
		"WordPerfect vector graphic",
		"3D Studio vector graphic",
	}[v]
}
