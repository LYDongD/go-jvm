package constants

import (
	"gojvm/ch07/instructions/base"
	"gojvm/ch07/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(fram *rtdata.Frame) {
	
}

