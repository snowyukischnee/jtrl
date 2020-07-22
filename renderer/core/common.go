package core

import "gonum.org/v1/gonum/mat"

type Point3 = mat.VecDense
type Vec3 = mat.VecDense

func NewVec3(x float64, y float64, z float64) *Vec3 {
	return mat.NewVecDense(3, []float64{x,y,z})
}