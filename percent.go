package sutils

import "strconv"

const (
	OnePercent Percent = 0.01
)

type Percent float64

func PercentFromInt(percent int) Percent {
	return Percent(percent) / 100.
}

func PercentFromInt64(percent int64) Percent {
	return Percent(percent) / 100.
}

func PercentFromFloat64(percent float64) Percent {
	return Percent(percent)
}

func (p Percent) String() string {
	return strconv.FormatFloat(100.*float64(p), 'f', -1, 64) + "%"
}

func (p Percent) Float64() float64 {
	return float64(p)
}

// Scale returns p% imply that val is 100%
func (p Percent) Scale(val float64) float64 {
	return float64(p) * val
}

// ScaleReversed returns 100% imply that val is p%
func (p Percent) ScaleReversed(val float64) float64 {
	if p == 0. {
		return 0.
	}
	return 1. / float64(p) * val
}
