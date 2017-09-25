package main

import "fmt"

func main() {
	flow := newFlow()
	for i := 1; i < 20; i++ {
		err := flow(i)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}
}

func newFlow() func(int) error {
	op1In := newFirstOp()
	opUIn := newUpperOp()
	opLIn := newLowerOp()
	opZuIn, opZlIn := newLastOp()

	return func(i int) error {
		upperI, lowerI, err := op1In(i)
		if err != nil {
			return err
		} else if lowerI != 0 {
			i, err = opLIn(lowerI)
		} else {
			i, err = opUIn(upperI)
		}
		if err != nil {
			return err
		}
		if lowerI != 0 {
			opZlIn(i)
		} else {
			opZuIn(i)
		}
		return nil
	}
}
func newFirstOp() func(int) (int, int, error) {
	return func(i int) (upperI, lowerI int, err error) {
		err = failUtil(i)
		if err != nil {
			return
		}
		if (i & 1) != 0 {
			upperI = i
		} else {
			lowerI = i
		}
		return
	}
}

func newUpperOp() func(int) (int, error) {
	return func(i int) (int, error) {
		fmt.Printf("Going upper: %d\n", i)
		err := failUtil(i)
		if err != nil {
			return 0, err
		}
		i++ // do own work
		return i, nil
	}
}

func newLowerOp() func(int) (int, error) {
	return func(i int) (int, error) {
		fmt.Printf("Going lower: %d\n", i)
		err := failUtil(i)
		if err != nil {
			return 0, err
		}
		i-- // do own work
		return i, nil
	}
}

func newLastOp() (upperIn, lowerIn func(int)) {
	internalFunc := func(i int) {
		fmt.Printf("reached last op: %d\n", i)
	}
	upperIn = func(i int) {
		i-- // do some work
		internalFunc(i)
	}
	lowerIn = func(i int) {
		i++ // do some work
		internalFunc(i)
	}
	return
}

var failNumber, failCount int

func failUtil(i int) error {
	if failNumber != i {
		failNumber = i
		failCount = 0
	}
	failCount++

	if failCount == (i % 3) {
		return fmt.Errorf("Its time for an error: %d", i)
	}
	return nil
}
