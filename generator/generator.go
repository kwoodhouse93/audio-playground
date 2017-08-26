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
	return func() []float32 {
		out := source()
		out[0] = (r.Float32() * 2) - 1
		out[1] = out[0]
		return out
	}
}

// UniformNoiseS returns a stereo uniform noise generator
func UniformNoiseS(source source.Source) source.Source {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() []float32 {
		out := source()
		out[0] = (r.Float32() * 2) - 1
		out[1] = (r.Float32() * 2) - 1
		return out
	}
}

// SineM returns a mono sine wave generator
func SineM(source source.Source, frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(math.Sin, source, frequency, phase, sampleRate)
}

// SineS returns a stereo sine wave generator
func SineS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(math.Sin, source, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// TriangleM returns a mono triangle wave generator
func TriangleM(source source.Source, frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(utils.Triangle, source, frequency, phase, sampleRate)
}

// TriangleS returns a stereo triangle wave generator
func TriangleS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(utils.Triangle, source, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// SawtoothM returns a mono sawtooth wave generator
func SawtoothM(source source.Source, frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(utils.Sawtooth, source, frequency, phase, sampleRate)
}

// SawtoothS returns a stereo sawtooth wave generator
func SawtoothS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(utils.Sawtooth, source, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// SquareM returns a mono square wave generator
func SquareM(source source.Source, frequency, phase, dutyCycle, sampleRate float64) source.Source {
	return applyWaveM(utils.SquareFunc(dutyCycle), source, frequency, phase, sampleRate)
}

// SquareS returns a stereo square wave generator
func SquareS(source source.Source, frequencyL, frequencyR, phaseL, phaseR, dutyCycle, sampleRate float64) source.Source {
	return applyWaveS(utils.SquareFunc(dutyCycle), source, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

func applyWaveM(waveFunc func(float64) float64, source source.Source, frequency, phase, sampleRate float64) source.Source {
	step := frequency / sampleRate
	return func() []float32 {
		out := source()
		out[0] = float32(waveFunc(2 * math.Pi * phase))
		_, phase = math.Modf(phase + step)
		out[1] = out[0]
		return out
	}
}

func applyWaveS(waveFunc func(float64) float64, source source.Source, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	stepL := frequencyL / sampleRate
	stepR := frequencyR / sampleRate
	return func() []float32 {
		out := source()
		out[0] = float32(waveFunc(2 * math.Pi * phaseL))
		_, phaseL = math.Modf(phaseL + stepL)
		out[1] = float32(waveFunc(2 * math.Pi * phaseR))
		_, phaseR = math.Modf(phaseR + stepR)
		return out
	}
}
