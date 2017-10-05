package alternatives

import (
	"testing"

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
