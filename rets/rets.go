package rets

import (
	"errors"
)

// NewThousandOp returns an operation that calls singleOp a thousand times via
// hundredOp and tenOp.
func NewThousandOp() func(int) (int, error) {
	op1 := NewHundredOp()
	op2 := NewHundredOp()
	op3 := NewHundredOp()
	op4 := NewHundredOp()
	op5 := NewHundredOp()
	op6 := NewHundredOp()
	op7 := NewHundredOp()
	op8 := NewHundredOp()
	op9 := NewHundredOp()
	op10 := NewHundredOp()

	return func(i int) (int, error) {
		var err error

		i, err = op1(i)
		if err != nil {
			return 0, err
		}
		i, err = op2(i)
		if err != nil {
			return 0, err
		}
		i, err = op3(i)
		if err != nil {
			return 0, err
		}
		i, err = op4(i)
		if err != nil {
			return 0, err
		}
		i, err = op5(i)
		if err != nil {
			return 0, err
		}
		i, err = op6(i)
		if err != nil {
			return 0, err
		}
		i, err = op7(i)
		if err != nil {
			return 0, err
		}
		i, err = op8(i)
		if err != nil {
			return 0, err
		}
		i, err = op9(i)
		if err != nil {
			return 0, err
		}
		i, err = op10(i)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
}

// NewHundredOp returns an operation that calls singleOp a hundred times via
// tenOp.
func NewHundredOp() func(int) (int, error) {
	op1 := newTenOp()
	op2 := newTenOp()
	op3 := newTenOp()
	op4 := newTenOp()
	op5 := newTenOp()
	op6 := newTenOp()
	op7 := newTenOp()
	op8 := newTenOp()
	op9 := newTenOp()
	op10 := newTenOp()

	return func(i int) (int, error) {
		var err error

		i, err = op1(i)
		if err != nil {
			return 0, err
		}
		i, err = op2(i)
		if err != nil {
			return 0, err
		}
		i, err = op3(i)
		if err != nil {
			return 0, err
		}
		i, err = op4(i)
		if err != nil {
			return 0, err
		}
		i, err = op5(i)
		if err != nil {
			return 0, err
		}
		i, err = op6(i)
		if err != nil {
			return 0, err
		}
		i, err = op7(i)
		if err != nil {
			return 0, err
		}
		i, err = op8(i)
		if err != nil {
			return 0, err
		}
		i, err = op9(i)
		if err != nil {
			return 0, err
		}
		i, err = op10(i)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
}

func newTenOp() func(int) (int, error) {
	op1 := newSingleOp()
	op2 := newSingleOp()
	op3 := newSingleOp()
	op4 := newSingleOp()
	op5 := newSingleOp()
	op6 := newSingleOp()
	op7 := newSingleOp()
	op8 := newSingleOp()
	op9 := newSingleOp()
	op10 := newSingleOp()

	return func(i int) (int, error) {
		var err error

		i, err = op1(i)
		if err != nil {
			return 0, err
		}
		i, err = op2(i)
		if err != nil {
			return 0, err
		}
		i, err = op3(i)
		if err != nil {
			return 0, err
		}
		i, err = op4(i)
		if err != nil {
			return 0, err
		}
		i, err = op5(i)
		if err != nil {
			return 0, err
		}
		i, err = op6(i)
		if err != nil {
			return 0, err
		}
		i, err = op7(i)
		if err != nil {
			return 0, err
		}
		i, err = op8(i)
		if err != nil {
			return 0, err
		}
		i, err = op9(i)
		if err != nil {
			return 0, err
		}
		i, err = op10(i)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
}

func newSingleOp() func(int) (int, error) {
	return func(i int) (int, error) {
		if i < 0 || i > 1000000 {
			return 0, errors.New("should not happen")
		}
		return i + 1, nil
	}
}
