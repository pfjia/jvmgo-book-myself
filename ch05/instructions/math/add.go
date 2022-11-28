package math

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Add double
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	stack.PushDouble(val1 + val2)
}

// Add float
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	stack.PushFloat(val1 + val2)
}

// Add int
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	stack.PushInt(val1 + val2)
}

// Add long
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	stack.PushLong(val1 + val2)
}

// Path: ch05/instructions/math/add.go
