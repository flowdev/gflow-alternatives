package classic

import (
	"errors"
	"testing"
)

// BenchmarkRets1000 benchmarks the implementation of the thousandOp that uses many return values.
func BenchmarkRets1000(b *testing.B) {
	thousandOp := newThousandOp()
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

// BenchmarkIllegalRets1000 benchmarks an illegal implementation of the thousandOp that uses many return values.
func BenchmarkIllegalRets1000(b *testing.B) {
	thousandOp := illegalThousandOp()
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

func newThousandOp() func(int) (int, error) {
	op1 := newHundredOp()
	op2 := newHundredOp()
	op3 := newHundredOp()
	op4 := newHundredOp()
	op5 := newHundredOp()
	op6 := newHundredOp()
	op7 := newHundredOp()
	op8 := newHundredOp()
	op9 := newHundredOp()
	op10 := newHundredOp()

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

func illegalThousandOp() func(int) (int, error) {
	op1 := illegalHundredOp()
	op2 := illegalHundredOp()
	op3 := illegalHundredOp()
	op4 := illegalHundredOp()
	op5 := illegalHundredOp()
	op6 := illegalHundredOp()
	op7 := illegalHundredOp()
	op8 := illegalHundredOp()
	op9 := illegalHundredOp()
	op10 := illegalHundredOp()

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

func newHundredOp() func(int) (int, error) {
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

func illegalHundredOp() func(int) (int, error) {
	op1 := illegalTenOp()
	op2 := illegalTenOp()
	op3 := illegalTenOp()
	op4 := illegalTenOp()
	op5 := illegalTenOp()
	op6 := illegalTenOp()
	op7 := illegalTenOp()
	op8 := illegalTenOp()
	op9 := illegalTenOp()
	op10 := illegalTenOp()

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

func illegalTenOp() func(int) (int, error) {
	op1 := illegalSingleOp
	op2 := illegalSingleOp
	op3 := illegalSingleOp
	op4 := illegalSingleOp
	op5 := illegalSingleOp
	op6 := illegalSingleOp
	op7 := illegalSingleOp
	op8 := illegalSingleOp
	op9 := illegalSingleOp
	op10 := illegalSingleOp

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

func illegalSingleOp(i int) (int, error) {
	if i < 0 || i > 1000000 {
		return 0, errors.New("should not happen")
	}
	return i + 1, nil
}
