package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//原始类型数组创建指令,它需要2个操作数，数组类型操作数和数组长度(从运行时操作数栈中获取)

//原始类型数组的类型操作数
const (
	AT_BOOLEAN = 4
	AT_CHAR = 5
	AT_FLOAT = 6
	AT_DOUBLE = 7
	AT_BYTE = 8
	AT_SHORT = 9
	AT_INT = 10
	AT_LONG = 11
)


type NEW_ARRAY struct {
	atype uint8 //数组元素类型操作数
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtdata.Frame) {
	//从操作数栈中获取数组长度count
	count := frame.OperandStack().PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	//当前方法所属类的类加载器，对原始类型数组类进行加载
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)

	//构造数组对象并入操作数栈
	arr := arrClass.NewArray(uint(count))
	frame.OperandStack().PushRef(arr)
}

//根据数组元素类型获取类对象
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN: return loader.LoadClass("[Z")
	case AT_BYTE: return loader.LoadClass("[B")
	case AT_CHAR: return loader.LoadClass("[C")
	case AT_SHORT: return loader.LoadClass("[S")
	case AT_INT: return loader.LoadClass("[I")
	case AT_LONG: return loader.LoadClass("[J")
	case AT_FLOAT: return loader.LoadClass("[F")
	case AT_DOUBLE: return loader.LoadClass("[D")
	default:
		panic("Invalid atype")
	}
}








