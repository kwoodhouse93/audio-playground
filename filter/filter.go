package filter

import (
	"math"
	"time"

	"github.com/kwoodhouse93/audio-playground/source"
	"github.com/kwoodhouse93/audio-playground/types"
	"github.com/kwoodhouse93/audio-playground/utils"
)

// Zero sets a signal to 0
func Zero(source source.Source) source.Source {
	return func(step int) []float32 {
		out := source(step)
		out[0] = 0
		out[1] = 0
		return out
	}
}

// Delay stores the samples for a given duration then plays them back delayed
func Delay(src source.Source, delay time.Duration, sampleRate float64) source.Source {
	delaySamples := utils.TimeToSteps(delay, sampleRate)
	delayBuf := make([]*types.DelayLine, 2)
	delayBuf[0] = types.NewDelayLine(delaySamples)
	delayBuf[1] = types.NewDelayLine(delaySamples)
	return source.Cached(func(step int) []float32 {
		out := utils.MakeSample(2)
		s := src(step)
		out[0] = delayBuf[0].Step(s[0])
		out[1] = delayBuf[1].Step(s[1])
		return out
	})
}

// DelayFB is a delay with a feedback setting.
func DelayFB(src source.Source, delay time.Duration, feedback float32, sampleRate float64) source.Source {
	delaySamples := utils.TimeToSteps(delay, sampleRate)
	delayBuf := make([]*types.DelayLine, 2)
	delayBuf[0] = types.NewDelayLine(delaySamples)
	delayBuf[1] = types.NewDelayLine(delaySamples)
	return source.Cached(func(step int) []float32 {
		out := utils.MakeSample(2)
		s := src(step)
		out[0] = delayBuf[0].Read()
		out[1] = delayBuf[1].Read()
		delayBuf[0].Write(s[0] + (out[0] * feedback))
		delayBuf[1].Write(s[1] + (out[1] * feedback))
		return out
	})
}

// LowPass is a basic low pass filter
// H(s) = 1 / (s^2 + s/Q + 1)
func LowPass(src source.Source, cornerFreq, Q, sampleRate float64) source.Source {
	cosw0, _, alpha := calcCoefPrereqs(cornerFreq, Q, sampleRate)

	b0 := (1 - cosw0) / 2
	b1 := 1 - cosw0
	b2 := (1 - cosw0) / 2
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

// HighPass is a basic high pass filter
// H(s) = s^2 / (s^2 + s/Q + 1)
func HighPass(src source.Source, cornerFreq, Q, sampleRate float64) source.Source {
	cosw0, _, alpha := calcCoefPrereqs(cornerFreq, Q, sampleRate)

	b0 := (1 + cosw0) / 2
	b1 := -(1 + cosw0)
	b2 := (1 + cosw0) / 2
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

// BandPassPeakQ is a basic bandpass filter with constant skirt gain, and peak gain = Q
// H(s) = s / (s^2 + s/Q + 1)
func BandPassPeakQ(src source.Source, shelfFreq, Q, sampleRate float64) source.Source {
	cosw0, sinw0, alpha := calcCoefPrereqs(shelfFreq, Q, sampleRate)

	b0 := sinw0 / 2
	b1 := float32(0.0)
	b2 := -sinw0 / 2
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

// BandPassPeak0 is a basic bandpass filter, iwth a constant 0 dB peak gain
// H(s) = (s/Q) / (s^2 + s/Q + 1)
func BandPassPeak0(src source.Source, shelfFreq, Q, sampleRate float64) source.Source {
	cosw0, _, alpha := calcCoefPrereqs(shelfFreq, Q, sampleRate)

	b0 := alpha
	b1 := float32(0.0)
	b2 := -alpha
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

// Notch is a basic notch filter
// H(s) = (s^2 + 1) / (s^2 + s/Q + 1)
func Notch(src source.Source, notchFreq, Q, sampleRate float64) source.Source {
	cosw0, _, alpha := calcCoefPrereqs(notchFreq, Q, sampleRate)

	b0 := float32(1.0)
	b1 := -2 * cosw0
	b2 := float32(1.0)
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

// AllPass is a basic all pass filter
// H(s) = (s^2 - s/Q + 1) / (s^2 + s/Q + 1)
// quadFreq is the frequency at which the input and output go into quadrature
func AllPass(src source.Source, quadFreq, Q, sampleRate float64) source.Source {
	cosw0, _, alpha := calcCoefPrereqs(quadFreq, Q, sampleRate)

	b0 := 1 - alpha
	b1 := -2 * cosw0
	b2 := 1 + alpha
	a0 := 1 + alpha
	a1 := -2 * cosw0
	a2 := 1 - alpha

	return source.Cached(applyBiquadTransfer(src, b0, b1, b2, a0, a1, a2))
}

func calcCoefPrereqs(cornerFreq, Q, sampleRate float64) (cosw0, sinw0, alpha float32) {
	w0 := 2 * math.Pi * (cornerFreq / sampleRate)
	cosw0 = float32(math.Cos(w0))
	sinw0 = float32(math.Sin(w0))
	alpha = sinw0 / (2 * float32(Q))
	return
}

func applyBiquadTransfer(source source.Source, b0, b1, b2, a0, a1, a2 float32) source.Source {
	x := utils.MakeBuffer(2, 3)
	y := utils.MakeBuffer(2, 3)
	return func(step int) []float32 {
		out := utils.MakeSample(2)
		input := source(step)
		for i := range y {
			y[i][2] = y[i][1]
			y[i][1] = y[i][0]
			x[i][2] = x[i][1]
			x[i][1] = x[i][0]
			x[i][0] = input[i]
		}
		out[0] = biquadTransfer(x[0], y[0], a0, a1, a2, b0, b1, b2)
		out[1] = biquadTransfer(x[1], y[1], a0, a1, a2, b0, b1, b2)
		return out
	}
}

func biquadTransfer(x, y []float32, a0, a1, a2, b0, b1, b2 float32) float32 {
	y[0] = (b0/a0)*x[0] + (b1/a0)*x[1] + (b2/a0)*x[2] - (a1/a0)*y[1] - (a2/a0)*y[2]
	return y[0]
}
