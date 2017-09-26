package oopmix

import (
	"errors"
	"testing"
)

// BenchmarkOOPMix1000 benchmarks the implementation of the thousandOp that uses a mix of OOP things.
func BenchmarkOOPMix1000(b *testing.B) {
	thousandOp := newThousandOp()
	thousandOp.SetOutPort(func(i int) {})
	thousandOp.SetErrorPort(func(err error) {})
	for n := 0; n < b.N; n++ {
		thousandOp.InPort(0)
	}
}

type thousandOp struct {
	InPort func(int)
	op1    *hundredOp
	op2    *hundredOp
	op3    *hundredOp
	op4    *hundredOp
	op5    *hundredOp
	op6    *hundredOp
	op7    *hundredOp
	op8    *hundredOp
	op9    *hundredOp
	op10   *hundredOp
}

func newThousandOp() *thousandOp {
	o := &thousandOp{
		op1:  newHundredOp(),
		op2:  newHundredOp(),
		op3:  newHundredOp(),
		op4:  newHundredOp(),
		op5:  newHundredOp(),
		op6:  newHundredOp(),
		op7:  newHundredOp(),
		op8:  newHundredOp(),
		op9:  newHundredOp(),
		op10: newHundredOp(),
	}
	o.InPort = o.op1.InPort

	o.op1.SetOutPort(o.op2.InPort)
	o.op2.SetOutPort(o.op3.InPort)
	o.op3.SetOutPort(o.op4.InPort)
	o.op4.SetOutPort(o.op5.InPort)
	o.op5.SetOutPort(o.op6.InPort)
	o.op6.SetOutPort(o.op7.InPort)
	o.op7.SetOutPort(o.op8.InPort)
	o.op8.SetOutPort(o.op9.InPort)
	o.op9.SetOutPort(o.op10.InPort)

	return o
}
func (op *thousandOp) SetOutPort(port func(int)) {
	op.op10.SetOutPort(port)
}
func (op *thousandOp) SetErrorPort(port func(error)) {
	op.op1.SetErrorPort(port)
	op.op2.SetErrorPort(port)
	op.op3.SetErrorPort(port)
	op.op4.SetErrorPort(port)
	op.op5.SetErrorPort(port)
	op.op6.SetErrorPort(port)
	op.op7.SetErrorPort(port)
	op.op8.SetErrorPort(port)
	op.op9.SetErrorPort(port)
	op.op10.SetErrorPort(port)
}

type hundredOp struct {
	InPort func(int)
	op1    *tenOp
	op2    *tenOp
	op3    *tenOp
	op4    *tenOp
	op5    *tenOp
	op6    *tenOp
	op7    *tenOp
	op8    *tenOp
	op9    *tenOp
	op10   *tenOp
}

func newHundredOp() *hundredOp {
	o := &hundredOp{
		op1:  newTenOp(),
		op2:  newTenOp(),
		op3:  newTenOp(),
		op4:  newTenOp(),
		op5:  newTenOp(),
		op6:  newTenOp(),
		op7:  newTenOp(),
		op8:  newTenOp(),
		op9:  newTenOp(),
		op10: newTenOp(),
	}
	o.InPort = o.op1.InPort

	o.op1.SetOutPort(o.op2.InPort)
	o.op2.SetOutPort(o.op3.InPort)
	o.op3.SetOutPort(o.op4.InPort)
	o.op4.SetOutPort(o.op5.InPort)
	o.op5.SetOutPort(o.op6.InPort)
	o.op6.SetOutPort(o.op7.InPort)
	o.op7.SetOutPort(o.op8.InPort)
	o.op8.SetOutPort(o.op9.InPort)
	o.op9.SetOutPort(o.op10.InPort)

	return o
}
func (op *hundredOp) SetOutPort(port func(int)) {
	op.op10.SetOutPort(port)
}
func (op *hundredOp) SetErrorPort(port func(error)) {
	op.op1.SetErrorPort(port)
	op.op2.SetErrorPort(port)
	op.op3.SetErrorPort(port)
	op.op4.SetErrorPort(port)
	op.op5.SetErrorPort(port)
	op.op6.SetErrorPort(port)
	op.op7.SetErrorPort(port)
	op.op8.SetErrorPort(port)
	op.op9.SetErrorPort(port)
	op.op10.SetErrorPort(port)
}

type tenOp struct {
	InPort func(int)
	op1    *singleOp
	op2    *singleOp
	op3    *singleOp
	op4    *singleOp
	op5    *singleOp
	op6    *singleOp
	op7    *singleOp
	op8    *singleOp
	op9    *singleOp
	op10   *singleOp
}

func newTenOp() *tenOp {
	o := &tenOp{
		op1:  &singleOp{},
		op2:  &singleOp{},
		op3:  &singleOp{},
		op4:  &singleOp{},
		op5:  &singleOp{},
		op6:  &singleOp{},
		op7:  &singleOp{},
		op8:  &singleOp{},
		op9:  &singleOp{},
		op10: &singleOp{},
	}
	o.InPort = o.op1.InPort

	o.op1.SetOutPort(o.op2.InPort)
	o.op2.SetOutPort(o.op3.InPort)
	o.op3.SetOutPort(o.op4.InPort)
	o.op4.SetOutPort(o.op5.InPort)
	o.op5.SetOutPort(o.op6.InPort)
	o.op6.SetOutPort(o.op7.InPort)
	o.op7.SetOutPort(o.op8.InPort)
	o.op8.SetOutPort(o.op9.InPort)
	o.op9.SetOutPort(o.op10.InPort)

	return o
}
func (op *tenOp) SetOutPort(port func(int)) {
	op.op10.SetOutPort(port)
}
func (op *tenOp) SetErrorPort(port func(error)) {
	op.op1.SetErrorPort(port)
	op.op2.SetErrorPort(port)
	op.op3.SetErrorPort(port)
	op.op4.SetErrorPort(port)
	op.op5.SetErrorPort(port)
	op.op6.SetErrorPort(port)
	op.op7.SetErrorPort(port)
	op.op8.SetErrorPort(port)
	op.op9.SetErrorPort(port)
	op.op10.SetErrorPort(port)
}

type singleOp struct {
	outPort   func(int)
	errorPort func(error)
}

func (op *singleOp) InPort(i int) {
	if i < 0 || i > 1000000 {
		op.errorPort(errors.New("should not happen"))
		return
	}
	op.outPort(i + 1)
}
func (op *singleOp) SetOutPort(port func(int)) {
	op.outPort = port
}
func (op *singleOp) SetErrorPort(port func(error)) {
	op.errorPort = port
}
