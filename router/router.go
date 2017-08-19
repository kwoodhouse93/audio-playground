package router

import (
	"github.com/kwoodhouse93/audio-playground/source"
)

// Gain returns a filter that multiplies the signals in the buffer by a constant value
func Gain(source source.Source, gain float32) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = out[0][i] * gain
			out[1][i] = out[1][i] * gain
		}
		return out
	}
}

// Sum2 returns a filter that adds two signals together
// Be careful using this without first reducing the gain of the sources
func Sum2(source1, source2 source.Source) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = source1(bufferSize)
		mix := source2(bufferSize)
		for i := range out[0] {
			out[0][i] += mix[0][i]
			out[1][i] += mix[1][i]
		}
		return out
	}
}

// Sum2Comp returns a function that adds two signals together and reduces
// their volume to compensate for the addition
func Sum2Comp(source1, source2 source.Source) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = source1(bufferSize)
		mix := source2(bufferSize)
		for i := range out[0] {
			out[0][i] = (out[0][i] + mix[0][i]) / 2
			out[1][i] = (out[1][i] + mix[1][i]) / 2
		}
		return out
	}
}

// Sum returns a filter that adds multiple signals together
// Failing to reduce the gain of these signals will likely cause severe clipping
func Sum(sources ...source.Source) source.Source {
	return func(bufferSize int) (out [][]float32) {
		curSource := sources[0]
		for _, source := range sources[1:] {
			curSource = Sum2(curSource, source)
		}
		return curSource(bufferSize)
	}
}

// SumComp returns a filter that adds multiple signals but compensates for
// the increase in volume that this would result in
func SumComp(sources ...source.Source) source.Source {
	return func(bufferSize int) (out [][]float32) {
		compGain := 1 / (float32(len(sources)))
		curSource := sources[0]
		for _, source := range sources[1:] {
			curSource = Sum2(Gain(curSource, compGain), Gain(source, compGain))
		}
		return curSource(bufferSize)
	}
}

// Mixer2 returns a filter that mixes two signals together
func Mixer2(source1, source2 source.Source, gain1, gain2 float32) source.Source {
	return func(bufferSize int) (out [][]float32) {
		return Sum2(
			Gain(source1, gain1),
			Gain(source2, gain2),
		)(bufferSize)
	}
}

//SourceGain represents a source and gain pair to use in a Mixer
type SourceGain struct {
	Source source.Source
	Gain   float32
}

// Mixer returns a filter that mixes any number of signals together
func Mixer(inputs []SourceGain) source.Source {
	return func(bufferSize int) (out [][]float32) {
		sources := make([]source.Source, len(inputs))
		for i, input := range inputs {
			sources[i] = Gain(input.Source, input.Gain)
		}
		return Sum(sources...)(bufferSize)
	}
}
