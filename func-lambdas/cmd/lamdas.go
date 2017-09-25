package main

import "fmt"

func main() {
	flow := newFlow(func(err error) {
		fmt.Printf("ERROR: %s\n", err)
	})
	for i := 1; i < 20; i++ {
		flow(i)
	}
}

func newFlow(handleError func(error)) func(int) {
	opZuIn, opZlIn := newLastOp()
	opLIn := newLowerOp(opZlIn, handleError)
	opUIn := newUpperOp(opZuIn, handleError)
	op1In := newFirstOp(opUIn, opLIn, handleError)

	return op1In
}
func newFirstOp(upperOp, lowerOp func(int), handleError func(error)) func(int) {
	return func(i int) {
		err := failUtil(i)
		if err != nil {
			handleError(err)
			return
		}
		if (i & 1) != 0 {
			upperOp(i)
		} else {
			lowerOp(i)
		}
	}
}

func newUpperOp(outOp func(int), handleError func(error)) func(int) {
	return func(i int) {
		fmt.Printf("Going upper: %d\n", i)
		err := failUtil(i)
		if err != nil {
			handleError(err)
			return
		}
		i++ // do own work
		outOp(i)
	}
}

func newLowerOp(outOp func(int), handleError func(error)) func(int) {
	return func(i int) {
		fmt.Printf("Going lower: %d\n", i)
		err := failUtil(i)
		if err != nil {
			handleError(err)
			return
		}
		i-- // do own work
		outOp(i)
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
