package lalg_test

import (
	"math"
	"raytracer/lalg"
	"testing"
)

const EPSILON = 0.001

func TestVec4Add(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := lalg.Vec4{3.0, 2.0, 4.0, 1.0}

	v := lalg.Vec4Add(&v1, &v2)

	if math.Abs(v.X()-4.0) > EPSILON {
		t.Errorf("`v.X()` expected %f but got %f", 4.0, v.X())
	}
	if math.Abs(v.Y()-4.0) > EPSILON {
		t.Errorf("`v.Y()` expected %f but got %f", 4.0, v.Y())
	}
	if math.Abs(v.Z()-7.0) > EPSILON {
		t.Errorf("`v.Z()` expected %f but got %f", 7.0, v.Z())
	}
	if math.Abs(v.W()-5.0) > EPSILON {
		t.Errorf("`v.W()` expected %f but got %f", 5.0, v.W())
	}
}

func TestVec4Sub(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := lalg.Vec4{3.0, 2.0, 4.0, 1.0}

	v := lalg.Vec4Sub(&v1, &v2)
	// expected value:
	e := lalg.Vec4{-2.0, 0.0, -1.0, 3.0}

	if math.Abs(v.X()-e.X()) > EPSILON {
		t.Errorf("`v.X()` expected %f but got %f", e.X(), v.X())
	}
	if math.Abs(v.Y()-e.Y()) > EPSILON {
		t.Errorf("`v.Y()` expected %f but got %f", e.Y(), v.Y())
	}
	if math.Abs(v.Z()-e.Z()) > EPSILON {
		t.Errorf("`v.Z()` expected %f but got %f", e.Z(), v.Z())
	}
	if math.Abs(v.W()-e.W()) > EPSILON {
		t.Errorf("`v.W()` expected %f but got %f", e.W(), v.W())
	}
}

func TestVec4Mul(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	f := 2.0

	v := lalg.Vec4Mul(&v1, f)
	// expected value:
	e := lalg.Vec4{2.0, 4.0, 6.0, 8.0}

	if math.Abs(v.X()-e.X()) > EPSILON {
		t.Errorf("`v.X()` expected %f but got %f", e.X(), v.X())
	}
	if math.Abs(v.Y()-e.Y()) > EPSILON {
		t.Errorf("`v.Y()` expected %f but got %f", e.Y(), v.Y())
	}
	if math.Abs(v.Z()-e.Z()) > EPSILON {
		t.Errorf("`v.Z()` expected %f but got %f", e.Z(), v.Z())
	}
	if math.Abs(v.W()-e.W()) > EPSILON {
		t.Errorf("`v.W()` expected %f but got %f", e.W(), v.W())
	}
}

func TestVec4Neg(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v := lalg.Vec4Neg(&v1)
	// expected value:
	e := lalg.Vec4{-1.0, -2.0, -3.0, -4.0}

	if math.Abs(v.X()-e.X()) > EPSILON {
		t.Errorf("`v.X()` expected %f but got %f", e.X(), v.X())
	}
	if math.Abs(v.Y()-e.Y()) > EPSILON {
		t.Errorf("`v.Y()` expected %f but got %f", e.Y(), v.Y())
	}
	if math.Abs(v.Z()-e.Z()) > EPSILON {
		t.Errorf("`v.Z()` expected %f but got %f", e.Z(), v.Z())
	}
	if math.Abs(v.W()-e.W()) > EPSILON {
		t.Errorf("`v.W()` expected %f but got %f", e.W(), v.W())
	}
}

func TestVec4Mag(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	magnitude := lalg.Vec4Mag(&v1)
	// expected value:
	e := math.Sqrt(30)

	if math.Abs(magnitude-e) > EPSILON {
		t.Errorf("`magnitude` expected %f but got %f", e, magnitude)
	}
}

func TestVec4Norm(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	normalized := lalg.Vec4Norm(&v1)
	// expected value:
	magnitude := math.Sqrt(30)
	e := lalg.Vec4{
		v1.X() / magnitude,
		v1.Y() / magnitude,
		v1.Z() / magnitude,
		v1.W() / magnitude,
	}

	if math.Abs(normalized.X()-e.X()) > EPSILON {
		t.Errorf("`normalized.X()` expected %f but got %f", e.X(), normalized.X())
	}
	if math.Abs(normalized.Y()-e.Y()) > EPSILON {
		t.Errorf("`normalized.Y()` expected %f but got %f", e.Y(), normalized.Y())
	}
	if math.Abs(normalized.Z()-e.Z()) > EPSILON {
		t.Errorf("`normalized.Z()` expected %f but got %f", e.Z(), normalized.Z())
	}
	if math.Abs(normalized.W()-e.W()) > EPSILON {
		t.Errorf("`normalized.W()` expected %f but got %f", e.W(), normalized.W())
	}
}

func TestVec4Dot(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := lalg.Vec4{2.0, 6.0, 1.0, 2.0}

	e := 25.0
	d := lalg.Vec4Dot(&v1, &v2)

	if d-e > EPSILON {
		t.Errorf("`Vec4Dot(v1, v2)` expected %f but got %f", e, d)
	}
}

func TestVec4Cross(t *testing.T) {
	v1 := lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := lalg.Vec4{2.0, 6.0, 1.0, 2.0}

	e := lalg.Vec4{-16.0, 5.0, 2.0, 0.0}
	cross := lalg.Vec4Cross(&v1, &v2)

	if math.Abs(cross.X()-e.X()) > EPSILON {
		t.Errorf("`cross.X()` expected %f but got %f", e.X(), cross.X())
	}
	if math.Abs(cross.Y()-e.Y()) > EPSILON {
		t.Errorf("`cross.Y()` expected %f but got %f", e.Y(), cross.Y())
	}
	if math.Abs(cross.Z()-e.Z()) > EPSILON {
		t.Errorf("`cross.Z()` expected %f but got %f", e.Z(), cross.Z())
	}
	if math.Abs(cross.W()-e.W()) > EPSILON {
		t.Errorf("`cross.W()` expected %f but got %f", e.W(), cross.W())
	}
}

func TestVec4Cmp(t *testing.T) {
	v1 := &lalg.Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := &lalg.Vec4{1.0, 2.0, 3.0, 4.5}
	v3 := &lalg.Vec4{1.0, 2.0, 3.0, 3.5}
	v4 := &lalg.Vec4{1.0, 2.0, 3.0, 4.0}

	result := lalg.Vec4Cmp(v1, v2)
	if result != -4 {
		t.Errorf("`v1` should be less than `v2")
	}

	result = lalg.Vec4Cmp(v1, v3)
	if result != 4 {
		t.Errorf("`v1` should be greater than `v2")
	}

	result = lalg.Vec4Cmp(v1, v4)
	if result != 0 {
		t.Errorf("`v1` should be equal `v2")
	}
}
