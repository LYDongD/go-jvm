package heap


type SymRef struct {
	cp *ConstantPool
	className string //完全限定类名
	class *Class
}

//类符号引用的解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}

	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if c.isAccessibleTo(d) {
		self.class = c
	}
}