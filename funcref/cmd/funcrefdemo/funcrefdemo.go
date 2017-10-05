package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		err := firstOp(i)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}
}

func firstOp(i int) error {
	err := failUtil(i)
	if err != nil {
		return err
	}
	if (i & 1) != 0 {
		return upperOp(i)
	}
	return lowerOp(i)
}

func upperOp(i int) error {
	fmt.Printf("Going upper: %d\n", i)
	err := failUtil(i)
	if err != nil {
		return err
	}
	i++ // do own work
	lastUpperOp(i)
	return nil
}

func lowerOp(i int) error {
	fmt.Printf("Going lower: %d\n", i)
	err := failUtil(i)
	if err != nil {
		return err
	}
	i-- // do own work
	lastLowerOp(i)
	return nil
}

func lastUpperOp(i int) {
	i-- // do some work
	lastOp(i)
}
func lastLowerOp(i int) {
	i++ // do some work
	lastOp(i)
}
func lastOp(i int) {
	fmt.Printf("reached last op: %d\n", i)
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
