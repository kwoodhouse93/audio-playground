package generator

import (
	"math"
	"math/rand"
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/types"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// UniformNoiseM returns a mono uniform noise generator
func UniformNoiseM() source.Source {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return source.Cached(func(step int) types.Sample {
		out := types.NewSample(2)
		out[0] = (r.Float32() * 2) - 1
		out[1] = out[0]
		return out
	})
}

// UniformNoiseS returns a stereo uniform noise generator
func UniformNoiseS() source.Source {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return source.Cached(func(step int) types.Sample {
		out := types.NewSample(2)
		out[0] = (r.Float32() * 2) - 1
		out[1] = (r.Float32() * 2) - 1
		return out
	})
}

// SineM returns a mono sine wave generator
func SineM(frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(math.Sin, frequency, phase, sampleRate)
}

// SineS returns a stereo sine wave generator
func SineS(frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(math.Sin, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// TriangleM returns a mono triangle wave generator
func TriangleM(frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(utils.Triangle, frequency, phase, sampleRate)
}

// TriangleS returns a stereo triangle wave generator
func TriangleS(frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(utils.Triangle, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// SawtoothM returns a mono sawtooth wave generator
func SawtoothM(frequency, phase, sampleRate float64) source.Source {
	return applyWaveM(utils.Sawtooth, frequency, phase, sampleRate)
}

// SawtoothS returns a stereo sawtooth wave generator
func SawtoothS(frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	return applyWaveS(utils.Sawtooth, frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

// SquareM returns a mono square wave generator
func SquareM(frequency, phase, dutyCycle, sampleRate float64) source.Source {
	return applyWaveM(utils.SquareFunc(dutyCycle), frequency, phase, sampleRate)
}

// SquareS returns a stereo square wave generator
func SquareS(frequencyL, frequencyR, phaseL, phaseR, dutyCycle, sampleRate float64) source.Source {
	return applyWaveS(utils.SquareFunc(dutyCycle), frequencyL, frequencyR, phaseL, phaseR, sampleRate)
}

func applyWaveM(waveFunc func(float64) float64, frequency, phase, sampleRate float64) source.Source {
	stepChange := frequency / sampleRate
	return source.Cached(func(step int) types.Sample {
		out := types.NewSample(2)
		out[0] = float32(waveFunc(2 * math.Pi * phase))
		_, phase = math.Modf(phase + stepChange)
		out[1] = out[0]
		return out
	})
}

func applyWaveS(waveFunc func(float64) float64, frequencyL, frequencyR, phaseL, phaseR, sampleRate float64) source.Source {
	stepChangeL := frequencyL / sampleRate
	stepChangeR := frequencyR / sampleRate
	return source.Cached(func(step int) types.Sample {
		out := types.NewSample(2)
		out[0] = float32(waveFunc(2 * math.Pi * phaseL))
		_, phaseL = math.Modf(phaseL + stepChangeL)
		out[1] = float32(waveFunc(2 * math.Pi * phaseR))
		_, phaseR = math.Modf(phaseR + stepChangeR)
		return out
	})
}
