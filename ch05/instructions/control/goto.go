package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}