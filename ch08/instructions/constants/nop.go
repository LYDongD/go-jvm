package constants

import (
	"gojvm/ch08/instructions/base"
	"gojvm/ch08/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(fram *rtdata.Frame) {
	
}

