package filter

import (
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

// // LowPass is a simple LPF
// // Current implementation is just about the worst LPF imaginable
// func LowPass(source source.Source) source.Source {
// 	var last = [2]float32{0, 0}
// 	return func(bufferSize int) (out [][]float32) {
// 		out = source(bufferSize)
// 		for i := range out[0] {
// 			tmp := [2]float32{out[0][i], out[1][i]}
// 			out[0][i] = out[0][i] + last[0]
// 			out[1][i] = out[1][i] + last[1]
// 			last[0] = tmp[0]
// 			last[1] = tmp[1]
// 		}
// 		return out
// 	}
// }
