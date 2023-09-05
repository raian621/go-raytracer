package lalg

import (
	"math"
	"testing"
)

const EPSILON = 0.001

func TestVec4Add(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{3.0, 2.0, 4.0, 1.0}

	v := Vec4Add(&v1, &v2)

	if math.Abs(v.x-4.0) > EPSILON {
		t.Errorf("`v.x` expected %f but got %f", 4.0, v.x)
	}
	if math.Abs(v.y-4.0) > EPSILON {
		t.Errorf("`v.y` expected %f but got %f", 4.0, v.y)
	}
	if math.Abs(v.z-7.0) > EPSILON {
		t.Errorf("`v.z` expected %f but got %f", 7.0, v.z)
	}
	if math.Abs(v.w-5.0) > EPSILON {
		t.Errorf("`v.w` expected %f but got %f", 5.0, v.w)
	}
}

func TestVec4Sub(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{3.0, 2.0, 4.0, 1.0}

	v := Vec4Sub(&v1, &v2)
	// expected value:
	e := Vec4{-2.0, 0.0, -1.0, 3.0}

	if math.Abs(v.x-e.x) > EPSILON {
		t.Errorf("`v.x` expected %f but got %f", e.x, v.x)
	}
	if math.Abs(v.y-e.y) > EPSILON {
		t.Errorf("`v.y` expected %f but got %f", e.y, v.y)
	}
	if math.Abs(v.z-e.z) > EPSILON {
		t.Errorf("`v.z` expected %f but got %f", e.z, v.z)
	}
	if math.Abs(v.w-e.w) > EPSILON {
		t.Errorf("`v.w` expected %f but got %f", e.w, v.w)
	}
}

func TestVec4Mul(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	f := 2.0

	v := Vec4Mul(&v1, f)
	// expected value:
	e := Vec4{2.0, 4.0, 6.0, 8.0}

	if math.Abs(v.x-e.x) > EPSILON {
		t.Errorf("`v.x` expected %f but got %f", e.x, v.x)
	}
	if math.Abs(v.y-e.y) > EPSILON {
		t.Errorf("`v.y` expected %f but got %f", e.y, v.y)
	}
	if math.Abs(v.z-e.z) > EPSILON {
		t.Errorf("`v.z` expected %f but got %f", e.z, v.z)
	}
	if math.Abs(v.w-e.w) > EPSILON {
		t.Errorf("`v.w` expected %f but got %f", e.w, v.w)
	}
}

func TestVec4Neg(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v := Vec4Neg(&v1)
	// expected value:
	e := Vec4{-1.0, -2.0, -3.0, -4.0}

	if math.Abs(v.x-e.x) > EPSILON {
		t.Errorf("`v.x` expected %f but got %f", e.x, v.x)
	}
	if math.Abs(v.y-e.y) > EPSILON {
		t.Errorf("`v.y` expected %f but got %f", e.y, v.y)
	}
	if math.Abs(v.z-e.z) > EPSILON {
		t.Errorf("`v.z` expected %f but got %f", e.z, v.z)
	}
	if math.Abs(v.w-e.w) > EPSILON {
		t.Errorf("`v.w` expected %f but got %f", e.w, v.w)
	}
}

func TestVec4Mag(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	magnitude := Vec4Mag(&v1)
	// expected value:
	e := math.Sqrt(30)

	if math.Abs(magnitude-e) > EPSILON {
		t.Errorf("`magnitude` expected %f but got %f", e, magnitude)
	}
}

func TestVec4Norm(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	normalized := Vec4Norm(&v1)
	// expected value:
	magnitude := math.Sqrt(30)
	e := Vec4{
		v1.x / magnitude,
		v1.y / magnitude,
		v1.z / magnitude,
		v1.w / magnitude,
	}

	if math.Abs(normalized.x-e.x) > EPSILON {
		t.Errorf("`normalized.x` expected %f but got %f", e.x, normalized.x)
	}
	if math.Abs(normalized.y-e.y) > EPSILON {
		t.Errorf("`normalized.y` expected %f but got %f", e.y, normalized.y)
	}
	if math.Abs(normalized.z-e.z) > EPSILON {
		t.Errorf("`normalized.z` expected %f but got %f", e.z, normalized.z)
	}
	if math.Abs(normalized.w-e.w) > EPSILON {
		t.Errorf("`normalized.w` expected %f but got %f", e.w, normalized.w)
	}
}

func TestVec4Dot(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{2.0, 6.0, 1.0, 2.0}

	e := 25.0
	d := Vec4Dot(&v1, &v2)

	if d-e > EPSILON {
		t.Errorf("`Vec4Dot(v1, v2)` expected %f but got %f", e, d)
	}
}

func TestVec4Cross(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{2.0, 6.0, 1.0, 2.0}

	e := Vec4{-16.0, 5.0, 2.0, 0.0}
	cross := Vec4Cross(v1, v2)

	if math.Abs(cross.x-e.x) > EPSILON {
		t.Errorf("`cross.x` expected %f but got %f", e.x, cross.x)
	}
	if math.Abs(cross.y-e.y) > EPSILON {
		t.Errorf("`cross.y` expected %f but got %f", e.y, cross.y)
	}
	if math.Abs(cross.z-e.z) > EPSILON {
		t.Errorf("`cross.z` expected %f but got %f", e.z, cross.z)
	}
	if math.Abs(cross.w-e.w) > EPSILON {
		t.Errorf("`cross.w` expected %f but got %f", e.w, cross.w)
	}
}
