package filter

import (
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// Zero sets a signal to 0
func Zero(source source.Source) source.Source {
	return func() []float32 {
		out := source()
		out[0] = 0
		out[1] = 0
		return out
	}
}

// Delay stores the samples for a given duration then plays them back delayed
func Delay(source source.Source, delay time.Duration, sampleRate float64) source.Source {
	delaySamples := utils.TimeToSteps(delay, sampleRate)
	delayBuf := utils.MakeBuffer(2, delaySamples)
	return func() []float32 {
		out := utils.MakeSample(2)
		// Pop from front and shift buffer
		out[0], delayBuf[0] = delayBuf[0][0], delayBuf[0][1:]
		out[1], delayBuf[1] = delayBuf[1][0], delayBuf[1][1:]
		// Evaluate the sample on the input
		s := source()
		// Push the input sample to the end of the buffer
		delayBuf[0] = append(delayBuf[0], s[0])
		delayBuf[1] = append(delayBuf[1], s[1])
		// Return popped sample
		return out
	}
}

// // LowPass is a simple LPF
// func LowPass(source source.Source) source.Source {
// }
