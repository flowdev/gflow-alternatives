package channel

import (
	"errors"
	"fmt"
	"io"
)

// ThousandOp holds the data necessary to call ten HunredOps.
type ThousandOp struct {
	op1  *HundredOp
	op2  *HundredOp
	op3  *HundredOp
	op4  *HundredOp
	op5  *HundredOp
	op6  *HundredOp
	op7  *HundredOp
	op8  *HundredOp
	op9  *HundredOp
	op10 *HundredOp
}

// NewThousandOp returns a new ThousandOp.
func NewThousandOp() *ThousandOp {
	o := &ThousandOp{
		op1:  NewHundredOp(),
		op2:  NewHundredOp(),
		op3:  NewHundredOp(),
		op4:  NewHundredOp(),
		op5:  NewHundredOp(),
		op6:  NewHundredOp(),
		op7:  NewHundredOp(),
		op8:  NewHundredOp(),
		op9:  NewHundredOp(),
		op10: NewHundredOp(),
	}

	op1To2 := make(chan int, 10)
	op2To3 := make(chan int, 10)
	op3To4 := make(chan int, 10)
	op4To5 := make(chan int, 10)
	op5To6 := make(chan int, 10)
	op6To7 := make(chan int, 10)
	op7To8 := make(chan int, 10)
	op8To9 := make(chan int, 10)
	op9To10 := make(chan int, 10)

	o.op1.SetOut(op1To2)

	o.op2.SetIn(op1To2)
	o.op2.SetOut(op2To3)

	o.op3.SetIn(op2To3)
	o.op3.SetOut(op3To4)

	o.op4.SetIn(op3To4)
	o.op4.SetOut(op4To5)

	o.op5.SetIn(op4To5)
	o.op5.SetOut(op5To6)

	o.op6.SetIn(op5To6)
	o.op6.SetOut(op6To7)

	o.op7.SetIn(op6To7)
	o.op7.SetOut(op7To8)

	o.op8.SetIn(op7To8)
	o.op8.SetOut(op8To9)

	o.op9.SetIn(op8To9)
	o.op9.SetOut(op9To10)

	o.op10.SetIn(op9To10)

	return o
}

// SetIn sets the input port of the ThousandOp.
func (op *ThousandOp) SetIn(port <-chan int) {
	op.op1.SetIn(port)
}

// SetOut sets the output port of the ThousandOp.
func (op *ThousandOp) SetOut(port chan<- int) {
	op.op10.SetOut(port)
}

// SetError sets the error port of the ThousandOp.
func (op *ThousandOp) SetError(port chan<- error) {
	op.op1.SetError(port)
	op.op2.SetError(port)
	op.op3.SetError(port)
	op.op4.SetError(port)
	op.op5.SetError(port)
	op.op6.SetError(port)
	op.op7.SetError(port)
	op.op8.SetError(port)
	op.op9.SetError(port)
	op.op10.SetError(port)
}

// Run the ThousandOp or rather all of its internal operations.
func (op *ThousandOp) Run() {
	op.op1.Run()
	op.op2.Run()
	op.op3.Run()
	op.op4.Run()
	op.op5.Run()
	op.op6.Run()
	op.op7.Run()
	op.op8.Run()
	op.op9.Run()
	op.op10.Run()
}

// HundredOp holds the data necessary to call ten tenOps.
type HundredOp struct {
	op1  *tenOp
	op2  *tenOp
	op3  *tenOp
	op4  *tenOp
	op5  *tenOp
	op6  *tenOp
	op7  *tenOp
	op8  *tenOp
	op9  *tenOp
	op10 *tenOp
}

// NewHundredOp returns a new HundredOp.
func NewHundredOp() *HundredOp {
	o := &HundredOp{
		op1:  newTenOp(),
		op2:  newTenOp(),
		op3:  newTenOp(),
		op4:  newTenOp(),
		op5:  newTenOp(),
		op6:  newTenOp(),
		op7:  newTenOp(),
		op8:  newTenOp(),
		op9:  newTenOp(),
		op10: newTenOp(),
	}

	op1To2 := make(chan int, 10)
	op2To3 := make(chan int, 10)
	op3To4 := make(chan int, 10)
	op4To5 := make(chan int, 10)
	op5To6 := make(chan int, 10)
	op6To7 := make(chan int, 10)
	op7To8 := make(chan int, 10)
	op8To9 := make(chan int, 10)
	op9To10 := make(chan int, 10)

	o.op1.SetOut(op1To2)

	o.op2.SetIn(op1To2)
	o.op2.SetOut(op2To3)

	o.op3.SetIn(op2To3)
	o.op3.SetOut(op3To4)

	o.op4.SetIn(op3To4)
	o.op4.SetOut(op4To5)

	o.op5.SetIn(op4To5)
	o.op5.SetOut(op5To6)

	o.op6.SetIn(op5To6)
	o.op6.SetOut(op6To7)

	o.op7.SetIn(op6To7)
	o.op7.SetOut(op7To8)

	o.op8.SetIn(op7To8)
	o.op8.SetOut(op8To9)

	o.op9.SetIn(op8To9)
	o.op9.SetOut(op9To10)

	o.op10.SetIn(op9To10)

	return o
}

// SetIn sets the input port of the HundredOp.
func (op *HundredOp) SetIn(port <-chan int) {
	op.op1.SetIn(port)
}

// SetOut sets the output port of the HundredOp.
func (op *HundredOp) SetOut(port chan<- int) {
	op.op10.SetOut(port)
}

// SetError sets the error port of the HundredOp.
func (op *HundredOp) SetError(port chan<- error) {
	op.op1.SetError(port)
	op.op2.SetError(port)
	op.op3.SetError(port)
	op.op4.SetError(port)
	op.op5.SetError(port)
	op.op6.SetError(port)
	op.op7.SetError(port)
	op.op8.SetError(port)
	op.op9.SetError(port)
	op.op10.SetError(port)
}

// Run the HundredOp or rather all of its internal operations.
func (op *HundredOp) Run() {
	op.op1.Run()
	op.op2.Run()
	op.op3.Run()
	op.op4.Run()
	op.op5.Run()
	op.op6.Run()
	op.op7.Run()
	op.op8.Run()
	op.op9.Run()
	op.op10.Run()
}

type tenOp struct {
	op1  *singleOp
	op2  *singleOp
	op3  *singleOp
	op4  *singleOp
	op5  *singleOp
	op6  *singleOp
	op7  *singleOp
	op8  *singleOp
	op9  *singleOp
	op10 *singleOp
}

func newTenOp() *tenOp {
	o := &tenOp{
		op1:  &singleOp{},
		op2:  &singleOp{},
		op3:  &singleOp{},
		op4:  &singleOp{},
		op5:  &singleOp{},
		op6:  &singleOp{},
		op7:  &singleOp{},
		op8:  &singleOp{},
		op9:  &singleOp{},
		op10: &singleOp{},
	}

	op1To2 := make(chan int, 10)
	op2To3 := make(chan int, 10)
	op3To4 := make(chan int, 10)
	op4To5 := make(chan int, 10)
	op5To6 := make(chan int, 10)
	op6To7 := make(chan int, 10)
	op7To8 := make(chan int, 10)
	op8To9 := make(chan int, 10)
	op9To10 := make(chan int, 10)

	o.op1.Out = op1To2

	o.op2.In = op1To2
	o.op2.Out = op2To3

	o.op3.In = op2To3
	o.op3.Out = op3To4

	o.op4.In = op3To4
	o.op4.Out = op4To5

	o.op5.In = op4To5
	o.op5.Out = op5To6

	o.op6.In = op5To6
	o.op6.Out = op6To7

	o.op7.In = op6To7
	o.op7.Out = op7To8

	o.op8.In = op7To8
	o.op8.Out = op8To9

	o.op9.In = op8To9
	o.op9.Out = op9To10

	o.op10.In = op9To10

	return o
}
func (op *tenOp) SetIn(port <-chan int) {
	op.op1.In = port
}
func (op *tenOp) SetOut(port chan<- int) {
	op.op10.Out = port
}
func (op *tenOp) SetError(port chan<- error) {
	op.op1.Error = port
	op.op2.Error = port
	op.op3.Error = port
	op.op4.Error = port
	op.op5.Error = port
	op.op6.Error = port
	op.op7.Error = port
	op.op8.Error = port
	op.op9.Error = port
	op.op10.Error = port
}

func (op *tenOp) Run() {
	op.op1.Run()
	op.op2.Run()
	op.op3.Run()
	op.op4.Run()
	op.op5.Run()
	op.op6.Run()
	op.op7.Run()
	op.op8.Run()
	op.op9.Run()
	op.op10.Run()
}

type singleOp struct {
	In    <-chan int
	Out   chan<- int
	Error chan<- error
}

func (op *singleOp) Run() {
	go func() {
		for {
			i, ok := <-op.In
			if !ok {
				close(op.Out)
				op.Error <- io.EOF // signal that we are done
				return
			}
			if i < 0 || i > 1000000 {
				op.Error <- errors.New("should not happen")
				continue
			}
			op.Out <- (i + 1)
		}
	}()
}

// ErrorOp handles errors by reporting them.
type ErrorOp struct {
	Error <-chan error
	Done  chan<- bool
}

// Run runs the ErrorOp with input from n components.
func (op *ErrorOp) Run(n int) {
	i := 0

	go func() {
		for err := range op.Error {
			if err != io.EOF {
				fmt.Printf("ERROR: %s\n", err)
			} else {
				i++
				if i >= n {
					break
				}
			}
		}
		op.Done <- true
	}()
}
