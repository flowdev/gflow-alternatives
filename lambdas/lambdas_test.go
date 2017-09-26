package classic

import (
	"errors"
	"testing"
)

// BenchmarkLambdas1000 benchmarks the implementation of the thousandOp that uses lambdas.
func BenchmarkLambdas1000(b *testing.B) {
	thousandOp := newThousandOp(func(i int) {}, func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

func newThousandOp(out func(int), handleError func(error)) (in func(int)) {
	op10 := newHundredOp(out, handleError)
	op9 := newHundredOp(op10, handleError)
	op8 := newHundredOp(op9, handleError)
	op7 := newHundredOp(op8, handleError)
	op6 := newHundredOp(op7, handleError)
	op5 := newHundredOp(op6, handleError)
	op4 := newHundredOp(op5, handleError)
	op3 := newHundredOp(op4, handleError)
	op2 := newHundredOp(op3, handleError)
	in = newHundredOp(op2, handleError)
	return
}

func newHundredOp(out func(int), handleError func(error)) (in func(int)) {
	op10 := newTenOp(out, handleError)
	op9 := newTenOp(op10, handleError)
	op8 := newTenOp(op9, handleError)
	op7 := newTenOp(op8, handleError)
	op6 := newTenOp(op7, handleError)
	op5 := newTenOp(op6, handleError)
	op4 := newTenOp(op5, handleError)
	op3 := newTenOp(op4, handleError)
	op2 := newTenOp(op3, handleError)
	in = newTenOp(op2, handleError)
	return
}

func newTenOp(out func(int), handleError func(error)) (in func(int)) {
	op10 := newSingleOp(out, handleError)
	op9 := newSingleOp(op10, handleError)
	op8 := newSingleOp(op9, handleError)
	op7 := newSingleOp(op8, handleError)
	op6 := newSingleOp(op7, handleError)
	op5 := newSingleOp(op6, handleError)
	op4 := newSingleOp(op5, handleError)
	op3 := newSingleOp(op4, handleError)
	op2 := newSingleOp(op3, handleError)
	in = newSingleOp(op2, handleError)
	return
}

func newSingleOp(out func(int), handleError func(error)) (in func(int)) {
	in = func(i int) {
		if i < 0 || i > 1000000 {
			handleError(errors.New("should not happen"))
			return
		}
		out(i + 1)
	}
	return
}
