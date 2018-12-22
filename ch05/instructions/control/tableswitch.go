package control

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtdata"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low   		  int32
	high	  	  int32
	jumpOffsets	  []int32 //index table
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	//make sure default offset address be the 4x, leave 1-3 padding
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtdata.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= low && index  <= hight {
		offset = self.jumpOffsets[index - low]
	}else {
		offset = self.defaultOffset
	}
	base.Branch(frame, offset)
}

