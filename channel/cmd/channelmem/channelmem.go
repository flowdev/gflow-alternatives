package main

import "github.com/flowdev/gflow-alternatives/channel"

func main() {
	thousandOp := channel.NewThousandOp()
	errOp := &channel.ErrorOp{}
	in := make(chan int, 10)
	out := make(chan int, 10)
	err := make(chan error, 10)
	done := make(chan bool, 1)
	thousandOp.SetIn(in)
	thousandOp.SetOut(out)
	thousandOp.SetError(err)
	errOp.Error = err
	errOp.Done = done

	errOp.Run(1000)
	thousandOp.Run()

	for i := 0; i < 1000; i++ {
		in <- 0
		<-out
	}
	close(in)
	<-done // wait for the ErrorOp
	close(done)
	close(err)
}
