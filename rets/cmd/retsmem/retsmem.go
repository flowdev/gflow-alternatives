package main

import "github.com/flowdev/gflow-alternatives/rets"

func main() {
	thousandOp := rets.NewThousandOp()
	for i := 0; i < 1000; i++ {
		thousandOp(0)
	}
}
