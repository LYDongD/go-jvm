package comparisons

import "gojvm/ch08/instructions/base"
import "gojvm/ch08/rtdata"

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
