package notes

// These consts represent the frequency of es in Western music
// nolint
const (
	C0      = 16.35
	Csharp0 = 17.32
	D0      = 18.35
	Dsharp0 = 19.45
	E0      = 20.60
	F0      = 21.83
	Fsharp0 = 23.12
	G0      = 24.50
	Gsharp0 = 25.96
	A0      = 27.50
	Asharp0 = 29.14
	B0      = 30.87
	C1      = 32.70
	Csharp1 = 34.65
	D1      = 36.71
	Dsharp1 = 38.89
	E1      = 41.20
	F1      = 43.65
	Fsharp1 = 46.25
	G1      = 49.00
	Gsharp1 = 51.91
	A1      = 55.00
	Asharp1 = 58.27
	B1      = 61.74
	C2      = 65.41
	Csharp2 = 69.30
	D2      = 73.42
	Dsharp2 = 77.78
	E2      = 82.41
	F2      = 87.31
	Fsharp2 = 92.50
	G2      = 98.00
	Gsharp2 = 103.83
	A2      = 110.00
	Asharp2 = 116.54
	B2      = 123.47
	C3      = 130.81
	Csharp3 = 138.59
	D3      = 146.83
	Dsharp3 = 155.56
	E3      = 164.81
	F3      = 174.61
	Fsharp3 = 185.00
	G3      = 196.00
	Gsharp3 = 207.65
	A3      = 220.00
	Asharp3 = 233.08
	B3      = 246.94
	C4      = 261.63
	Csharp4 = 277.18
	D4      = 293.66
	Dsharp4 = 311.13
	E4      = 329.63
	F4      = 349.23
	Fsharp4 = 369.99
	G4      = 392.00
	Gsharp4 = 415.30
	A4      = 440.00
	Asharp4 = 466.16
	B4      = 493.88
	C5      = 523.25
	Csharp5 = 554.37
	D5      = 587.33
	Dsharp5 = 622.25
	E5      = 659.25
	F5      = 698.46
	Fsharp5 = 739.99
	G5      = 783.99
	Gsharp5 = 830.61
	A5      = 880.00
	Asharp5 = 932.33
	B5      = 987.77
	C6      = 1046.50
	Csharp6 = 1108.73
	D6      = 1174.66
	Dsharp6 = 1244.51
	E6      = 1318.51
	F6      = 1396.91
	Fsharp6 = 1479.98
	G6      = 1567.98
	Gsharp6 = 1661.22
	A6      = 1760.00
	Asharp6 = 1864.66
	B6      = 1975.53
	C7      = 2093.00
	Csharp7 = 2217.46
	D7      = 2349.32
	Dsharp7 = 2489.02
	E7      = 2637.02
	F7      = 2793.83
	Fsharp7 = 2959.96
	G7      = 3135.96
	Gsharp7 = 3322.44
	A7      = 3520.00
	Asharp7 = 3729.31
	B7      = 3951.07
	C8      = 4186.01
	Csharp8 = 4434.92
	D8      = 4698.63
	Dsharp8 = 4978.03
	E8      = 5274.04
	F8      = 5587.65
	Fsharp8 = 5919.91
	G8      = 6271.93
	Gsharp8 = 6644.88
	A8      = 7040.00
	Asharp8 = 7458.62
	B8      = 7902.13
)

// Duration represents the length of a note in Western music
type Duration float64

// nolint
const (
	Large                Duration = 8.0
	Long                 Duration = 4.0
	DoubleWhole          Duration = 2.0
	Whole                Duration = 1.0
	Half                 Duration = 0.5
	Quarter              Duration = 0.25
	Eighth               Duration = 0.125
	Sixteenth            Duration = 0.0625
	ThirtySecond         Duration = 0.03125
	SixtyFourth          Duration = 0.015625
	HundredTwentyEighth  Duration = 0.0078125
	TwoHundredFiftySixth Duration = 0.00390625
)
