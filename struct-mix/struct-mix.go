package main

import "fmt"

func main() {
	f := newFlow()
	f.SetErrorPort(func(err error) {
		fmt.Printf("ERROR: %s\n", err)
	})
	for i := 1; i < 20; i++ {
		f.InPort(i)
	}
}

type flow struct {
	op1 firstOp
	opU upperOp
	opL lowerOp
	// opZ lastOp

	InPort func(int)
}

func newFlow() *flow {
	f := &flow{}
	f.op1 = firstOp{}
	f.opU = upperOp{}
	f.opL = lowerOp{}
	opZ := lastOp{}

	f.op1.SetUpperPort(f.opU.InPort)
	f.op1.SetLowerPort(f.opL.InPort)
	f.opU.SetOutPort(opZ.UpperInPort)
	f.opL.SetOutPort(opZ.LowerInPort)

	f.InPort = f.op1.InPort
	return f
}
func (f *flow) SetErrorPort(port func(error)) {
	f.op1.SetErrorPort(port)
	f.opU.SetErrorPort(port)
	f.opL.SetErrorPort(port)
}

type firstOp struct {
	upperPort func(int)
	lowerPort func(int)
	errorPort func(error)
}

func (op *firstOp) InPort(i int) {
	err := failUtil(i)
	if err != nil {
		op.errorPort(err)
		return
	}
	if (i & 1) != 0 {
		op.upperPort(i)
	} else {
		op.lowerPort(i)
	}
}
func (op *firstOp) SetUpperPort(port func(int)) {
	op.upperPort = port
}
func (op *firstOp) SetLowerPort(port func(int)) {
	op.lowerPort = port
}
func (op *firstOp) SetErrorPort(port func(error)) {
	op.errorPort = port
}

type upperOp struct {
	outPort   func(int)
	errorPort func(error)
}

func (op *upperOp) InPort(i int) {
	fmt.Printf("Going upper: %d\n", i)
	err := failUtil(i)
	if err != nil {
		op.errorPort(err)
		return
	}
	i++ // do own work
	op.outPort(i)
}
func (op *upperOp) SetOutPort(port func(int)) {
	op.outPort = port
}
func (op *upperOp) SetErrorPort(port func(error)) {
	op.errorPort = port
}

type lowerOp struct {
	outPort   func(int)
	errorPort func(error)
}

func (op *lowerOp) InPort(i int) {
	fmt.Printf("Going lower: %d\n", i)
	err := failUtil(i)
	if err != nil {
		op.errorPort(err)
		return
	}
	i-- // do own work
	op.outPort(i)
}
func (op *lowerOp) SetOutPort(port func(int)) {
	op.outPort = port
}
func (op *lowerOp) SetErrorPort(port func(error)) {
	op.errorPort = port
}

type lastOp struct {
}

func (op *lastOp) UpperInPort(i int) {
	i--
	op.internalFunc(i)
}
func (op *lastOp) LowerInPort(i int) {
	i++
	op.internalFunc(i)
}
func (op *lastOp) internalFunc(i int) {
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
