package source

// A Source creates a buffer based on the request made from the stream (via the sink)
type Source func(bufferSize int) (out [][]float32)

// NewSource returns a buffer with the specified number of channels
func NewSource(channels int) Source {
	ch := channels
	return func(bufferSize int) (out [][]float32) {
		out = make([][]float32, ch)
		for c := range out {
			out[c] = make([]float32, bufferSize)
		}
		return out
	}
}

// NewStereoSource returns a buffer with 2 output channels
func NewStereoSource() Source {
	return NewSource(2)
}
