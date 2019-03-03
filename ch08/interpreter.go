package main

import (
	"fmt"
	"gojvm/ch08/instructions"
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
	"gojvm/ch08/rtdata/heap"
)

//解释器：解释执行
func interpret(method *heap.Method, logInst bool) {

	//thread -> frame -> execute bytecode in frame
	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread) //异常捕获
	//logInst是否打印指令执行日志到console
	loop(thread, logInst)
}

//如果解释器执行过程中抛出异常，则会执行下列方
func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		//打印栈信息
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtdata.Thread, logInst bool) {

	reader := &base.BytecodeReader{}
	for {

		//获取当前指令pc
		frame := thread.PopFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		//读取当前指令字节码
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()

		//创建指令
		inst := instructions.NewInstruction(opcode)
		//获取指令操作数
		inst.FetchOperands(reader)
		//更新PC
		frame.SetNextPC(reader.PC())

		//logger
		if logInst {
			logInstruction(frame, inst)
		}

		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

//打印线程栈信息，包括类,方法和描述符
func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v\n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

//打印执行指令
func logInstruction(frame *rtdata.Frame, ins base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v \n", className, methodName, pc, ins, ins)
}


