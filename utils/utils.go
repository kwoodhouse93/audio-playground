package utils

import "math"

// Triangle returns the value of a triangle wave at x radians
// FIXME: Should a triangle wave start at 0 for x=0?
func Triangle(x float64) float64 {
	var sign float64 = 1
	if math.Mod(x, 2*math.Pi) > math.Pi {
		sign = -1
	}
	return sign * (((2 / math.Pi) * math.Mod(x, math.Pi)) - 1)
}
