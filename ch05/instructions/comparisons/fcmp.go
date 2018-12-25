package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type FCOMPG struct {
	base.NoOperandsInstruction
}

func (self *FCOMPG) Execute(frame *rtdata.Frame) {
	_fcmp(frame, true)
}

type FCOMPL struct {
	base.NoOperandsInstruction
}

func (self *FCOMPL) Execute(frame *rtdata.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag { //if v1 or v2 is NaN
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
