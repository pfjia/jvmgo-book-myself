package control

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/rtda"
)

// Access jump table by index and jump
type TABLE_SWITCH struct {
	/**
	 * 默认情况下执行跳转所需的字节码偏移量
	 */
	defaultOffset int32

	/**
	 * low 和 high 对应 case 的取值范围
	 */
	low  int32
	high int32
	/**
	 * 索引表，里面存放 high-low+1 个int值，对应 low-high 范围内的case值
	 */
	jumpOffsets []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

// Execute
/**
 * 1. 从操作数栈中弹出一个int变量
 * 2. 判断该变量是否在 low 和 high 给定的范围之间
 * 3. 如果在范围之内，按照索引表跳转
 * 4. 如果不在范围之内，按照默认偏移量跳转
 */
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
