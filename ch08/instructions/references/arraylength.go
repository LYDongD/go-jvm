package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
)

//获取数组长度指令，只需要一个操作数：数组引用本身，运行时从操作数栈弹出

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}