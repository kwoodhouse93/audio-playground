package utils

import "math"

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
