package detect

import "time"

var (
	zeroFloat64 float64
	zeroString  string
)

func Float64(floats ...float64) float64 {
	for _, f := range floats {
		if f != zeroFloat64 {
			return f
		}
	}
	return zeroFloat64
}

func String(strings ...string) string {
	for _, s := range strings {
		if s != zeroString {
			return s
		}
	}
	return zeroString
}

func Time(times ...time.Time) time.Time {
	for _, t := range times {
		if !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
