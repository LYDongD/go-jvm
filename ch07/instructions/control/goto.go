package control

import "gojvm/ch07/instructions/base"
import "gojvm/ch07/rtdata"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
