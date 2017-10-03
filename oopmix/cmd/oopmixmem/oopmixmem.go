package main

import "github.com/flowdev/gflow-alternatives/oopmix"

func main() {
	thousandOp := oopmix.NewThousandOp()
	thousandOp.SetOutPort(func(i int) {})
	thousandOp.SetErrorPort(func(err error) {})
	for i := 0; i < 1000; i++ {
		thousandOp.InPort(0)
	}
}
