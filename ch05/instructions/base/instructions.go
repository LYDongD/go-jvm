package base

import (
	"go-jvm/ch05/rtdata"
)

type Instructions interface {
	//fetch operands from byte code using byte reader
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtdata.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing todo
}

//goto instruction
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//local variables index instruction
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
}

//constant pool index struction
type index16Instruction struct {
	Index uint
}

func (self *index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}


