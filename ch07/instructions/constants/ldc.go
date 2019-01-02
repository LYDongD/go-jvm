package constants

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
)

type LDC struct {
	base.Index8Instruction
}

type LDC_W struct {
	base.Index16Instruction
}

type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC2_W) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	}
}


func _ldc(frame *rtdata.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	default:
		panic("todo: idc")
	}
}

