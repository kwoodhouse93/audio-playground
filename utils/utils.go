package utils

import (
	"math"
	"time"
)

// MakeBuffer returns an empty, but initialised buffer
func MakeBuffer(channels, samples int) (out [][]float32) {
	out = make([][]float32, channels)
	for c := range out {
		out[c] = make([]float32, samples)
	}
	return out
}

// TimeToSteps converts a duration to a number of samples based on the sample rate
func TimeToSteps(duration time.Duration, sampleRate float64) int {
	return int(duration.Seconds() * sampleRate)
}

// Triangle returns the value of a triangle wave at x radians
// Output is in the range [-1.0, 1.0], and a cycle is 2*pi radians
// FIXME: Should a triangle wave start at 0 for x=0?
func Triangle(x float64) float64 {
	var sign float64 = 1
	if math.Mod(x, 2*math.Pi) > math.Pi {
		sign = -1
	}
	return sign * (((2 / math.Pi) * math.Mod(x, math.Pi)) - 1)
}

// Sawtooth returns the value of a sawtooth wave at x radians
// Output is in the range [-1.0, 1.0], and a cycle is 2*pi radians
func Sawtooth(x float64) float64 {
	return ((1 / math.Pi) * math.Mod(x, 2*math.Pi)) - 1
}

// SquareFunc returns a function which returns the value of a square
// wave at x radians, with the given duty cycle.
// Output is in the range [-1.0, 1.0], and a cycle is 2*pi radians
func SquareFunc(dutyCycle float64) func(float64) float64 {
	return func(x float64) float64 {
		if math.Mod(x, 2*math.Pi) < (2 * math.Pi * dutyCycle) {
			return 1
		}
		return -1
	}
}
