package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
	"math"
)


type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	v := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(v)
}

type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	v := math.Mod(v1, v2)
	stack.PushDouble(v)
}


type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithematicException: / by zero")
	}
	v := v1 % v2
	stack.PushInt(v)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithematicException: / by zero")
	}
	v := v1 % v2
	stack.PushLong(v)
}

