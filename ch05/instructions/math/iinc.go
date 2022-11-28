package math

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetInt(self.Index)
	val += self.Const
	frame.LocalVars().SetInt(self.Index, val)
}
