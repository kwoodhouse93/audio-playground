package router

import (
	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// Gain returns a filter that multiplies the signals in the buffer by a constant value
func Gain(source source.Source, gain float32) source.Source {
	return func(step int) []float32 {
		out := source(step)
		out[0] = out[0] * gain
		out[1] = out[1] * gain
		return out
	}
}

// Sum2 returns a filter that adds two signals together
// Be careful using this without first reducing the gain of the sources
func Sum2(source1, source2 source.Source) source.Source {
	return func(step int) []float32 {
		out := utils.MakeSample(2)
		s1 := source1(step)
		s2 := source2(step)
		out[0] = s1[0] + s2[0]
		out[1] = s1[1] + s2[1]
		return out
	}
}

// Sum2Comp returns a function that adds two signals together and reduces
// their volume to compensate for the addition
func Sum2Comp(source1, source2 source.Source) source.Source {
	return func(step int) []float32 {
		out := utils.MakeSample(2)
		s1 := source1(step)
		s2 := source2(step)
		out[0] = (s1[0] + s2[0]) / 2
		out[1] = (s1[1] + s2[1]) / 2
		return out
	}
}

// Sum returns a filter that adds multiple signals together
// Failing to reduce the gain of these signals will likely cause severe clipping
func Sum(sources ...source.Source) source.Source {
	return func(step int) []float32 {
		out := sources[0]
		for _, source := range sources[1:] {
			out = Sum2(out, source)
		}
		return out(step)
	}
}

// SumComp returns a filter that adds multiple signals but compensates for
// the increase in volume that this would result in
func SumComp(sources ...source.Source) source.Source {
	return func(step int) []float32 {
		compGain := 1 / (float32(len(sources)))
		out := sources[0]
		for _, source := range sources[1:] {
			out = Sum2(Gain(out, compGain), Gain(source, compGain))
		}
		return out(step)
	}
}

// Mixer2 returns a filter that mixes two signals together
func Mixer2(source1, source2 source.Source, gain1, gain2 float32) source.Source {
	return func(step int) []float32 {
		return Sum2(
			Gain(source1, gain1),
			Gain(source2, gain2),
		)(step)
	}
}

//SourceGain represents a source and gain pair to use in a Mixer
type SourceGain struct {
	Source source.Source
	Gain   float32
}

// Mixer returns a filter that mixes any number of signals together
func Mixer(inputs []SourceGain) source.Source {
	return func(step int) []float32 {
		sources := make([]source.Source, len(inputs))
		for i, input := range inputs {
			sources[i] = Gain(input.Source, input.Gain)
		}
		return Sum(sources...)(step)
	}
}
