package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs		  int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	reader.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32() // case-value paire
	self.matchOffsets = reader.ReaderInt32s(self.npairs * 2) //use 2 units to save a paire
}

func (self *LOOKUP_SWITCH) Execute(frame *rtdata.Frame) {
	key := frame.OperandStack().PopInt()
	for i := 0; i < self.npairs * 2; i += 2 {
		if key == self.matchOffsets[i] {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, offset)
			return
		}
	}
} 