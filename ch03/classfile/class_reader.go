package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

//read u1
func (self *ClassReader) readUinit8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//read u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//read u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndiam.Uint32
	self.data = self.data[4:]
	return val
}

//read u8
func (self *ClassReader) readUini64() uint64 {
	val := binary.BigEndiam.Uint64
	self.data = self.data[8:]
	return val
}

//read table
func (self *ClassReader) readUinit16s() []uint16 {
	//read size from the begin u2
	n := self.readUinit16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return
}

//read bytes
func (self *ClassReader) readBytes(n uint32) []byte {
	val := self.data[:n]
	self.data = self.data[n:]
	return val
}
