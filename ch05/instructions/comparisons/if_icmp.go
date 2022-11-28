package comparisons

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Branch if int comparison with zero succeeds
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
