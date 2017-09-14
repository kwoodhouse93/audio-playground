package types

import (
	"fmt"
	"strings"
)

// guardState is used to ensure the DelayLine's low-level Read and Write
// functions are used correctly
type guardState bool

const (
	readyToRead  guardState = false
	readyToWrite guardState = true
)

// DelayLine represents a circular buffer that can be used to delay a signal
// by a fixed number of samples.
type DelayLine struct {
	buf   []float32
	ptr   uintptr
	guard guardState
}

// NewDelayLine creates a new DelayLine for the given number of samples.
func NewDelayLine(len int) *DelayLine {
	if len < 1 {
		panic("DelayLine cannot have length < 1")
	}
	buf := make([]float32, len)
	return &DelayLine{
		buf: buf,
		ptr: 0,
	}
}

// Step returns the output value and writes the next input value back to the
// front of the delay line.
func (d *DelayLine) Step(in float32) (out float32) {
	out = d.Read()
	d.Write(in)
	return out
}

// Read only outputs the value at the end of the delay line, and must be
// followed by a Write.
func (d *DelayLine) Read() (out float32) {
	if d.guard != readyToRead {
		panic("Cannot read twice without writing")
	}
	out = d.buf[d.ptr]
	d.guard = readyToWrite
	return out
}

// Write writes the input value to the buffer and must be preceded by a Read.
func (d *DelayLine) Write(in float32) {
	if d.guard != readyToWrite {
		panic("Cannot write without reading first")
	}
	d.buf[d.ptr] = in
	d.ptr = (d.ptr + 1) % uintptr(len(d.buf))
	d.guard = readyToRead
}

// String returns a string representation of the delay line
func (d DelayLine) String() string {
	str := []string{fmt.Sprintf("%v\n", d.buf)}
	str = append(str, fmt.Sprintf(" "))
	for i := range d.buf {
		if d.ptr == uintptr(i) {
			str = append(str, fmt.Sprintf("^ "))
			continue
		}
		str = append(str, fmt.Sprintf("  "))
	}
	return strings.Join(str, "")
}

// SampleDelayLine is a delay line that operates on samples (with multiple
// channels).
type SampleDelayLine struct {
	bufs  []*DelayLine
	guard guardState
}

// NewSampleDelayLine creates a new SampleDelayLine.
func NewSampleDelayLine(channels, len int) *SampleDelayLine {
	bufs := make([]*DelayLine, channels)
	for i := range bufs {
		bufs[i] = NewDelayLine(len)
	}
	return &SampleDelayLine{
		bufs: bufs,
	}
}

// Step returns the output sample and writes the next input sample back to the
// front of the delay line.
func (d *SampleDelayLine) Step(in Sample) (out Sample) {
	if in.Channels() != len(d.bufs) {
		panic("Input sample must have same number of channels as SampleDelayLine")
	}
	out = NewSample(len(d.bufs))
	for i := range d.bufs {
		out[i] = d.bufs[i].Step(in[i])
	}
	return out
}

// Read only outputs the sample at the end of the delay line, and must be
// followed by a Write.
func (d *SampleDelayLine) Read() (out Sample) {
	if d.guard != readyToRead {
		panic("Cannot read twice without writing")
	}
	out = NewSample(len(d.bufs))
	for i := range d.bufs {
		out[i] = d.bufs[i].Read()
	}
	d.guard = readyToWrite
	return out
}

// Write writes the input sample to the buffer and must be preceded by a Read.
func (d *SampleDelayLine) Write(in Sample) {
	if in.Channels() != len(d.bufs) {
		panic("Input sample must have same number of channels as SampleDelayLine")
	}
	if d.guard != readyToWrite {
		panic("Cannot write without reading first")
	}
	for i := range d.bufs {
		d.bufs[i].Write(in[i])
	}
	d.guard = readyToRead
}
