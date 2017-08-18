package main

import (
	"log"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/gordonklaus/portaudio"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Important! Call gl.Init only under the presence of an active OpenGL context,
	// i.e., after MakeContextCurrent.
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	portaudio.Initialize()
	defer portaudio.Terminate()

	s, err := NewSine(30, sampleRate)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	err = s.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	err = s.Stop()
	if err != nil {
		panic(err)
	}
}

const sampleRate = 44100

type Sine struct {
	*portaudio.Stream
	step, phase float64
}

func NewSine(f, sampleRate float64) (*Sine, error) {
	s := &Sine{nil, f / sampleRate, 0}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.genSine)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Sine) genSine(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * s.phase))
		_, s.phase = math.Modf(s.phase + s.step)
	}
}
