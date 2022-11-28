package comparisons

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopRef()
	val1 := frame.OperandStack().PopRef()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopRef()
	val1 := frame.OperandStack().PopRef()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
