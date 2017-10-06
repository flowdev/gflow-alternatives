package main

import (
	"fmt"
	"io"
)

func main() {
	f := newFlow()
	done := make(chan bool, 3)
	in := make(chan int, 10)
	f.SetDone(done)
	f.SetIn(in)
	f.Run()
	for i := 1; i < 20; i++ {
		in <- i
	}
	close(in)
	<-done // wait for the upperIn of lastOp ...
	<-done // wait for the lowerIn of lastOp ...
	<-done // ... and the errorOp
	close(done)
}

type flow struct {
	op1   firstOp
	opU   upperOp
	opL   lowerOp
	opZ   lastOp
	opErr errorOp
}

func newFlow() *flow {
	f := &flow{
		op1:   firstOp{},
		opU:   upperOp{},
		opL:   lowerOp{},
		opZ:   lastOp{},
		opErr: errorOp{},
	}

	op12U := make(chan int, 10)
	op12L := make(chan int, 10)
	opU2Z := make(chan int, 10)
	opL2Z := make(chan int, 10)
	opX2Err := make(chan error, 10)

	f.op1.Upper = op12U
	f.op1.Lower = op12L
	f.op1.Error = opX2Err

	f.opU.In = op12U
	f.opU.Out = opU2Z
	f.opU.Error = opX2Err

	f.opL.In = op12L
	f.opL.Out = opL2Z
	f.opL.Error = opX2Err

	f.opZ.UpperIn = opU2Z
	f.opZ.LowerIn = opL2Z

	f.opErr.Error = opX2Err

	return f
}
func (f *flow) SetIn(port <-chan int) {
	f.op1.In = port
}
func (f *flow) SetDone(port chan<- bool) {
	f.opZ.Done = port
	f.opErr.Done = port
}
func (f *flow) Run() {
	f.opErr.Run(3)
	f.opZ.RunUpperIn()
	f.opZ.RunLowerIn()
	f.opU.Run()
	f.opL.Run()
	f.op1.Run()
}

type firstOp struct {
	In    <-chan int
	Upper chan<- int
	Lower chan<- int
	Error chan<- error
}

func (op *firstOp) Run() {
	go func() {
		for {
			i, ok := <-op.In
			if !ok {
				close(op.Upper)
				close(op.Lower)
				op.Error <- io.EOF // signal that we are done
				return
			}
			err := failUtil(i)
			if err != nil {
				op.Error <- err
				continue
			}
			if (i & 1) != 0 {
				op.Upper <- i
			} else {
				op.Lower <- i
			}
		}
	}()
}

type upperOp struct {
	In    <-chan int
	Out   chan<- int
	Error chan<- error
}

func (op *upperOp) Run() {
	go func() {
		for {
			i, ok := <-op.In
			if !ok {
				close(op.Out)
				op.Error <- io.EOF // signal that we are done
				return
			}
			fmt.Printf("Going upper: %d\n", i)
			err := failUtil(i)
			if err != nil {
				op.Error <- err
				continue
			}
			i++ // do own work
			op.Out <- i
		}
	}()
}

type lowerOp struct {
	In    <-chan int
	Out   chan<- int
	Error chan<- error
}

func (op *lowerOp) Run() {
	go func() {
		for {
			i, ok := <-op.In
			if !ok {
				close(op.Out)
				op.Error <- io.EOF // signal that we are done
				return
			}
			fmt.Printf("Going lower: %d\n", i)
			err := failUtil(i)
			if err != nil {
				op.Error <- err
				continue
			}
			i-- // do own work
			op.Out <- i
		}
	}()
}

type lastOp struct {
	UpperIn <-chan int
	LowerIn <-chan int
	Done    chan<- bool
}

func (op *lastOp) RunUpperIn() {
	go func() {
		for {
			i, ok := <-op.UpperIn
			if !ok {
				op.Done <- true
				return
			}
			i--
			op.internalFunc(i)
		}
	}()
}
func (op *lastOp) RunLowerIn() {
	go func() {
		for {
			i, ok := <-op.LowerIn
			if !ok {
				op.Done <- true
				return
			}
			i++
			op.internalFunc(i)
		}
	}()
}
func (op *lastOp) internalFunc(i int) {
	fmt.Printf("reached last op: %d\n", i)
}

type errorOp struct {
	Error <-chan error
	Done  chan<- bool
}

func (op *errorOp) Run(n int) {
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
