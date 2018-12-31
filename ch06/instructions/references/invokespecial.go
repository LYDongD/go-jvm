package references

import (
	"gojvm/ch06/instructions/base"
	"gojvm/ch06/rtdata"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//hack
func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef()
}
