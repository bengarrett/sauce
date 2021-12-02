package record

// CharacterBase files more commonly referred as text files.
type CharacterBase uint

const (
	Ascii CharacterBase = iota
	Ansi
	AnsiMation
	RipScript
	PCBoard
	Avatar
	Html
	Source
	TundraDraw
)

func (c CharacterBase) String() string {
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
func (c CharacterBase) Desc() string {
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
func (d *Data) Description() (s string) {
	dt, ft := unsignedBinary1(d.Datatype), unsignedBinary1(d.Filetype)
	c := CharacterBase(ft)
	if TypeOfData(dt) != Characters {
		return s
	}
	switch c {
	case Ascii,
		Ansi,
		AnsiMation,
		RipScript,
		PCBoard,
		Avatar,
		Html,
		Source,
		TundraDraw:
		return c.Desc()
	}
	return s
}
