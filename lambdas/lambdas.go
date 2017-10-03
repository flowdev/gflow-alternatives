package lambdas

import (
	"errors"
)

// NewThousandOp returns an operation that calls singleOp a thousand times via
// hundresOp and tenOp.
func NewThousandOp(out func(int), handleError func(error)) (in func(int)) {
	op10 := NewHundredOp(out, handleError)
	op9 := NewHundredOp(op10, handleError)
	op8 := NewHundredOp(op9, handleError)
	op7 := NewHundredOp(op8, handleError)
	op6 := NewHundredOp(op7, handleError)
	op5 := NewHundredOp(op6, handleError)
	op4 := NewHundredOp(op5, handleError)
	op3 := NewHundredOp(op4, handleError)
	op2 := NewHundredOp(op3, handleError)
	in = NewHundredOp(op2, handleError)
	return
}

// NewHundredOp returns an operation that calls singleOp a hundred times via
// tenOp.
func NewHundredOp(out func(int), handleError func(error)) (in func(int)) {
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
