package control

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Branch always
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
