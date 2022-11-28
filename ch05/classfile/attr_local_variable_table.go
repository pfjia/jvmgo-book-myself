package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:   reader.readUint16(),
			length:    reader.readUint16(),
			nameIndex: reader.readUint16(),
			descIndex: reader.readUint16(),
			index:     reader.readUint16(),
		}
	}
}

type LocalVariableTableEntry struct {
	startPc   uint16
	length    uint16
	nameIndex uint16
	descIndex uint16
	index     uint16
}
