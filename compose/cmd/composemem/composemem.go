package main

import (
	"github.com/flowdev/gflow-alternatives/compose"
)

func main() {
	for i := 0; i < 1000; i++ {
		compose.ThousandOp(0)
	}
}
