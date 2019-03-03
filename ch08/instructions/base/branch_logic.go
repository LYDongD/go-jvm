package base

import "gojvm/ch08/rtdata"

func Branch(frame *rtdata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
