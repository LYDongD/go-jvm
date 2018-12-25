package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//取低5位(32个数)即可表示可移动的最大位数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

type ISHR struct {
	base.NoOperandsInstruction
}

//算术右移，高位补符号位，采用golang默认的位移运算符
func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

//无符号(逻辑)右移，高位补0
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//取低6位(64个数)即可表示可移动的最大位数
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//取低6位(64个数)即可表示可移动的最大位数
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}


type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
