package lalg

import (
	"math"
)

type Vec4 struct {
	x, y, z, w float64
}

func Vec4Add(v1 *Vec4, v2 *Vec4) *Vec4 {
	var v Vec4 = Vec4{
		v1.x + v2.x,
		v1.y + v2.y,
		v1.z + v2.z,
		v1.w + v2.w,
	}

	return &v
}

func Vec4Sub(v1 *Vec4, v2 *Vec4) *Vec4 {
	var v Vec4 = Vec4{
		v1.x - v2.x,
		v1.y - v2.y,
		v1.z - v2.z,
		v1.w - v2.w,
	}

	return &v
}

func Vec4Mul(v1 *Vec4, f float64) *Vec4 {
	var v Vec4 = Vec4{
		v1.x * f,
		v1.y * f,
		v1.z * f,
		v1.w * f,
	}

	return &v
}
func Vec4Neg(v1 *Vec4) *Vec4 {
	var v Vec4 = Vec4{
		-v1.x,
		-v1.y,
		-v1.z,
		-v1.w,
	}

	return &v
}

func Vec4Mag(v1 *Vec4) float64 {
	sum := v1.x*v1.x + v1.y*v1.y + v1.z*v1.z + v1.w*v1.w
	return math.Sqrt(sum)
}

func Vec4Norm(v *Vec4) *Vec4 {
	m := Vec4Mag(v)
	return &Vec4{
		v.x / m,
		v.y / m,
		v.z / m,
		v.w / m,
	}
}

func Vec4Dot(v1 *Vec4, v2 *Vec4) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z + v1.w*v2.w
}

func Vec4Cross(v1 *Vec4, v2 *Vec4) *Vec4 {

}
