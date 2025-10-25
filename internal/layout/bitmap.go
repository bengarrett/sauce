package layout

// Bitmap graphic and animation files.
// See http://www.acid.org/info/sauce/sauce.htm#FileType

// Bitmap graphic and animation files.
type Bitmap uint

const (
	Gif Bitmap = iota
	Pcx
	Lbm
	Tga
	Fli
	Flc
	Bmp
	Gl
	Dl
	Wpg
	Png
	Jpg
	Mpg
	Avi
)

func (b Bitmap) String() string {
	if b > Avi {
		return ""
	}
	return [...]string{
		"GIF image",
		"ZSoft Paintbrush image",
		"DeluxePaint image",
		"Targa true color image",
		"Autodesk Animator animation",
		"Autodesk Animator animation",
		"BMP Windows/OS2 bitmap",
		"Grasp GL animation",
		"DL animation",
		"WordPerfect graphic",
		"PNG image",
		"Jpeg photo",
		"MPEG video",
		"AVI video",
	}[b]
}
