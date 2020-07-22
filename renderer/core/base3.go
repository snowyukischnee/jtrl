package core

import "math"

type Base3 struct {
	X float64
	Y float64
	Z float64
}

func NewBase3(x float64, y float64, z float64) *Base3 {
	return &Base3{X: x, Y: y, Z: z}
}

func (v *Base3) Neg() *Base3 {
	return NewBase3(-v.X, -v.Y, -v.Z)
}

func (v *Base3) AddV(vo *Base3) *Base3 {
	return NewBase3(v.X + vo.X, v.Y + vo.Y, v.Z + vo.Z)
}

func (v *Base3) SubV(vo *Base3) *Base3 {
	return NewBase3(v.X - vo.X, v.Y - vo.Y, v.Z - vo.Z)
}

func (v *Base3) MulV(vo *Base3) *Base3 {
	return NewBase3(v.X * vo.X, v.Y * vo.Y, v.Z * vo.Z)
}

func (v *Base3) AddScalar(t float64) *Base3 {
	return NewBase3(v.X + t, v.Y + t, v.Z + t)
}

func (v *Base3) SubScalar(t float64) *Base3 {
	return NewBase3(v.X - t, v.Y - t, v.Z - t)
}

func (v *Base3) MulScalar(t float64) *Base3 {
	return NewBase3(v.X * t, v.Y * t, v.Z * t)
}

func (v *Base3) DivScalar(t float64) *Base3 {
	if t != 0 {
		return NewBase3(v.X / t, v.Y / t, v.Z / t)
	} else {
		return NewBase3(0, 0, 0)
	}
}

func (v *Base3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func AddScaledV(v0 *Base3, v1 *Base3, t float64) *Base3 {
	return v0.AddV(v1.MulScalar(t))
}
func (v *Base3) Dot(vo *Base3) float64 {
	return v.X * vo.X + v.Y * vo.Y + v.Z * vo.Z
}

func (v *Base3) Cross(vo *Base3) *Base3 {
	return NewBase3(v.Y * vo.Z - v.Z * vo.Y, v.Z * vo.X - v.X * vo.Z, v.X * vo.Y - v.Y * vo.X)
}

func (v *Base3) UnitVector() *Base3 {
	return v.DivScalar(v.Length())
}
