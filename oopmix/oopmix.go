package oopmix

import (
	"errors"
)

// ThousandOp holds the data necessary to call ten HunredOps.
type ThousandOp struct {
	InPort func(int)
	op1    *HundredOp
	op2    *HundredOp
	op3    *HundredOp
	op4    *HundredOp
	op5    *HundredOp
	op6    *HundredOp
	op7    *HundredOp
	op8    *HundredOp
	op9    *HundredOp
	op10   *HundredOp
}

// NewThousandOp returns a new ThousandOp.
func NewThousandOp() *ThousandOp {
	o := &ThousandOp{
		op1:  NewHundredOp(),
		op2:  NewHundredOp(),
		op3:  NewHundredOp(),
		op4:  NewHundredOp(),
		op5:  NewHundredOp(),
		op6:  NewHundredOp(),
		op7:  NewHundredOp(),
		op8:  NewHundredOp(),
		op9:  NewHundredOp(),
		op10: NewHundredOp(),
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

// SetOutPort sets the output port of the ThousandOp.
func (op *ThousandOp) SetOutPort(port func(int)) {
	op.op10.SetOutPort(port)
}

// SetErrorPort sets the error port of the ThousandOp.
func (op *ThousandOp) SetErrorPort(port func(error)) {
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

// HundredOp holds the data necessary to call ten tenOps.
type HundredOp struct {
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

// NewHundredOp returns a new HundredOp.
func NewHundredOp() *HundredOp {
	o := &HundredOp{
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

// SetOutPort sets the output port of the HundredOp.
func (op *HundredOp) SetOutPort(port func(int)) {
	op.op10.SetOutPort(port)
}

// SetErrorPort sets the error port of the HundredOp.
func (op *HundredOp) SetErrorPort(port func(error)) {
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
