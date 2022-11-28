package classfile

type CodeAttribute struct {
	cp ConstantPool
	/**
	 * 操作数栈的最大深度
	 */
	maxStack uint16

	/**
	 * 局部变量表的最大容量
	 */
	maxLocals uint16

	/**
	 * 字节码
	 */
	code []byte

	/**
	 * 异常处理表
	 */
	exceptionTable []*ExceptionTableEntry

	/**
	 * 属性表
	 */
	attributes []AttributeInfo
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	self.code = reader.readBytes(reader.readUint32())
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

type ExceptionTableEntry struct {
	/**
	 * 起始pc
	 */
	startPc uint16

	/**
	 * 结束pc
	 */
	endPc uint16

	/**
	 * 异常处理器pc
	 */
	handlerPc uint16

	/**
	 * 捕获的异常类型
	 */
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}