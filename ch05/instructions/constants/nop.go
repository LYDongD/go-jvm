package constants

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(fram *rtdata.Frame) {
	
}

