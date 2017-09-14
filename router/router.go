package router

import (
	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/types"
)

// Gain returns a filter that multiplies the signals in the buffer by a constant value
func Gain(src source.Source, gain float32) source.Source {
	return func(step int) types.Sample {
		in := src(step)
		return in.Gain(gain)
	}
}

// Sum2 returns a filter that adds two signals together
// Be careful using this without first reducing the gain of the sources
func Sum2(src1, src2 source.Source) source.Source {
	return func(step int) types.Sample {
		s1 := src1(step)
		s2 := src2(step)
		return s1.Sum(s2)
	}
}

// Sum2Comp returns a function that adds two signals together and reduces
// their volume to compensate for the addition
func Sum2Comp(src1, src2 source.Source) source.Source {
	return func(step int) types.Sample {
		s1 := src1(step)
		s2 := src2(step)
		return s1.Sum(s2).Gain(0.5)
	}
}

// Sum returns a filter that adds multiple signals together
// Failing to reduce the gain of these signals will likely cause severe clipping
func Sum(srcs ...source.Source) source.Source {
	out := srcs[0]
	for _, src := range srcs[1:] {
		out = Sum2(out, src)
	}
	return out
}

// SumComp returns a filter that adds multiple signals but compensates for
// the increase in volume that this would result in
func SumComp(srcs ...source.Source) source.Source {
	compGain := 1 / (float32(len(srcs)))
	out := srcs[0]
	for _, src := range srcs[1:] {
		out = Sum2(Gain(out, compGain), Gain(src, compGain))
	}
	return out
}

// Mixer2 returns a filter that mixes two signals together
func Mixer2(src1, src2 source.Source, gain1, gain2 float32) source.Source {
	return Sum2(
		Gain(src1, gain1),
		Gain(src2, gain2),
	)
}

//SourceGain represents a source and gain pair to use in a Mixer
type SourceGain struct {
	Source source.Source
	Gain   float32
}

// Mixer returns a filter that mixes any number of signals together
func Mixer(inputs []SourceGain) source.Source {
	sources := make([]source.Source, len(inputs))
	for i, input := range inputs {
		sources[i] = Gain(input.Source, input.Gain)
	}
	return Sum(sources...)
}
