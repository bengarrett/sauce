package layout

// An audio file data type.
// See http://www.acid.org/info/sauce/sauce.htm#FileType

// Audio or music files.
type Audio uint

const (
	Mod Audio = iota
	Composer669
	Stm
	S3m
	Mtm
	Far
	Ult
	Amf
	Dmf
	Okt
	Rol
	Cmf
	Midi
	Sadt
	Voc
	Wave
	Smp8
	Smp8s
	Smp16
	Smp16s
	Patch8
	Patch16
	Xm
	Hsc
	It
)

func (a Audio) String() string {
	if a > It {
		return ""
	}
	return [...]string{
		"NoiseTracker module",
		"Composer 669 module",
		"ScreamTracker module",
		"ScreamTracker 3 module",
		"MultiTracker module",
		"Farandole Composer module",
		"Ultra Tracker module",
		"Dual Module Player module",
		"X-Tracker module",
		"Oktalyzer module",
		"AdLib Visual Composer FM audio",
		"Creative Music FM audio",
		"MIDI audio",
		"SAdT composer FM audio",
		"Creative Voice File",
		"Waveform audio",
		"single channel 8-bit sample",
		"stereo 8-bit sample",
		"single channel 16-bit sample",
		"stereo 16-bit sample",
		"8-bit patch file",
		"16-bit patch file",
		"Extended Module",
		"Hannes Seifert Composition FM audio",
		"Impulse Tracker module",
	}[a]
}

func (ti *Infos) audio(ft uint8) {
	switch Audio(ft) {
	case Smp8, Smp8s, Smp16, Smp16s:
		ti.Info1.Info = "sample rate"
	case Mod, Composer669, Stm, S3m, Mtm, Far, Ult, Amf, Dmf, Okt,
		Rol, Cmf, Midi, Sadt, Voc, Wave, Patch8, Patch16, Xm, Hsc, It:
		return
	}
}
