package conversions

import "gojvm/ch05/instructions/base"
import "gojvm/ch05/rtdata"

// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
