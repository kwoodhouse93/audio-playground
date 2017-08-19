package main

import (
	"fmt"
	"time"

	"github.com/gordonklaus/portaudio"

	"github.com/kwoodhouse93/audio-playground/generator"
	"github.com/kwoodhouse93/audio-playground/router"
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

	s := source.New(p.Output.Channels)

	sine := generator.SineM(s, 82.41, 0, sampleRate)
	sineS := generator.SineS(s, 261.63, 440, 0, 0, sampleRate)
	noise := generator.UniformNoiseS(s)

	mix := router.Mixer([]router.SourceGain{
		{Source: sine, Gain: 0.1},
		{Source: sineS, Gain: 0.3},
		{Source: noise, Gain: 0.02},
	})
	sink := sink.New(mix)

	st, err := portaudio.OpenStream(p, sink)
	panicOnErr(err)

	defer st.Close()

	err = st.Start()
	panicOnErr(err)

	time.Sleep(2 * time.Second)

	err = st.Stop()
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
