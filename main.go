package main

import (
	"fmt"
	"time"

	"github.com/gordonklaus/portaudio"

	"github.com/kwoodhouse93/audio-playground/filter"
	"github.com/kwoodhouse93/audio-playground/generator"
	"github.com/kwoodhouse93/audio-playground/notes"
	"github.com/kwoodhouse93/audio-playground/router"
	"github.com/kwoodhouse93/audio-playground/sequence"
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

	s := source.New()

	noise := generator.UniformNoiseS(s)
	// sine := generator.SineM(s, 440, 0, sampleRate)
	// sineS := generator.SineS(s, 261.63, 440, 0, 0, sampleRate)
	// tri := generator.TriangleM(s, 440, 0, sampleRate)
	// triS := generator.TriangleS(s, 523.25, 659.25, 0, 0, sampleRate)
	// saw := generator.SawtoothM(s, 440, 0, sampleRate)
	// sawS := generator.SawtoothS(s, 261.63, 440, 0, 0, sampleRate)
	// sqr := generator.SquareM(s, 440.0, 0, 0.5, sampleRate)
	// sqrS := generator.SquareS(s, 261.63, 440.0, 0, 0, 0.5, sampleRate)
	//
	// mix := router.Mixer([]router.SourceGain{
	// 	{Source: noise, Gain: 0.0},
	// 	{Source: sine, Gain: 0.0},
	// 	{Source: sineS, Gain: 0.6},
	// 	{Source: tri, Gain: 0.0},
	// 	{Source: triS, Gain: 0.4},
	// 	{Source: saw, Gain: 0.0},
	// 	{Source: sawS, Gain: 0.0},
	// 	{Source: sqr, Gain: 0.0},
	// 	{Source: sqrS, Gain: 0.0},
	// })

	// // Noise 4/4
	nLfSqr := generator.SquareM(s, 2, 0, 0.1, sampleRate)
	nPulse := sequence.Pulse(nLfSqr, 50*time.Millisecond, 0.5, sampleRate)
	nGate := sequence.Gate(noise, nPulse, 0.5)
	nDly := filter.DelayFB(nGate, 100*time.Millisecond, 0.5, sampleRate)
	nSum := router.Mixer2(nGate, nDly, 0.5, 0.5)

	// Am chord
	sineSAm := generator.SineS(s, notes.C4, notes.A4, 0, 0, sampleRate)
	triSAm := generator.TriangleS(s, notes.C5, notes.E5, 0, 0, sampleRate)
	mixAm := router.Mixer2(sineSAm, triSAm, 0.6, 0.4)

	// E chord
	sineSE := generator.SineS(s, notes.B4, notes.E4, 0, 0, sampleRate)
	triSE := generator.TriangleS(s, notes.Gsharp4, notes.E3, 0, 0, sampleRate)
	mixE := router.Mixer2(sineSE, triSE, 0.6, 0.4)

	// mLfSqr := generator.SquareM(s, 0.5, 0, 0.1, sampleRate)
	// mPulse := sequence.Pulse(mLfSqr, 1990*time.Millisecond, 0.5, sampleRate)
	// mGate := sequence.Gate(mix, mPulse, 0.5)

	mSeq := sequence.Sequencer([]source.Source{
		mixAm,
		mixE,
	}, 2*time.Second, sampleRate)

	sum := router.SumComp(nSum, mSeq)

	// lpf := filter.LowPass(sum)
	// dly := filter.Delay(sum, 1000*time.Millisecond, sampleRate)
	// sumDly := router.Mixer2(dly, sum, 0.2, 0.6)

	sink := sink.New(sum)

	st, err := portaudio.OpenStream(p, sink)
	panicOnErr(err)

	defer st.Close()

	err = st.Start()
	panicOnErr(err)

	time.Sleep(8 * time.Second)

	err = st.Stop()
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
