package constants

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
)

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}