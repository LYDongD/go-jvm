package references

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//取出被调用类的引用，如果为空，则为空指针异常
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		//hack
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}

	//调用不在一个包内的父类的protected方法,则调用类引用必须是当前类或当前类的子类
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		ref.Class().IsSubClassOf(currentClass) {
			panic("java.lang.IllegalAccessorError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtdata.OperandStack, descriptor string) {
	switch descriptor{
	case"(Z)V":fmt.Printf("%v\n",stack.PopInt()!=0)
	case"(C)V":fmt.Printf("%c\n",stack.PopInt())
	case"(B)V":fmt.Printf("%v\n",stack.PopInt())
	case"(S)V":fmt.Printf("%v\n",stack.PopInt())
	case"(I)V":fmt.Printf("%v\n",stack.PopInt())
	case"(F)V":fmt.Printf("%v\n",stack.PopFloat())
	case"(J)V":fmt.Printf("%v\n",stack.PopLong())
	case"(D)V":fmt.Printf("%v\n",stack.PopDouble())
	default:
		panic("println:"+descriptor)}
		stack.PopRef()
}


