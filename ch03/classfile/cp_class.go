package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = rader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
