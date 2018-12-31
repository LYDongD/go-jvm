package references

import (
	"gojvm/ch06/instructions/base"
	"gojvm/ch06/rtdata"
	"gojvm/ch06/rtdata/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	//弹出再压入，不改变栈的状态
	ref := stack.PopRef()
	stack.PushRef(ref)

	//null可以转任何类型
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
