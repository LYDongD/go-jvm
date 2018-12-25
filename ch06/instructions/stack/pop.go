package stack

import "gojvm/ch06/instructions/base"
import "gojvm/ch06/rtdata"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
