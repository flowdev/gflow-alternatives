package classic

import (
	"errors"
)

// ThousandOp calls singleOp a thousand times via HundredOp and tenOp.
func ThousandOp(i int) (int, error) {
	var err error

	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	i, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// HundredOp calls singleOp a hundred times via tenOp.
func HundredOp(i int) (int, error) {
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
