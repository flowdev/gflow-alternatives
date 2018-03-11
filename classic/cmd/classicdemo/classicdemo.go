package main

import (
	"fmt"

	"github.com/flowdev/gflow-alternatives/util"
)

func main() {
	for i := 1; i < 20; i++ {
		err := firstOp(i)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}
}

func firstOp(i int) error {
	err := util.FailUtil(i)
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
	err := util.FailUtil(i)
	if err != nil {
		return err
	}
	i++ // do own work
	lastUpperOp(i)
	return nil
}

func lowerOp(i int) error {
	fmt.Printf("Going lower: %d\n", i)
	err := util.FailUtil(i)
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
