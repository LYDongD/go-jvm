package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
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

