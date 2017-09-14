package source

import "github.com/kwoodhouse93/audio-playground/types"

// A Source returns the next sample to be output to the stream (via the sink).
type Source func(step int) types.Sample

// Cached returns a source that caches the result of the evalFunc once per
// time step, even if called multiple times within that time step.
// It's intended use is to wrap the output of other sources to ensure multiple
// evaluations cannot occur in the same time step. If multiple evaluation is
// allowed to occcur, side effects of those evaluations can cause unusual
// errors, such as generating a wave with double the frequency it should have
// (by calculating two samples per time step, and only being able to pass one
// to the output).
func Cached(evalFunc func(int) types.Sample) Source {
	curStep := 0
	var curSample types.Sample
	return func(step int) types.Sample {
		if curStep == (step - 1) {
			curSample = evalFunc(step)
			curStep = step
		}
		return curSample
	}
}
