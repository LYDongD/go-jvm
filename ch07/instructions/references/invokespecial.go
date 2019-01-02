package references

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//hack
func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef()
}
