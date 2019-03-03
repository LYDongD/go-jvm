package stores

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//按索引写入数组元素, 该指令需要三个操作数：需要写入的元素值，索引和数组引用

type AASTORE struct {
	 base.NoOperandsInstruction
}

func (self *AASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (self *BASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := int8(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)
	refs[index] = val
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (self *CASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := uint16(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Chars()
	checkIndex(len(refs), index)
	refs[index] = val
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (self *DASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Doubles()
	checkIndex(len(refs), index)
	refs[index] = val
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (self *LASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Longs()
	checkIndex(len(refs), index)
	refs[index] = val
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (self *FASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Floats32()
	checkIndex(len(refs), index)
	refs[index] = val
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *SASTORE) Execute(frame *rtdata.Frame) {
	//获取操作数
	stack := frame.OperandStack()
	val := int16(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	//获取底层golang实现的数组
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)
	refs[index] = val
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
