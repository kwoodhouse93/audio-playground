package sink

import "github.com/kwoodhouse93/audio-playground/source"

// A Sink takes an output buffer and fills it with samples to be streamed
type Sink func(outputBuffer [][]float32)

// New returns a Sink function object
func New(source source.Source) Sink {
	curStep := 0
	return func(out [][]float32) {
		for i := range out[0] {
			curStep++
			sample := source(curStep)
			out[0][i] = sample[0]
			out[1][i] = sample[1]
		}
	}
}
