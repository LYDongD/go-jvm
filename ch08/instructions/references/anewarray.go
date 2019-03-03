package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//引用类型数组创建指令，它需要2个操作数，一个是类符号引用的常量池索引，从常量池获取；一个是数组长度，从操作数栈弹出


type ANEW_ARRAY struct {
	base.Index16Instruction //需要一个类的符号引用，去常量池加载
}

func (self *ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	//去常量池拿类的符号引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)

	//解析数组的引用类
	componentClass := classRef.ResolvedClass()

	//从操作数栈中获取数组长度
	count := frame.OperandStack().PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	//构造引用类型数组实例并压入操作数栈
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	frame.OperandStack().PushRef(arr)
}
