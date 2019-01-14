package references

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.incompatibleClassChangeError")
	}

	base.InvokeMethod(frame, resolvedMethod)
}
