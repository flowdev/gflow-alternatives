package main

import (
	"github.com/flowdev/gflow-alternatives/lambdas"
)

func main() {
	thousandOp := lambdas.NewThousandOp(func(i int) {}, func(err error) {})
	for i := 0; i < 1000; i++ {
		thousandOp(0)
	}
}
