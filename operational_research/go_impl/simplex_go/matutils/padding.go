package matutils

import "gonum.org/v1/gonum/mat"

func LeftPad(m *mat.Dense, padding int) *mat.Dense {
	pad := mat.NewDense(m.RawMatrix().Rows, padding, nil)
	return HConcat(pad, m)
}

func RightPad(m *mat.Dense, padding int) *mat.Dense {
	pad := mat.NewDense(m.RawMatrix().Rows, padding, nil)
	return HConcat(m, pad)
}
