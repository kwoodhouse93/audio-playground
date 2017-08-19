package generator

import (
	"math"

	"github.com/kwoodhouse93/audio-playground/source"
)

// NewSineM returns a mono sine wave generator
func NewSineM(source source.Source, frequency, phase, sampleRate float64) source.Source {
	step := frequency / sampleRate
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			sample := float32(math.Sin(2 * math.Pi * phase))
			_, phase = math.Modf(phase + step)
			out[0][i] = sample
			out[1][i] = sample
		}
		return out
	}
}

// NewSineS returns a stereo sine wave generator
func NewSineS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	stepL := frequencyL / sampleRate
	stepR := frequencyR / sampleRate
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = float32(math.Sin(2 * math.Pi * phaseL))
			_, phaseL = math.Modf(phaseL + stepL)
			out[1][i] = float32(math.Sin(2 * math.Pi * phaseR))
			_, phaseR = math.Modf(phaseR + stepR)
		}
		return out
	}
}
