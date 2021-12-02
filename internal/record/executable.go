// A executable file. Any executable file. .exe, .dll, .bat, ...
// Executable scripts such as .vbs should be tagged as Source.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package record

// Executable program files.
type Executable uint

const (
	Exe Executable = iota
)

func (e Executable) String() string {
	if e > Exe {
		return ""
	}
	return [...]string{
		"Executable program file",
	}[e]
}
