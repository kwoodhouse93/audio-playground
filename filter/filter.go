package filter

import (
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/utils"
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

// Delay stores the samples for a given duration then plays them back delayed
func Delay(source source.Source, delay time.Duration, sampleRate float64) source.Source {
	delaySamples := utils.TimeToSteps(delay, sampleRate)
	delayBuf := utils.MakeBuffer(2, delaySamples)
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			var new [2]float32
			new[0], delayBuf[0] = delayBuf[0][0], delayBuf[0][1:]
			new[1], delayBuf[1] = delayBuf[1][1], delayBuf[1][1:]
			delayBuf[0] = append(delayBuf[0], out[0][i])
			delayBuf[1] = append(delayBuf[1], out[1][i])
			out[0][i] = new[0]
			out[1][i] = new[1]
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
