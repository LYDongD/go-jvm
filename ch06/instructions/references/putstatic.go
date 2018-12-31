package references

import (
	"gojvm/ch06/instructions/base"
	"gojvm/ch06/rtdata"
	"gojvm/ch06/rtdata/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

//给静态成员变量赋值
func (self *PUT_STATIC) Execute(frame *rtdata.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//static final 静态常量只能在clinit方法中赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	//描述符包含了static field的类型信息
	descriptor := field.Descriptor()
	//给成员变量赋值 =》从类的静态成员变量表中给指定成员变量id(索引)填值
	slotId := field.SlotId()
	slots := class.StaticVars()
	//static field 的值从stack弹出
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z','B','C','S','I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L':
		slots.SetRef(slotId, stack.PopRef())
	}

}
