package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/instructions/loads"
	"go-jvm/ch05/instructions/math"
	"go-jvm/ch05/instructions/stores"
	"go-jvm/ch05/rtdata"
)

//contains an instruction, which is an interface type
type WIDE struct {
	modificationInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		//extends index width to 2 bytes
		inst.Index = reader.ReadUint16()
		self.modificationInstruction = inst
	case 0x84:
		inst := &math.IINC()
		inst.Index = reader.ReadUint16()
		inst.Const = reader.ReadUint16()
		self.modificationInstruction = inst
	case 0Xa9:
		panic("Unsupported opcode 0xa9!")
	}
}

func (self *WIDE) Execute(frame *rtdata.Frame) {
	self.modificationInstruction.Execute(frame)
}
