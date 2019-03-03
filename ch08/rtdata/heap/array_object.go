package heap

//返回特定类型的数组数据，支持byte, short, int, long, double, float， ref 类型数组
//该数组是通过Object的data属性拿到，对于普通对象，data是实例变量数组；对于数组对象，data指向特定类型的数组数据

//支持byte,bool 数组
func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}

//short array
func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}

func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}

func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}

func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}

func (self *Object) Floats32() []float32 {
	return self.data.([]float32)
}

func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}

func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}


//获取数组长度, 根据数组元素类型返回对应的长度，意味着数组对象只能存储同一种类型
func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8: return int32(len(self.data.([]int8)))
	case []int16: return int32(len(self.data.([]int16)))
	case []int32: return int32(len(self.data.([]int32)))
	case []int64: return int32(len(self.data.([]int64)))
	case []uint16: return int32(len(self.data.([]uint16)))
	case []float32: return int32(len(self.data.([]float32)))
	case []float64: return int32(len(self.data.([]float64)))
	case []*Object: return int32(len(self.data.([]*Object)))
	default:
		panic("not array")
	}
}

