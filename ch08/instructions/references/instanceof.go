package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()

	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	}else {
		stack.PushInt(0)
	}
}
