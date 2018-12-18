package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	
}