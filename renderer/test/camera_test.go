package test

import (
	"renderer/core"
	"testing"
)

func TestCameraView(t *testing.T) {
	origin := core.NewBase3(1,1,0)
	width := 2.0
	height := 2.0
	focalLength := 1.0
	camera := core.NewCamera(origin, focalLength, width, height)
	camera.LookAt(core.NewBase3(1,1,1), core.NewBase3(0,1,0))
	llc := camera.CastRay(0,0)
	if !(llc.Direction.X == -1 && llc.Direction.Y == -1 && llc.Direction.Z == 1) {
		t.Fail()
	}
}
