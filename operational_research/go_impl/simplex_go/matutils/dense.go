package matutils

import "gonum.org/v1/gonum/mat"

func Transpose(m *mat.Dense) *mat.Dense {
	return mat.DenseCopyOf(m.T())
}
