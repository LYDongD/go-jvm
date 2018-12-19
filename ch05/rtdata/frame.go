package rtdata

type Frame struct {
	lower *Frame
	localVars LocalVars 
	operandStack *OperandStack
}

/* 
*   解析class文件时已经可以获取以下信息(编译器预先计算好))，保存在方法区的methodInfo结构中
*	maxLocals 最大局部变量数
*	maxStack 操作数栈最大深度
*/
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

//todo
func (self *Frame) Thread() *Thread {
	return nil
}






