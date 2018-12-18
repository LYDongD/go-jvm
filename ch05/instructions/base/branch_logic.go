package base

import (
	"go-jvm/ch05/rtdata"
)

func Branch(frame *rtdata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPc := pc + offset
	frame.SetNextPC(nextPC)
}