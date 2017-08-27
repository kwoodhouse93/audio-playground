package source

import (
	"github.com/kwoodhouse93/audio-playground/utils"
)

// A Source returns the next sample to be output to the stream (via the sink)
// type Source func(bufferSize int) (out [][]float32)
type Source func(step int) []float32

// New returns a buffer with the specified number of channels
func New() Source {
	return func(step int) []float32 {
		return utils.MakeSample(2)
	}
}

func Cached(evalFunc func(int) []float32) Source {
	curStep := 0
	var curSample []float32
	return func(step int) []float32 {
		if curStep == (step - 1) {
			curSample = evalFunc(step)
			curStep = step
		}
		return curSample
	}
}
