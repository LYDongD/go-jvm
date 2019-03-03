package math

import "gojvm/ch08/instructions/base"
import "gojvm/ch08/rtdata"

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (self *LOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
