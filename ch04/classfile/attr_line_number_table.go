package classfile

type LineNumberTableAttribute struct {
	lineNumberTabale []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTabale = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTabale {
		self.lineNumberTabale[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
