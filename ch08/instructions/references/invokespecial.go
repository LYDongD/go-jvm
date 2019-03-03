package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//适用于以下方法的调用: <init>方法；私有方法；super.method()
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}


func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	//调用类
	currentClass := frame.Method().Class()
	//调用类的常量池
	cp := currentClass.ConstantPool()
	//操作数为方法符号引用在常量池的索引
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	//分别解析类和方法
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	//如果方法是构造函数，则构造方法声明的类必须是从方法引用解析出来的类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	//目标方法不能是静态的
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//当前方法的操作数栈包括this引用和参数列表
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointException")
	}

	//todo 权限检查?
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}


	//处理super调用的情况(调用超类的非构造函数），需要进一步去父类查找目标方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)


}
