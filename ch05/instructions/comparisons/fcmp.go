package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type FCOMPG struct {
	base.NoOperandsInstruction
}


type FCOMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtdata.Frame gFlag bool) {
	stack := frame.OperandStack()
	v
}