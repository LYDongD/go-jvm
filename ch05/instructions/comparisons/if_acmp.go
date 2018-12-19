package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type IF_ACMEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMEQ) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()

	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMNE struct {
	base.BranchInstruction
}

func (self *IF_ACMNE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()

	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}