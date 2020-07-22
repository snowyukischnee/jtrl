package core



type Ray struct {
	Origin *Point3
	Direction *Vec3
}

func NewRay(origin *Vec3, direction *Point3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (ray *Ray) At(t float64) *Point3 {
	var ret Point3
	ret.AddScaledVec(ray.Origin, t, ray.Direction)
	return &ret
}

