package heap

import "gojvm/ch07/classfile"

//方法的符号引用
type MethodRef struct {
	MemberRef
	//缓存解析的方法目标引用
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.cp = cp
	methodRef.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return methodRef
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}

	return self.method
}

//方法符号引用解析，动态绑定的最后阶段，解析类找到方法所属的类，然后在该类中查找目标方法引用
func (self *MethodRef) resolveMethodRef() *Method {

	//获取当前类A
	d := self.cp.class

	//获取方法调用类B， A持有B的引用，即在A类中调用B类的方法
	c := self.ResolvedClass()

	//这里解析的是非接口方法
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//根据方法名和描述符在方法调用类中查找
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method

	return method
}



func lookupMethod(class *Class, name, descriptor string) *Method {

	//在类继承体系中寻找
	method := LookupMethodInClass(class, name, descriptor)

	if method == nil {
		//在类接口中寻找
		method = lookupInInterfaces(class.interfaces, name, descriptor)
	}

	return method
}