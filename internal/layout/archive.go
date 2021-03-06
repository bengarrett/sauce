// An archive file data type.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package layout

// Archive and compressed files.
type Archive uint

const (
	Zip Archive = iota
	Arj
	Lzh
	Arc
	Tar
	Zoo
	Rar
	Uc2
	Pak
	Sqz
)

func (a Archive) String() string {
	if (a) > (Sqz) {
		return ""
	}
	return [...]string{
		"ZIP compressed archive",
		"ARJ compressed archive",
		"LHA compressed archive",
		"ARC compressed archive",
		"Tarball tape archive",
		"ZOO compressed archive",
		"RAR compressed archive",
		"UltraCompressor II compressed archive",
		"PAK ARC compressed archive",
		"Squeeze It compressed archive",
	}[a]
}
