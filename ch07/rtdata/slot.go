package rtdata

import "gojvm/ch07/rtdata/heap"

type Slot struct {
	num int32
	ref *heap.Object
}