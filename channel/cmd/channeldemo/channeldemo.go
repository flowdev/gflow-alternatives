package main

import (
	"fmt"
	"sync"
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
	op12Err := make(chan error, 10)
	opU2Err := make(chan error, 10)
	opL2Err := make(chan error, 10)

	f.op1.Upper = op12U
	f.op1.Lower = op12L
	f.op1.Error = op12Err

	f.opU.In = op12U
	f.opU.Out = opU2Z
	f.opU.Error = opU2Err

	f.opL.In = op12L
	f.opL.Out = opL2Z
	f.opL.Error = opL2Err

	f.opZ.UpperIn = opU2Z
	f.opZ.LowerIn = opL2Z

	f.opErr.Error1 = op12Err
	f.opErr.Error2 = opU2Err
	f.opErr.Error3 = opL2Err

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
	f.opErr.Run()
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
				close(op.Error)
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
				close(op.Error)
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
				close(op.Error)
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
	Error1 <-chan error
	Error2 <-chan error
	Error3 <-chan error
	Done   chan<- bool
}

func (op *errorOp) Run() {
	in := mergeErrors(op.Error1, op.Error2, op.Error3)

	go func() {
		for err := range in {
			fmt.Printf("ERROR: %s\n", err)
		}
		op.Done <- true
	}()
}

func mergeErrors(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
