package generator

import (
	"math"
	"math/rand"
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// UniformNoiseM returns a mono uniform noise generator
func UniformNoiseM(source source.Source) source.Source {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = (r.Float32() * 2) - 1
			out[1][i] = out[0][i]
		}
		return out
	}
}

// UniformNoiseS returns a stereo uniform noise generator
func UniformNoiseS(source source.Source) source.Source {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = (r.Float32() * 2) - 1
			out[1][i] = (r.Float32() * 2) - 1
		}
		return out
	}
}

// SineM returns a mono sine wave generator
func SineM(source source.Source, frequency, phase, sampleRate float64) source.Source {
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

// SineS returns a stereo sine wave generator
func SineS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
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

// TriangleM returns a mono triangle wave generator
func TriangleM(source source.Source, frequency, phase, sampleRate float64) source.Source {
	step := frequency / sampleRate
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = float32(utils.Triangle(2 * math.Pi * phase))
			_, phase = math.Modf(phase + step)
			out[1][i] = out[0][i]
		}
		return out
	}
}

// TriangleS returns a stereo triangle wave generator
func TriangleS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	stepL := frequencyL / sampleRate
	stepR := frequencyR / sampleRate
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = float32(utils.Triangle(2 * math.Pi * phaseL))
			_, phaseL = math.Modf(phaseL + stepL)
			out[1][i] = float32(utils.Triangle(2 * math.Pi * phaseR))
			_, phaseR = math.Modf(phaseR + stepR)
		}
		return out
	}
}
