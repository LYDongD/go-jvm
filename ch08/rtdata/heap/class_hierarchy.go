package heap


func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self  //if t = s lawful

	if s == t {
		return true
	}

	//分成2种情况，s是数组类或非数组类
	if !s.IsArray() {
		//进一步判断S是否为接口
		if !s.IsInterface() {
			//s继承t或s实现t => t可以被s赋值
			if !t.IsInterface() {
				return s.isSubClassOf(t)
			} else {
				return s.isImplements(t)
			}
		}else {
			//如果t不是接口，则必须是接口的父类Object
			if !t.IsInterface() {
				return t.isJlObject()
			}else {
				return t.isSuperInterfaceOf(s)
			}
		}
	}else {
		//进一步判断t是否为数组
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			}else { //如果t是接口，则必须是数组的接口，即cloneable和serializable
				return t.isJlCloneable() || t.isJioSerializable()
			}
		}else {
			//如果t,s 都是数组，可以考察数组元素类型，相同或可以被赋值
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
}

// self extends c
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}
