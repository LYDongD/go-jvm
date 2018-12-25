package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtdata.Frame) {
	v := frame.OperandStack().PopInt()
	if v == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}