package comparisons

import "gojvm/ch07/instructions/base"
import "gojvm/ch07/rtdata"

// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtdata.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtdata.Frame) {
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
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
