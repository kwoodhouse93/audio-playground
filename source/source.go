package source

import "github.com/kwoodhouse93/audio-playground/utils"

// A Source creates a buffer based on the request made from the stream (via the sink)
type Source func(bufferSize int) (out [][]float32)

// New returns a buffer with the specified number of channels
func New(channels int) Source {
	ch := channels
	return func(bufferSize int) (out [][]float32) {
		return utils.MakeBuffer(ch, bufferSize)
	}
}
