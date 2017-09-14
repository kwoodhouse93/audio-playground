package types

// Sample represents an audio sample with a fixed number of audio channels.
type Sample []float32

// NewSample creates a new Sample with the given number of channels.
func NewSample(channels int) Sample {
	return make([]float32, channels)
}

// Channels returns the number of channels that the given sample represents.
func (s Sample) Channels() int {
	return len(s)
}

// Set sets the given sample to value v.
func (s Sample) Set(v float32) Sample {
	for i := range s {
		s[i] = v
	}
	return s
}

// Gain multiplies all values in s by a constant value.
func (s Sample) Gain(g float32) Sample {
	for i := range s {
		s[i] = s[i] * g
	}
	return s
}

// Sum adds s2 to s.
func (s Sample) Sum(s2 Sample) Sample {
	if s.Channels() != s2.Channels() {
		panic("Samples must have same number of channels to multiply")
	}
	for i := range s {
		s[i] = s[i] + s2[i]
	}
	return s
}

// Mult multiplies s by s2.
func (s Sample) Mult(s2 Sample) Sample {
	if s.Channels() != s2.Channels() {
		panic("Samples must have same number of channels to multiply")
	}
	for i := range s {
		s[i] = s[i] * s2[i]
	}
	return s
}
