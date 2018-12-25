package constants

import (
	"gojvm/ch06/instructions/base"
	"gojvm/ch06/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(fram *rtdata.Frame) {
	
}

