package classicrets

import (
	"errors"
)

// ThousandOp calls singleOp a thousand times via HundredOp and tenOp.
func ThousandOp(i int) (portOut int, portError error) {
	var err error
	pi := i

	pi, err = HundredOp(i)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = HundredOp(pi)
	if err != nil {
		return 0, err
	}
	return pi, nil
}

// HundredOp calls singleOp a hundred times via tenOp.
func HundredOp(i int) (portOut int, portError error) {
	var err error
	pi := i

	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = tenOp(pi)
	if err != nil {
		return 0, err
	}
	return pi, nil
}

func tenOp(i int) (portOut int, portError error) {
	var err error
	pi := i

	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	pi, err = singleOp(pi)
	if err != nil {
		return 0, err
	}
	return pi, nil
}

func singleOp(i int) (portOut int, portError error) {
	if i < 0 || i > 1000000 {
		return 0, errors.New("should not happen")
	}
	i++
	return i, nil
}
