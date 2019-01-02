package loads

import "gojvm/ch07/instructions/base"
import "gojvm/ch07/rtdata"

// Load int from local variable
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtdata.Frame) {
	_iload(frame, self.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *rtdata.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtdata.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtdata.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtdata.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
