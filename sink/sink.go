package sink

import "github.com/kwoodhouse93/audio-playground/source"

// A Sink takes an output buffer and fills it with samples to be streamed
type Sink func(outputBuffer [][]float32)

// New returns a Sink function object
func New(source source.Source) Sink {
	return func(out [][]float32) {
		buf := source(len(out[0]))
		for i := range out {
			copy(out[i], buf[i])
		}
	}
}
