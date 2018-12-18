package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopSlot()
}

func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
