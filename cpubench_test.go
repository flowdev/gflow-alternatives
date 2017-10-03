package alternatives

import (
	"testing"

	"github.com/flowdev/gflow-alternatives/classic"
	"github.com/flowdev/gflow-alternatives/lambdas"
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
