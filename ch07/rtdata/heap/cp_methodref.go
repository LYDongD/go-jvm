package heap

import "gojvm/ch07/classfile"

type MethodRef struct {
	MemberRef
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

func (self *MethodRef) resolveMethodRef() *Method {

	//解析方法前先解析方法的调用类，解析成功前self.class为空，需要通过当前类加载器加载该方法的调用类
	d := self.cp.class
	c := self.ResolvedClass()

	//这里解析的是非接口方法
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//根据方法名和描述符查找
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {

	//在类继承体系中寻找
	method := LookupMethodInClass(class, name, descriptor)

	if method == nil {
		//在类接口中寻找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}

	return method
}