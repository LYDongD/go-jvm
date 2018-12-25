package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type DCOMPG struct {
	base.NoOperandsInstruction
}

func (self *DCOMPG) Execute(frame *rtdata.Frame) {
	_dcmp(frame, true)
}


type DCOMPL struct {
	base.NoOperandsInstruction
}


func (self *DCOMPL) Execute(frame *rtdata.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	}else if v1 == v2 {
		stack.PushInt(0)
	}else if v1 < v2 {
		stack.PushInt(-1)
	}else if gFlag {
		stack.PushInt(1)
	}else {
		stack.PushInt(-1)
	}
}