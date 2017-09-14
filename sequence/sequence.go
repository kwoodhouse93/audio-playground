package sequence

import (
	"math"
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/types"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// Pulse returns an object that outputs a pulse at value 1.0 for a specified
// duration after its control source goes high
func Pulse(control source.Source, duration time.Duration, threshold, sampleRate float64) source.Source {
	var (
		active     = false
		pulseSteps = utils.TimeToSteps(duration, sampleRate)
		curStep    int
	)
	return source.Cached(func(step int) types.Sample {
		out := types.NewSample(2)
		ctl := control(step)
		if active {
			out.Set(1)
			curStep--
			if curStep == 0 {
				active = false
			}
			return out
		}
		if float64(ctl[0]) > threshold {
			active = true
			curStep = pulseSteps
		}
		out.Set(0)
		return out
	})
}

// Gate allows the input signal to pass only when the control signal is above
// a certain threshold
func Gate(source, control source.Source, threshold float64) source.Source {
	return func(step int) types.Sample {
		ctl := control(step)
		input := source(step)
		if math.Abs(float64(ctl[0])) > threshold {
			return input
		}
		return input.Set(0)
	}
}

// Sequencer outputs from each input source in turn, for a fixed duration
func Sequencer(sources []source.Source, period time.Duration, sampleRate float64) source.Source {
	var (
		channel  = 0
		seqSteps = utils.TimeToSteps(period, sampleRate)
		curStep  = seqSteps
	)
	return source.Cached(func(step int) types.Sample {
		samples := make([]types.Sample, len(sources))
		for ch, source := range sources {
			samples[ch] = source(step)
		}
		curStep--
		if curStep <= 0 {
			channel = (channel + 1) % len(sources)
			curStep = seqSteps
		}
		return samples[channel]
	})
}
