package base

import (
	"awesomeProject/ch05/rtda"
)

type Instruction interface {
	// FetchOperands /** 从字节码中提取操作数 */
	FetchOperands(reader *BytecodeReader)
	// Execute /** 执行指令逻辑 */
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction /** 跳转指令 */
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction /** 存储和加载类指令都需要根据索引存取局部变量表 */
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction /** 用于访问运行时常量池的指令 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
