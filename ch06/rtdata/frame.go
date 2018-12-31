package rtdata

import (
	"gojvm/ch06/rtdata/heap"
)

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPc       int
}

/*
*   解析class文件时已经可以获取以下信息(编译器预先计算好))，保存在方法区的methodInfo结构中
*	maxLocals 最大局部变量数
*	maxStack 操作数栈最大深度
 */
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		method: method,
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPc = pc
}

func (self *Frame) NextPC() int {
	return self.nextPc
}

func (self *Frame) Method() *heap.Method {
	return self.method
}
