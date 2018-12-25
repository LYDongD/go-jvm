package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP_X1 struct {
	base.NoOperandsInstruction
}
type DUP_X2 struct {
	base.NoOperandsInstruction
}

type DUP2 struct {
	base.NoOperandsInstruction
}
type DUP2_X1 struct {
	base.NoOperandsInstruction
}
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

