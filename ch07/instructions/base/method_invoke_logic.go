package base

import (
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
	"fmt"
)

//方法调用，从当前主类栈帧传递参数到新的栈帧
func InvokeMethod(invokerFrame *rtdata.Frame, method *heap.Method) {
	//创建新的方法栈帧
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)

	//入栈
	thread.PushFrame(newFrame)

	//根据方法的参数表，传递参数, 迭代法
	//参数个数 <= slot个数(long或double占用两个slot, 可能还包含隐性参数等)
	argsSlotCount := int(method.ArgSlotCount())
	if argsSlotCount > 0 {
		//参数从后往前从操作数栈中弹出
		for i := argsSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			//将参数写入新栈帧的局部变量表
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		}else {
			panic(fmt.Sprint("native method: %v.%v%v \n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
