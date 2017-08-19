package main

import (
	"fmt"
	"time"

	"github.com/gordonklaus/portaudio"

	"github.com/kwoodhouse93/audio-playground/filter"
	"github.com/kwoodhouse93/audio-playground/generator"
	"github.com/kwoodhouse93/audio-playground/sink"
	"github.com/kwoodhouse93/audio-playground/source"
)

var sampleRate float64

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	h, err := portaudio.DefaultHostApi()
	panicOnErr(err)
	p := portaudio.LowLatencyParameters(nil, h.DefaultOutputDevice)
	p.Input.Channels = 0
	p.Output.Channels = 2
	sampleRate = p.SampleRate
	fmt.Printf("Sample rate: %f\n", sampleRate)

	source := source.NewSource(p.Output.Channels)
	// sine := generator.NewSineM(source, 440, 0, sampleRate)
	// gain := filter.NewGain(sine, 0.4)
	sineS := generator.NewSineS(source, 300, 440, 0, 0, sampleRate)
	gain := filter.NewGain(sineS, 0.2)
	sink := sink.NewSink(gain)

	s, err := portaudio.OpenStream(p, sink)
	panicOnErr(err)

	defer s.Close()

	err = s.Start()
	panicOnErr(err)

	time.Sleep(2 * time.Second)

	err = s.Stop()
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
