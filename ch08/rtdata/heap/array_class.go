package heap

//数组类不是类加载器加载创建的，而是在运行时生成
//创建数组，根据数组类名"[类型描述符"模式进行创建
func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("Not array class: " + self.name)
	}

	switch self.name {
	case "[Z": return &Object{self, make([]int8, count)}
	case "[B": return &Object{self, make([]int8, count)}
	case "[C": return &Object{self, make([]uint16, count)}
	case "[S": return &Object{self, make([]int16, count)}
	case "[I": return &Object{self, make([]int32, count)}
	case "[L": return &Object{self, make([]int64, count)}
	case "[F": return &Object{self, make([]float32, count)}
	case "[D": return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

//数组类的类名是以[开头的
func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

