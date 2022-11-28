package math

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Multiply double
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	stack.PushDouble(val1 * val2)
}

// Multiply float
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	stack.PushFloat(val1 * val2)
}

// Multiply int
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	stack.PushInt(val1 * val2)
}

// Multiply long
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	stack.PushLong(val1 * val2)
}

// Path: ch05/instructions/math/mul.go
