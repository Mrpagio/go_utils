package utils

import "math"

func IsFracPartZero(f float64) bool {
	_, frac := math.Modf(f)
	return math.Abs(frac) < 1e-9
}
