package loads

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//数组取元素指令，针对不同类型元素，一共定义8条指令; 该指令需要2个操作数，即元素索引和数组引用，从运行时操作数栈弹出

type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}


type BALOAD struct {
	base.NoOperandsInstruction
}

func (self *BALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)

	//todo
	stack.PushInt(int32(refs[index]))
}

type CALOAD struct {
	base.NoOperandsInstruction
}


func (self *CALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Chars()
	checkIndex(len(refs), index)

	//todo check
	stack.PushInt(int32(refs[index]))
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (self *DALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Doubles()
	checkIndex(len(refs), index)
	stack.PushDouble(refs[index])
}

type FALOAD struct {
	base.NoOperandsInstruction
}


func (self *FALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Floats32()
	checkIndex(len(refs), index)
	stack.PushFloat(refs[index])
}


type IALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Ints()
	checkIndex(len(refs), index)
	stack.PushInt(refs[index])
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (self *LALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Longs()
	checkIndex(len(refs), index)
	stack.PushLong(refs[index])
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *SALOAD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	//返回真正持有的golang引用数组
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)

	//TODO
	stack.PushInt(int32(refs[index]))
}


func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLength int, index int32) {
	if index < 0 || index  > int32(arrLength) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
