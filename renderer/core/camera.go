package core

type Camera struct {
	Origin *Point3
	FocalLength float64
	Width float64
	Height float64
	up *Vec3
	right *Vec3
	lowerLeftCorner *Point3
}

func NewCamera(origin *Point3, focalLength float64, width float64, height float64) *Camera {
	return &Camera{
		Origin:      origin,
		FocalLength: focalLength,
		Width:       width,
		Height:      height,
	}
}

func (camera *Camera) LookAt(target *Point3, up *Vec3) {
	viewDir := target.SubV(camera.Origin)
	planeNormal := viewDir.UnitVector()
	focalPoint := planeNormal.MulScalar(camera.FocalLength).AddV(camera.Origin)
	uDir := focalPoint.AddV(up).SubV(camera.Origin)
	camera.right = uDir.Cross(planeNormal).UnitVector()
	camera.up = planeNormal.Cross(camera.right).UnitVector()
	camera.lowerLeftCorner = focalPoint.SubV(camera.up.MulScalar(camera.Height / 2)).SubV(camera.right.MulScalar(camera.Width / 2))
}

func (camera *Camera) CastRay(u float64, v float64) *Ray {
	return NewRay(
		camera.Origin,
		camera.lowerLeftCorner.AddV(camera.right.MulScalar(u)).AddV(camera.up.MulScalar(v)).SubV(camera.Origin),
	)
}