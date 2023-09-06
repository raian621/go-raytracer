package lalg

import (
	"fmt"
	"strings"
)

type Mat4 [4]*Vec4

func (m Mat4) Stringify() string {
	var sb strings.Builder
	sb.WriteString("{\n")
	for i := 0; i < 4; i++ {
		sb.WriteString("  ")
		sb.WriteString(fmt.Sprint(m[i]))
		sb.WriteString(",\n")
	}
	sb.WriteString("}\n")

	return sb.String()
}

func Mat4Add(m1 *Mat4, m2 *Mat4) *Mat4 {
	return &Mat4{
		Vec4Add(m1[0], m2[0]),
		Vec4Add(m1[1], m2[1]),
		Vec4Add(m1[2], m2[2]),
		Vec4Add(m1[3], m2[3]),
	}
}

func Mat4Sub(m1 *Mat4, m2 *Mat4) *Mat4 {
	return &Mat4{
		Vec4Sub(m1[0], m2[0]),
		Vec4Sub(m1[1], m2[1]),
		Vec4Sub(m1[2], m2[2]),
		Vec4Sub(m1[3], m2[3]),
	}
}

func Mat4Id() *Mat4 {
	return &Mat4{
		&Vec4{1, 0, 0, 0},
		&Vec4{0, 1, 0, 0},
		&Vec4{0, 0, 1, 0},
		&Vec4{0, 0, 0, 1},
	}
}

func Mat4Zero() *Mat4 {
	return &Mat4{
		&Vec4{},
		&Vec4{},
		&Vec4{},
		&Vec4{},
	}
}

/*
Comparator function for Mat4. Mat4s are ordered lexicographically using this
function. Returns (-col, -row) if m1 is greater than m2, returns 0, 0 if m1 and
m2 are equal, returns (col, row) if v1 is less than v2. row and col being equal
to 1 + the indices of the nonmatching elements
*/
func Mat4Cmp(m1 *Mat4, m2 *Mat4) (int, int) {
	for i := 0; i < 4; i++ {
		rowCmp := Vec4Cmp(m1[i], m2[i])
		if rowCmp == 0 {
			continue
		} else if rowCmp < 0 {
			return -(i + 1), rowCmp
		} else {
			return i + 1, rowCmp
		}
	}
	return 0, 0
}

func Mat4Mult(m1 *Mat4, m2 *Mat4) *Mat4 {
	m := Mat4Zero()
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			for k := 0; k < 4; k++ {
				m[row][col] += m1[row][k] * m2[k][col]
			}
		}
	}
	return m
}

func Mat4MultScalar(m *Mat4, s float64) *Mat4 {
	return &Mat4{
		Vec4Mul(m[0], s),
		Vec4Mul(m[1], s),
		Vec4Mul(m[2], s),
		Vec4Mul(m[3], s),
	}
}

func Mat4MultVec4(m *Mat4, v *Vec4) *Vec4 {
	returnVec := &Vec4{}
	for row := 0; row < 4; row++ {
		returnVec[row] = Vec4Dot(m[row], v)
	}
	return returnVec
}

func Mat4Translate(x, y, z float64) *Mat4 {
	return &Mat4{
		&Vec4{1, 0, 0, x},
		&Vec4{0, 1, 0, y},
		&Vec4{0, 0, 1, z},
		&Vec4{0, 0, 0, 1},
	}
}

func Mat4Transpose(m *Mat4) *Mat4 {
	returnMat4 := Mat4Zero()
	// set the values on the diagonal
	for mid := 0; mid < 4; mid++ {
		returnMat4[mid][mid] = m[mid][mid]
	}
	// set other values
	for row := 1; row < 4; row++ {
		for col := 0; col < row; col++ {
			returnMat4[row][col] = m[col][row]
			returnMat4[col][row] = m[row][col]
		}
	}
	return returnMat4
}

func Mat4Minor(m *Mat4, row, col int) float64 {
	rows := make([]int, 3)
	cols := make([]int, 3)

	for i := 0; i < 3; i++ {
		// copy every row index except for i
		if i < row {
			rows[i] = i
		} else {
			rows[i] = i + 1
		}
		// copy every col index except for j
		if i < col {
			cols[i] = i
		} else {
			cols[i] = i + 1
		}
	}

	return mat4MinorAux(m, rows, cols)
}

func mat4MinorAux(m *Mat4, rows, cols []int) float64 {
	if len(rows) == 2 {
		return m[rows[0]][cols[0]]*m[rows[1]][cols[1]] -
			m[rows[0]][cols[1]]*m[rows[1]][cols[0]]
	}

	i := 0
	var minor float64 = 0
	for j := 0; j < len(cols); j++ {
		newRows := make([]int, len(rows)-1)
		newCols := make([]int, len(cols)-1)

		for k := 0; k < len(newRows); k++ {
			// copy every row index except for i
			if k < i {
				newRows[k] = rows[k]
			} else {
				newRows[k] = rows[k+1]
			}
			// copy every col index except for j
			if k < j {
				newCols[k] = cols[k]
			} else {
				newCols[k] = cols[k+1]
			}
		}

		if (i+j)%2 == 0 {
			minor += m[rows[i]][cols[j]] * mat4MinorAux(m, newRows, newCols)
		} else {
			minor -= m[rows[i]][cols[j]] * mat4MinorAux(m, newRows, newCols)
		}
	}
	return minor
}

func Mat4Cofactor(m *Mat4, row, col int) float64 {
	if (row+col)%2 == 0 {
		return Mat4Minor(m, row, col)
	}
	return -Mat4Minor(m, row, col)
}
