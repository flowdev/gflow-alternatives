package main

import "github.com/flowdev/gflow-alternatives/funcref"

func main() {
	for i := 0; i < 1000; i++ {
		funcref.ThousandOp(0)
	}
}
