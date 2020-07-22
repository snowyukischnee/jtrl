package utils

import "math/rand"

func RandF64(lower float64, upper float64) float64 {
	return lower + rand.Float64() * (upper - lower);
}
