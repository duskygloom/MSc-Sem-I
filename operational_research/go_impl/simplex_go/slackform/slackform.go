package slackform

import (
	"fmt"
	"simplex/matutils"
	"simplex/stdform"

	"gonum.org/v1/gonum/mat"
)

type SlackForm struct {
	a, b, c  *mat.Dense
	basic    *mat.Dense // basic variables
	nonBasic *mat.Dense // non-basic variables
}

func NewSlack(eqCoeff, eqValue, objCoeff, basic, nonBasic *mat.Dense) (*SlackForm, error) {
	numEq, numVar := eqValue.RawMatrix().Cols, objCoeff.RawMatrix().Cols
	if eqCoeff.RawMatrix().Rows != numEq {
		return nil, fmt.Errorf("NUMBER OF CONSTRAINTS DO NOT MATCH")
	}
	if eqCoeff.RawMatrix().Cols != numVar {
		return nil, fmt.Errorf("NUMBER OF VARIABLES DO NOT MATCH")
	}
	if basic.RawMatrix().Cols+nonBasic.RawMatrix().Cols != numVar {
		return nil, fmt.Errorf("BASIC AND NON-BASIC ARE INCORRECT")
	}
	return &SlackForm{a: eqCoeff, b: eqValue, c: objCoeff, basic: basic, nonBasic: nonBasic}, nil
}

func FromStd(std *stdform.StdForm) (*SlackForm, error) {
	numEq, numVar := std.B().RawMatrix().Cols, std.C().RawMatrix().Cols
	A := mat.NewDense(numEq, numEq+numVar, nil)
	for i := range numEq {
		for j := range numVar {
			val := std.A().At(i, j)
			A.Set(i, j, val)
		}
		A.Set(i, i+numVar, 1)
	}
	C := matutils.RightPad(std.C(), numEq)
	return NewSlack(A, std.B(), C, mat.NewDense(1, numEq, nil), mat.NewDense(1, numVar, nil))
}

func (f *SlackForm) NumEq() int {
	return f.b.RawMatrix().Cols
}

func (f *SlackForm) NumVars() int {
	return f.c.RawMatrix().Cols
}

func (f *SlackForm) Format() (fmt.Formatter, error) {
	// add constraint values
	constraints := matutils.HConcat(matutils.Transpose(f.b), f.a)
	// add basic variables
	constraints = matutils.HConcat(matutils.Transpose(f.basic), constraints)
	// pad obj coefficients
	objective := matutils.LeftPad(f.c, 2)
	// pad zj
	zj := matutils.LeftPad(f.z(), 2)
	// add obj coefficients
	result := matutils.VConcat(objective, constraints)
	// add zj
	result = matutils.VConcat(result, zj)
	// whew finally done
	return mat.Formatted(result), nil
}

func (f *SlackForm) z() *mat.Dense {
	prod := mat.NewDense(1, f.NumVars(), nil)
	prod.Mul(f.basic, f.a)
	return prod
}

// Calculates Zj-Cj and checks if all values are non-negative.
// Returns true if all values are non-negative else false.
func (f *SlackForm) Proceed() bool {
	prod := f.z()
	prod.Sub(prod, f.c)
	fmt.Println(mat.Formatted(prod))
	for i := range f.NumVars() {
		if prod.At(0, i) < 0 {
			return true
		}
	}
	return false
}

func (f *SlackForm) Incoming() int {
	prod := f.z()
	prod.Sub(prod, f.c)
	index, value := 0, 0.0
	fmt.Println(mat.Formatted(prod))
	for i := range f.NumVars() {
		if prod.At(0, i) < value {
			index = i
			value = prod.At(0, i)
		}
	}
	return index
}

func (f *SlackForm) Outgoing() int {

}
