package references

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
)

//静态方法调用
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()


	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	class := resolvedMethod.Class()
	//如果类还未初始化，则需要先初始化
	if !class.InitStarted() {
		//重置PC
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	if !resolvedMethod.IsStatic() {
		panic("java.lang.incompatibleClassChangeError")
	}

	base.InvokeMethod(frame, resolvedMethod)
}
