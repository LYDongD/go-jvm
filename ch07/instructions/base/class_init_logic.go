package base

import (
	"gojvm/ch07/rtdata"
	"gojvm/ch07/rtdata/heap"
)

func InitClass(thread *rtdata.Thread, class *heap.Class) {
	//调整初始状态
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

//初始化类
func scheduleClinit(thread *rtdata.Thread, class *heap.Class) {

	//获取init方法
	clinit := class.GetClInitMethod()
	if clinit != nil {
		//execute init

		//压栈
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

//初始化超类
func initSuperClass(thread *rtdata.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		//超类未初始化才进行初始化，递归调用
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
