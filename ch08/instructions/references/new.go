package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//构造对象，操作数为类对象引用索引
type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtdata.Frame) {
	//去常量池获取类对象引用索引
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)

	//解析类
	class := classRef.ResolvedClass()

	//初始化类，所以new指令会触发类的初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	//检查
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	//构造对象实例，初始化实例变量
	ref := class.NewObject()

	//压入当前栈帧操作数栈
	frame.OperandStack().PushRef(ref)
}










