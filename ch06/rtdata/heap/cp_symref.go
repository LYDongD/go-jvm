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

//类符号引用的解析: d类调用c类，d持有c的类符号引用
func (self *SymRef) resolveClassRef() {
	d := self.cp.class //符号引用所在常量池指向的类
	c := d.loader.LoadClass(self.className) //加载符号引用表示的类
	if c.isAccessibleTo(d) { //如果该类可以被访问，例如d可以调用c
		self.class = c  //设置符号引用指向该类
	}
}