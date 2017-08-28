package meter

import (
	"time"

	"github.com/kwoodhouse93/audio-playground/notes"
)

// Meter represents a musical timing in beats and bars
type Meter struct {
	BeatsPerMinute float64
	BeatsPerBar    float64
	BeatValue      notes.Duration
}

// NewCommonTime returns a simple 4/4 meter at the specified tempo
func NewCommonTime(bpm float64) *Meter {
	return &Meter{
		BeatsPerMinute: bpm,
		BeatsPerBar:    4,
		BeatValue:      notes.Quarter,
	}
}

// New returns a new meter with the specified parameters
func New(bpm, beatsPerBar float64, beatValue notes.Duration) *Meter {
	return &Meter{
		BeatsPerMinute: bpm,
		BeatsPerBar:    beatsPerBar,
		BeatValue:      beatValue,
	}
}

// NoteToTime converts a notes.Duration to a time.Duration based on the meter
func (m Meter) NoteToTime(noteVal notes.Duration) time.Duration {
	return time.Duration((float64(noteVal/m.BeatValue) / m.BeatsPerMinute) * float64(time.Minute))
}

// NoteToFreq converts a notes.Duration into a frequency with period equal to
// that note length
func (m Meter) NoteToFreq(noteVal notes.Duration) float64 {
	duration := m.NoteToTime(noteVal)
	return 1 / float64(duration.Seconds())
}
