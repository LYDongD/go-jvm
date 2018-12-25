package main

import (
	"fmt"
	"gojvm/ch06/classfile"
	"gojvm/ch06/instructions"
	"gojvm/ch06/instructions/base"
	"gojvm/ch06/rtdata"
)


func interpret(methodInfo *classfile.MemberInfo) {
	//get bytecode and operate env
	codeAtrr := methodInfo.CodeAttribute()
	maxLocals := codeAtrr.MaxLocals()
	maxStack := codeAtrr.MaxStack()
	bytecode := codeAtrr.Code()

	//thread -> frame -> execute bytecode in frame
	thread := rtdata.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
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
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
