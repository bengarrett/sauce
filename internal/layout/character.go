// A character based file. These files are typically interpreted sequentially.
// Also known as streams.
// See http://www.acid.org/info/sauce/sauce.htm#FileType
package layout

// Character files more commonly referred as text files.
type Character uint

const (
	ASCII Character = iota
	Ansi
	AnsiMation
	RipScript
	PCBoard
	Avatar
	HTML
	Source
	TundraDraw
)

func (c Character) String() string {
	if c > TundraDraw {
		return ""
	}
	return [...]string{
		"ASCII text",
		"ANSI color text",
		"ANSIMation",
		"RIPScript",
		"PCBoard color text",
		"Avatar color text",
		"HTML markup",
		"Programming source code",
		"TundraDraw color text",
	}[c]
}

// Desc is the character description.
func (c Character) Desc() string {
	if c > TundraDraw {
		return ""
	}
	return [...]string{
		"ASCII text file with no formatting codes or color codes.",
		"ANSI text file with coloring codes and cursor positioning.",
		"ANSIMation are ANSI text files that rely on fixed screen sizes.",
		"RIPScript are Remote Imaging Protocol graphics.",
		"PCBoard color codes and macros, and ANSI codes.",
		"Avatar color codes, and ANSi codes.",
		"HTML markup files.",
		"Source code for a programming language.",
		"TundraDraw files, like ANSI, but with a custom palette.",
	}[c]
}

func (d *Layout) Description() string {
	dt, ft := UnsignedBinary1(d.Datatype),
		UnsignedBinary1(d.Filetype)
	c := Character(ft)
	if TypeOfData(dt) != Characters {
		return ""
	}
	switch c {
	case ASCII,
		Ansi,
		AnsiMation,
		RipScript,
		PCBoard,
		Avatar,
		HTML,
		Source,
		TundraDraw:
		return c.Desc()
	}
	return ""
}
