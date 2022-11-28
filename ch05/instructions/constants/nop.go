package constants

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

type NOP struct {
	// nothing to do
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
