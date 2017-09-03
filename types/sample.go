package types

type Sample []float32

func NewSample(channels int) Sample {
	return make([]float32, channels)
}
