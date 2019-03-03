package control

import "gojvm/ch08/instructions/base"
import "gojvm/ch08/rtdata"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
