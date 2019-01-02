package main

import (
	"fmt"
	"gojvm/ch07/instructions"
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
)


func interpret(method *heap.Method) {

	//thread -> frame -> execute bytecode in frame
	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}

//handle panic -> recover
func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtdata.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()

		//bycode(opcode) -> instructions
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc:%2d instruction:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
