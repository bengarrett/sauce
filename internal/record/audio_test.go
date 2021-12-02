package record_test

import (
	"testing"

	"github.com/bengarrett/sauce/internal/record"
)

func TestAudio_String(t *testing.T) {
	tests := []struct {
		name string
		a    record.Audio
		want string
	}{
		{"out of range", 999, ""},
		{"first", record.Mod, "NoiseTracker module"},
		{"midi", record.Midi, "MIDI audio"},
		{"okt", record.Okt, "Oktalyzer module"},
		{"last", record.It, "Impulse Tracker module"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Audio.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
