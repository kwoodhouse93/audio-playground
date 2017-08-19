package filter

import "github.com/kwoodhouse93/audio-playground/source"

// NewPassthrough returns a filter that doesn't modify the source
func NewPassthrough(source source.Source) source.Source {
	return source
}

// NewGain returns a filter that multiplies the signals in the buffer by a constant value
func NewGain(source source.Source, gain float32) source.Source {
	return func(bufferSize int) (out [][]float32) {
		out = source(bufferSize)
		for i := range out[0] {
			out[0][i] = out[0][i] * gain
			out[1][i] = out[1][i] * gain
		}
		return out
	}
}
