package main

import (
	"fmt"

	"github.com/flowdev/gflow-alternatives/util"
)

func main() {
	for i := 1; i < 20; i++ {
		err := flow(i)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}
}

func flow(i int) error {
	var strI *struct{ i int }
	var strErr *struct{ err error }

	upperI, lowerI, err := firstOp(i)
	if err != nil {
		return err
	} else if lowerI != nil {
		strI, strErr = lowerOp(*lowerI)
	} else {
		strI, strErr = upperOp(*upperI)
	}
	if strErr != nil {
		return strErr.err
	}
	if lowerI != nil {
		lastOpPortLower(strI.i)
	} else {
		lastOpPortUpper(strI.i)
	}
	return nil
}

func firstOp(i int) (portUpper *int, portLower *int, portError error) {
	err := util.FailUtil(i)
	if err != nil {
		return nil, nil, err
	}
	if (i & 1) != 0 {
		return &i, nil, nil
	}
	return nil, &i, nil
}

func upperOp(i int) (portOut *struct{ i int }, portError *struct{ err error }) {
	fmt.Printf("Going upper: %d\n", i)
	err := util.FailUtil(i)
	if err != nil {
		return nil, &struct{ err error }{err: err}
	}
	i++ // do own work
	return &struct{ i int }{i: i}, nil
}

func lowerOp(i int) (portOut *struct{ i int }, portError *struct{ err error }) {
	fmt.Printf("Going lower: %d\n", i)
	err := util.FailUtil(i)
	if err != nil {
		return nil, &struct{ err error }{err: err}
	}
	i-- // do own work
	return &struct{ i int }{i: i}, nil
}

func lastOpPortUpper(i int) {
	i-- // do some work
	lastOp(i)
}
func lastOpPortLower(i int) {
	i++ // do some work
	lastOp(i)
}
func lastOp(i int) {
	fmt.Printf("reached last op: %d\n", i)
}
