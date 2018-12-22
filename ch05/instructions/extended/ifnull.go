package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
