package control

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
)

type RETURN struct {
	base.NoOperandsInstruction
}

type IRETURN struct {
	base.NoOperandsInstruction
}
type ARETURN struct {
	base.NoOperandsInstruction
}

type LRETURN struct {
	base.NoOperandsInstruction
}

type DRETURN struct {
	base.NoOperandsInstruction
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtdata.Frame) {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}


func (self *LRETURN) Execute(frame rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}


func (self *ARETURN) Execute(frame rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *DRETURN) Execute(frame rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

func (self *FRETURN) Execute(frame rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}