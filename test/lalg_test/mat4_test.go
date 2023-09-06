package lalg_test

import (
	"math"
	"raytracer/lalg"
	"testing"
)

func TestMat4Cmp(t *testing.T) {
	m1 := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 3, 2},
	}
	m2 := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 3.5, 2},
	}
	m3 := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 2.5, 2},
	}

	rowRes, colRes := lalg.Mat4Cmp(m1, m2)
	if rowRes != -4 && colRes != -3 {
		t.Errorf("`m1` should be less than `m2, (%d, %d)", rowRes, colRes)
	}

	rowRes, colRes = lalg.Mat4Cmp(m1, m3)
	if rowRes != 4 && colRes != 3 {
		t.Errorf("`m1` should be greater than `m3`, (%d, %d)", rowRes, colRes)
	}

	rowRes, colRes = lalg.Mat4Cmp(m1, m1)
	if rowRes != 0 && colRes != 0 {
		t.Errorf("`m1` should be equal `m1, (%d, %d)", rowRes, colRes)
	}
}

func TestMat4Mul(t *testing.T) {
	m1 := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 3, 2},
	}
	m2 := &lalg.Mat4{
		&lalg.Vec4{-2, 1, 2, 3},
		&lalg.Vec4{3, 2, 1, -1},
		&lalg.Vec4{4, 3, 6, 5},
		&lalg.Vec4{1, 2, 7, 8},
	}
	e := &lalg.Mat4{
		&lalg.Vec4{20, 22, 50, 48},
		&lalg.Vec4{44, 54, 114, 108},
		&lalg.Vec4{40, 58, 110, 102},
		&lalg.Vec4{16, 26, 46, 42},
	}

	m := lalg.Mat4Mult(m1, m2)

	rowRes, colRes := lalg.Mat4Cmp(m, e)
	if rowRes != 0 {
		t.Errorf("Matrix multiplication failed (%d, %d) %s", colRes, rowRes,
			m.Stringify(),
		)
	}
}

func TestMat4MulVec4(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 3, 2},
	}
	v := &lalg.Vec4{1, 2, 3, 4}
	result := lalg.Mat4MultVec4(m, v)
	e := &lalg.Vec4{30, 70, 70, 30}

	if lalg.Vec4Cmp(result, e) != 0 {
		t.FailNow()
	}
}

func TestMat4MulScalar(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{1, 1, 1, 1},
		&lalg.Vec4{1, 1, 1, 1},
		&lalg.Vec4{1, 1, 1, 1},
		&lalg.Vec4{1, 1, 1, 1},
	}
	result := lalg.Mat4MultScalar(m, 4)
	e := &lalg.Mat4{
		&lalg.Vec4{4, 4, 4, 4},
		&lalg.Vec4{4, 4, 4, 4},
		&lalg.Vec4{4, 4, 4, 4},
		&lalg.Vec4{4, 4, 4, 4},
	}

	rowRes, _ := lalg.Mat4Cmp(result, e)
	if rowRes != 0 {
		t.Fail()
		t.Log(result.Stringify())
	}
}

func TestMat4Transpose(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{1, 2, 3, 4},
		&lalg.Vec4{5, 6, 7, 8},
		&lalg.Vec4{9, 8, 7, 6},
		&lalg.Vec4{5, 4, 3, 2},
	}
	e := &lalg.Mat4{
		&lalg.Vec4{1, 5, 9, 5},
		&lalg.Vec4{2, 6, 8, 4},
		&lalg.Vec4{3, 7, 7, 3},
		&lalg.Vec4{4, 8, 6, 2},
	}
	result := lalg.Mat4Transpose(m)
	rowRes, _ := lalg.Mat4Cmp(result, e)
	if rowRes != 0 {
		t.Fail()
		t.Log(result.Stringify())
	}
}

func TestMat4MinorAndCofactor(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{-2, -8, 3, 5},
		&lalg.Vec4{-3, 1, 7, 3},
		&lalg.Vec4{1, 2, -9, 6},
		&lalg.Vec4{-6, 7, 7, -9},
	}

	minor := lalg.Mat4Minor(m, 0, 0)
	cofactor := lalg.Mat4Cofactor(m, 0, 0)
	var eMinor float64 = 690
	var eCofactor float64 = 690
	if math.Abs(eMinor-minor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eMinor, minor)
	}
	if math.Abs(eCofactor-cofactor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eCofactor, cofactor)
	}

	minor = lalg.Mat4Minor(m, 0, 1)
	cofactor = lalg.Mat4Cofactor(m, 0, 1)
	eMinor = -447
	eCofactor = 447
	if math.Abs(eMinor-minor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eMinor, minor)
	}
	if math.Abs(eCofactor-cofactor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eCofactor, cofactor)
	}

	minor = lalg.Mat4Minor(m, 0, 2)
	cofactor = lalg.Mat4Cofactor(m, 0, 2)
	eMinor = 210
	eCofactor = 210
	if math.Abs(eMinor-minor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eMinor, minor)
	}
	if math.Abs(eCofactor-cofactor) >= lalg.EPSILON {
		t.Fail()
		t.Logf("%f, %f\n", eCofactor, cofactor)
	}
}

func TestMat4Determinant(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{-2, -8, 3, 5},
		&lalg.Vec4{-3, 1, 7, 3},
		&lalg.Vec4{1, 2, -9, 6},
		&lalg.Vec4{-6, 7, 7, -9},
	}
	determinant := lalg.Mat4Determinant(m)
	var e float64 = -4071

	if math.Abs(determinant-e) != 0 {
		t.Errorf("determinant: %f", determinant)
		t.Errorf("e:           %f", e)
	}
}

func TestUninvertibleMat4(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{-2, -8, 3, 5},
		&lalg.Vec4{-3, 1, 7, 3},
		&lalg.Vec4{1, 2, -9, 6},
		&lalg.Vec4{0, 0, 0, 0},
	}
	_, err := lalg.Mat4Inverse(m)
	if err == nil {
		t.Errorf("Error was not returned when uninvertible matrix was passed")
	}
}

func TestMat4Inverse(t *testing.T) {
	m := &lalg.Mat4{
		&lalg.Vec4{8, -5, 9, 2},
		&lalg.Vec4{7, 5, 6, 1},
		&lalg.Vec4{-6, 0, 9, 6},
		&lalg.Vec4{-3, 0, -9, -4},
	}
	inverse, _ := lalg.Mat4Inverse(m)
	e := &lalg.Mat4{
		&lalg.Vec4{-0.15385, -0.15385, -0.28205, -0.53846},
		&lalg.Vec4{-0.07692, 0.12308, 0.02564, 0.03077},
		&lalg.Vec4{0.35897, 0.35897, 0.43590, 0.92308},
		&lalg.Vec4{-0.69231, -0.69231, -0.76923, -1.92308},
	}

	rowCmp, _ := lalg.Mat4Cmp(inverse, e)
	if rowCmp != 0 {
		t.Error(inverse.Stringify())
	}
}
