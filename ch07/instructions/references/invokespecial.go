package references

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}


func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	//如果方法是构造函数，则构造方法声明的类必须是从方法引用解析出来的类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	//弹出this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointException")
	}

	//todo 这一段权限检查没看懂
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass)
	&& resolvedMethod.Class().GetPackageName() != currentClass.getPackageName()
	&& ref.Class() != currentClass &&!ref.Class().IsSubClass(currentClass) {
		panic("java.lang.IllegalAccessError")
	}




}
