package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(fram *rtdata.Frame) {
	
}

