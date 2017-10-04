package compose

import (
	"errors"
)

// ThousandOp calls singleOp a thousand times via HundredOp and tenOp.
func ThousandOp(i int) (int, error) {
	return compose(
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
		HundredOp,
	)(i)
}

// HundredOp calls singleOp a hundred times via tenOp.
func HundredOp(i int) (int, error) {
	return compose(
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
		tenOp,
	)(i)
}

func tenOp(i int) (int, error) {
	return compose(
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
		singleOp,
	)(i)
}

func singleOp(i int) (int, error) {
	if i < 0 || i > 1000000 {
		return 0, errors.New("should not happen")
	}
	return i + 1, nil
}

func compose(funcs ...func(int) (int, error)) func(int) (int, error) {
	return func(i int) (int, error) {
		var err error
		for _, fun := range funcs {
			i, err = fun(i)
			if err != nil {
				return 0, err
			}
		}
		return i, nil
	}
}
