package conversions

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}