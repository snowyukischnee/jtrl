package utils

import (
	"math"
	"math/rand"
	"renderer/core"
)

func RandF64(lower float64, upper float64) float64 {
	return lower + rand.Float64() * (upper - lower);
}

func RandUnitVector() *core.Base3 {
	a := RandF64(0, 2 * math.Pi)
	z := RandF64(-1, 1)
	r := math.Sqrt(1 - z * z)
	return core.NewBase3(r * math.Cos(a), r * math.Sin(a), z)
}
