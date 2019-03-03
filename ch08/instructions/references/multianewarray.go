package references

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//多维数组创建指令，该指令包含2个操作数，紧跟在指令后，无需从操作数栈中加载
//1 类在常量池中的符号引用 2 数组的维度
//该指令还需要多个操作数作为数组各个维度的长度，需要在运行时确定，即从操作数栈中弹出

type MULTI_ANEW_ARRAY struct {
	index uint16 //类符号引用在常量池中的索引
	dimensions uint8 //维度
}

//从字节码中读取操作数
func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	//从常量池中获取数组类的符号引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)

	//类加载，解析符号引用为具体类对象,这里的类是数组类，不是元素的类
	arrClass := classRef.ResolvedClass()

	//获取各个维度的长度
	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))

	//构建数组对象
	arr := newMultiDimensionArray(counts, arrClass)
	stack.PushRef(arr)
}

//根据维度个数，弹栈获取维度长度
func popAndCheckCounts(stack *rtdata.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >-0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

//递归构建多维数组，先构建引用数组，对于每个引用，递归构建多维数组
func newMultiDimensionArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}










