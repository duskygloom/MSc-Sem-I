package stdform

import (
	"fmt"
	"simplex/matutils"

	"gonum.org/v1/gonum/mat"
)

type StdForm struct {
	a *mat.Dense // constraints coefficients
	b *mat.Dense // constraints value
	c *mat.Dense // objective coefficients
}

func (f *StdForm) A() *mat.Dense {
	return f.a
}

func (f *StdForm) B() *mat.Dense {
	return f.b
}

func (f *StdForm) C() *mat.Dense {
	return f.c
}

func NewStd(eqCoeff, eqValue, objCoeff *mat.Dense) (*StdForm, error) {
	if eqCoeff.RawMatrix().Rows != eqValue.RawMatrix().Cols {
		return nil, fmt.Errorf("NUMBER OF CONSTRAINTS DO NOT MATCH")
	}
	if eqCoeff.RawMatrix().Cols != objCoeff.RawMatrix().Cols {
		return nil, fmt.Errorf("NUMBER OF VARIABLES DO NOT MATCH")
	}
	return &StdForm{a: eqCoeff, b: eqValue, c: objCoeff}, nil
}

func (f *StdForm) Format() (fmt.Formatter, error) {
	constraints := matutils.HConcat(matutils.Transpose(f.b), f.a)
	objective := matutils.LeftPad(f.c, 1)
	result := matutils.VConcat(objective, constraints)
	return mat.Formatted(result), nil
}
