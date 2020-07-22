package test

import (
	"renderer/core"
	"renderer/utils"
	"testing"
)

func TestRayCast(t *testing.T)  {
	x0 := utils.RandF64(-100, 100)
	y0 := utils.RandF64(-100, 100)
	z0 := utils.RandF64(-100, 100)
	x1 := utils.RandF64(-100, 100)
	y1 := utils.RandF64(-100, 100)
	z1 := utils.RandF64(-100, 100)
	tt := utils.RandF64(-100, 100)
	a := core.NewBase3(x0, y0, z0)
	b := core.NewBase3(x1, y1, z1)
	x2 := x0 + tt * x1
	y2 := y0 + tt * y1
	z2 := z0 + tt * z1
	ray := core.NewRay(a, b)
	rt := ray.At(tt)
	if !(x2 == rt.X && y2 == rt.Y && z2 == rt.Z) {
		t.Fail()
	}
}
