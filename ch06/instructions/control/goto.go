package control

import "gojvm/ch06/instructions/base"
import "gojvm/ch06/rtdata"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
