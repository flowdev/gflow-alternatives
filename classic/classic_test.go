package classic

import (
	"errors"
	"testing"
)

// BenchmarkClassic1000 benchmarks the classical implementation of the thousandOp.
func BenchmarkClassic1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

func thousandOp(i int) (int, error) {
	var err error

	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = hundredOp(i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func hundredOp(i int) (int, error) {
	var err error

	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	i, err = tenOp(i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func tenOp(i int) (int, error) {
	var err error

	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	i, err = singleOp(i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func singleOp(i int) (int, error) {
	if i < 0 || i > 1000000 {
		return 0, errors.New("should not happen")
	}
	return i + 1, nil
}
