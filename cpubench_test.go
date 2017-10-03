package alternatives

import (
	"testing"

	"github.com/flowdev/gflow-alternatives/classic"
	"github.com/flowdev/gflow-alternatives/lambdas"
	"github.com/flowdev/gflow-alternatives/oopmix"
)

// BenchmarkClassic1000 benchmarks the classical implementation of the thousandOp.
func BenchmarkClassic1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		classic.ThousandOp(0)
	}
}

// BenchmarkLambdas1000 benchmarks the implementation of the thousandOp that uses lambdas.
func BenchmarkLambdas1000(b *testing.B) {
	thousandOp := lambdas.NewThousandOp(func(i int) {}, func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp(0)
	}
}

// BenchmarkOOPMix1000 benchmarks the implementation of the thousandOp that uses a mix of OOP things.
func BenchmarkOOPMix1000(b *testing.B) {
	thousandOp := oopmix.NewThousandOp()
	thousandOp.SetOutPort(func(i int) {})
	thousandOp.SetErrorPort(func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp.InPort(0)
	}
}
