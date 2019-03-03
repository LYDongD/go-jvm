package heap

//可以表示普通对象或数组
type Object struct {
	class *Class //类对象的指针
	data interface{} //Slots或各种类型的数组
}


func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) Class() *Class {
	return self.class
}

// self的类是否可以赋值给目标类： 1 self是目标类的子类 2 self 是目标接口的实现类
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}