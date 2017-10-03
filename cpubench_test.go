package alternatives

import (
	"testing"

	"github.com/flowdev/gflow-alternatives/classic"
)

// BenchmarkClassic1000 benchmarks the classical implementation of the thousandOp.
func BenchmarkClassic1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		classic.ThousandOp(0)
	}
}
