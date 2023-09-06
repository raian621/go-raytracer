package lalg

import (
	"math"
)

type Vec4 [4]float64

func (v Vec4) X() float64 {
	return v[0]
}

func (v Vec4) Y() float64 {
	return v[1]
}

func (v Vec4) Z() float64 {
	return v[2]
}

func (v Vec4) W() float64 {
	return v[3]
}

func Vec4Add(v1 *Vec4, v2 *Vec4) *Vec4 {
	return &Vec4{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
		v1.Z() + v2.Z(),
		v1.W() + v2.W(),
	}
}

func Vec4Sub(v1 *Vec4, v2 *Vec4) *Vec4 {
	return &Vec4{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
		v1.Z() - v2.Z(),
		v1.W() - v2.W(),
	}
}

func Vec4Mul(v1 *Vec4, f float64) *Vec4 {
	return &Vec4{
		v1.X() * f,
		v1.Y() * f,
		v1.Z() * f,
		v1.W() * f,
	}
}

func Vec4Neg(v1 *Vec4) *Vec4 {
	return &Vec4{
		-v1.X(),
		-v1.Y(),
		-v1.Z(),
		-v1.W(),
	}
}

func Vec4Mag(v1 *Vec4) float64 {
	sum := v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z() + v1.W()*v1.W()
	return math.Sqrt(sum)
}

func Vec4Norm(v *Vec4) *Vec4 {
	m := Vec4Mag(v)
	return &Vec4{
		v.X() / m,
		v.Y() / m,
		v.Z() / m,
		v.W() / m,
	}
}

func Vec4Dot(v1 *Vec4, v2 *Vec4) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z() + v1.W()*v2.W()
}

func Vec4Cross(v1 *Vec4, v2 *Vec4) *Vec4 {
	return &Vec4{
		v1.Y()*v2.Z() - v1.Z()*v2.Y(),
		v1.Z()*v2.X() - v1.X()*v2.Z(),
		v1.X()*v2.Y() - v1.Y()*v2.X(),
		0.0,
	}
}

/*
Comparator function for Vec4. Vec4s are ordered lexicographically using this
function. Returns -n if v1 is greater than v2, returns 0 if v1 and v2 are equal,
returns n if v1 is less than v2. `n` being equal to 1 + the index of the
nonmatching elements
*/
func Vec4Cmp(v1 *Vec4, v2 *Vec4) int {
	for i := 0; i < 4; i++ {
		if math.Abs(v1[i]-v2[i]) < EPSILON {
			continue
		} else if v1[i] < v2[i] {
			return -(i + 1)
		} else {
			return i + 1
		}
	}
	return 0
}
