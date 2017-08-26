package source

import "github.com/kwoodhouse93/audio-playground/utils"

// A Source returns the next sample to be output to the stream (via the sink)
// type Source func(bufferSize int) (out [][]float32)
type Source func() []float32

// New returns a buffer with the specified number of channels
func New() Source {
	return func() []float32 {
		return utils.MakeSample(2)
	}
}
