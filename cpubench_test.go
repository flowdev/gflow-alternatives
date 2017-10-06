package alternatives

import (
	"testing"

	"github.com/flowdev/gflow-alternatives/channel"
	"github.com/flowdev/gflow-alternatives/classic"
	"github.com/flowdev/gflow-alternatives/funcref"
	"github.com/flowdev/gflow-alternatives/lambdas"
	"github.com/flowdev/gflow-alternatives/oopmix"
	"github.com/flowdev/gflow-alternatives/rets"
)

// BenchmarkClassic1000 benchmarks the classical implementation of the
// thousandOp.
func BenchmarkClassic1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		classic.ThousandOp(0)
	}
}

// BenchmarkChannel1000 benchmarks the implementation of the thousandOp that
// uses a mix of OOP things.
func BenchmarkChannel1000(b *testing.B) {
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

	for n := 0; n < b.N; n++ {
		in <- 0
		<-out
	}
	close(in)
	<-done // wait for the ErrorOp
	close(done)
	close(err)
}

// BenchmarkFuncref1000 benchmarks the implementation of the thousandOp that
// uses a function reference.
func BenchmarkFuncref1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		funcref.ThousandOp(0)
	}
}

// BenchmarkLambdas1000 benchmarks the implementation of the thousandOp that
// uses lambdas.
func BenchmarkLambdas1000(b *testing.B) {
	thousandOp := lambdas.NewThousandOp(func(i int) {}, func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

// BenchmarkOOPMix1000 benchmarks the implementation of the thousandOp that
// uses a mix of OOP things.
func BenchmarkOOPMix1000(b *testing.B) {
	thousandOp := oopmix.NewThousandOp()
	thousandOp.SetOutPort(func(i int) {})
	thousandOp.SetErrorPort(func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp.InPort(0)
	}
}

// BenchmarkRets1000 benchmarks the implementation of the thousandOp that uses
// many return values.
func BenchmarkRets1000(b *testing.B) {
	thousandOp := rets.NewThousandOp()
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}
