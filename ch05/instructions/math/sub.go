package math

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Subtract double
type DSUB struct {
	base.NoOperandsInstruction
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	stack.PushDouble(val1 - val2)
}

// Subtract float
type FSUB struct {
	base.NoOperandsInstruction
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	stack.PushFloat(val1 - val2)
}

// Subtract int
type ISUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	stack.PushInt(val1 - val2)
}

// Subtract long
type LSUB struct {
	base.NoOperandsInstruction
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	stack.PushLong(val1 - val2)
}

// Path: ch05/instructions/math/sub.go

