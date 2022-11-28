package control

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	/**
	 * 存放 npairs 个成对的值，每对值中的第一个值是 case 值，第二个值是跳转偏移量
	 */
	matchOffsets []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
