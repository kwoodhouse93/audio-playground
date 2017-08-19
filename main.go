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

	noise := generator.UniformNoiseS(s)
	sine := generator.SineM(s, 445, 0, sampleRate)
	sineS := generator.SineS(s, 261.63, 440, 0, 0, sampleRate)
	tri := generator.TriangleM(s, 440, 0, sampleRate)
	triS := generator.TriangleS(s, 261.63, 440, 0, 0, sampleRate)
	saw := generator.SawtoothM(s, 440, 0, sampleRate)
	sawS := generator.SawtoothS(s, 261.63, 440, 0, 0, sampleRate)

	mix := router.Mixer([]router.SourceGain{
		{Source: noise, Gain: 0.0},
		{Source: sine, Gain: 0.0},
		{Source: sineS, Gain: 0.0},
		{Source: tri, Gain: 0.0},
		{Source: triS, Gain: 0.0},
		{Source: saw, Gain: 0.0},
		{Source: sawS, Gain: 0.0},
	})
	sink := sink.New(mix)

	st, err := portaudio.OpenStream(p, sink)
	panicOnErr(err)

	defer st.Close()

	err = st.Start()
	panicOnErr(err)

	time.Sleep(1 * time.Second)

	err = st.Stop()
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
