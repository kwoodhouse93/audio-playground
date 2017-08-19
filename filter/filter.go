package filter

import (
	"math"
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
)

// Pass returns a filter that doesn't modify the source
func Pass(source source.Source) source.Source {
	return source
}

// Zero sets a signal to 0
func Zero(source source.Source) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = 0
			out[1][i] = 0
		}
		return out
	}
}

// Pulse returns an object that outputs a pulse at value 1.0 for a specified
// duration after its control source goes high
func Pulse(control source.Source, duration time.Duration, threshold, sampleRate float64) source.Source {
	var (
		active     = false
		pulseSteps = int(duration.Seconds() * sampleRate)
		curStep    int
	)
	return func(bufferSize int) (out [][]float32) {
		out = control(bufferSize)
		for i := range out[0] {
			if active {
				out[0][i] = 1
				curStep--
				if curStep <= 0 {
					active = false
				}
				continue
			}
			if float64(out[0][i]) > threshold {
				active = true
				curStep = pulseSteps
			}
			out[0][i] = 0
		}
		out[1] = out[0]
		return out
	}
}

// Gate allows the input signal to pass only when the control signal is above
// a certain threshold
func Gate(source, control source.Source, threshold float64) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = control(bufferSize)
		input := source(bufferSize)
		for i := range out[0] {
			if math.Abs(float64(out[0][i])) > threshold {
				out[0][i] = input[0][i]
				out[1][i] = input[1][i]
			} else {
				out[0][i] = 0
				out[1][i] = 0
			}
		}
		return out
	}
}
