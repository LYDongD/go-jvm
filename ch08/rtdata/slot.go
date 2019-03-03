package rtdata

import "gojvm/ch08/rtdata/heap"

type Slot struct {
	num int32
	ref *heap.Object
}