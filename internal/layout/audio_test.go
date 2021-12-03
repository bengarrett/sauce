package layout_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/layout"
)

func TestAudio_String(t *testing.T) {
	tests := []struct {
		name string
		a    layout.Audio
		want string
	}{
		{"out of range", 999, ""},
		{"first", layout.Mod, "NoiseTracker module"},
		{"midi", layout.Midi, "MIDI audio"},
		{"okt", layout.Okt, "Oktalyzer module"},
		{"last", layout.It, "Impulse Tracker module"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Audio.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
