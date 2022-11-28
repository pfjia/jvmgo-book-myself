package rtda

type Frame struct {
	lower *Frame
	/**
	 * 本地变量表指针
	 */
	localVars LocalVars
	/**
	 * 操作数栈指针
	 */
	operandStack *OperandStack

	thread *Thread

	/**
	 * 实现跳转指令新增
	 */
	nextPC int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
