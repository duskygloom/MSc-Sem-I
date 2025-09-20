package matutils

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func HConcat(a, b *mat.Dense) *mat.Dense {
	arows, acols := a.Dims()
	brows, bcols := b.Dims()
	if arows != brows {
		panic(fmt.Errorf("SHOULD HAVE THE SAME NUMBER OF ROWS"))
	}
	c := mat.NewDense(arows, acols+bcols, nil)
	for i := range arows {
		for j := range acols {
			c.Set(i, j, a.At(i, j))
		}
		for j := range bcols {
			c.Set(i, j+acols, b.At(i, j))
		}
	}
	return c
}

func VConcat(a, b *mat.Dense) *mat.Dense {
	arows, acols := a.Dims()
	brows, bcols := b.Dims()
	if acols != bcols {
		panic(fmt.Errorf("SHOULD HAVE THE SAME NUMBER OF COLS"))
	}
	c := mat.NewDense(arows+brows, acols, nil)
	for i := range arows {
		for j := range acols {
			c.Set(i, j, a.At(i, j))
		}
	}
	for i := range brows {
		for j := range acols {
			c.Set(i+arows, j, b.At(i, j))
		}
	}
	return c
}
