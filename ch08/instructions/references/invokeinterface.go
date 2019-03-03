package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

type INVOKE_INTERFACE struct {
	//16(常量池索引) + 8（参数数量）+ 8（历史遗留）
	index uint
}


func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadInt16())
	reader.ReadUint8() //slot count 被调用方法的参数数量,历史遗留
	reader.ReadUint8() // must be 0 向后兼容
}

func (self *INVOKE_INTERFACE) Execute(frame *rtdata.Frame) {

	//定位方法，获取符号引用并解析成目标引用
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolveInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//获取接口方法(被调用方法)的实例引用(当前方法的操作数栈保存了被调用方法的参数slot list及被调用类的实例引用ref)
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	//方法实例必须是方法所属的类的实现类 todo
	//if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
	//	panic("java.lang.IncompatibleClassChangeError")
	//}

	//去实例中找目标方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}