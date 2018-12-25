package instructions

import (
	"fmt"
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/instructions/comparisions"
	"go-jvm/ch05/instructions/constants"
	"go-jvm/ch05/instructions/control"
	"go-jvm/ch05/instructions/conversions"
	"go-jvm/ch05/instructions/extended"
	"go-jvm/ch05/instructions/loads"
	"go-jvm/ch05/instructions/math"
	"go-jvm/ch05/instructions/stack"
	"go-jvm/ch05/instructions/stores"
)

var (
	nop         = &NOP{}
	aconst_null = &ACONST_NULL{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	default:
		panic(fmt.Errorf("Unsupported opcaode: %x", opcode))
	}

}
