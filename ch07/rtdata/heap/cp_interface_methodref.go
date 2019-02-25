package heap

import "gojvm/ch07/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)

	return ref
}

//解析方法的符号引用, 如果已经解析，则直接返回解析好的目标方法引用
func (self *InterfaceMethodRef) ResolveInterfaceMethod() *Method {
	if self.method == nil {
		return self.resolveInterfaceMethodRef()
	}

	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethodRef() *Method{
	//获取主类和被调用方法的类
	d := self.cp.class
	c := self.ResolvedClass()

	//接口方法的类必须是接口
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//从方法的调用类中，根据符号引用的名称和方法描述符查找目标方法引用
	method := lookupInterfaceMethod(c, self.name, self.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	//判断主类能否访问该接口方法, 主要看方法是否为pubLic
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAceesError")
	}

	return method
}

//在调用类中查找接口方法
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	//迭代法, 找到名称和描述符完全匹配的方法
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	//如果当前调用接口类找不到，则递归地向上查找
	return lookupInInterfaces(iface.interfaces, name, descriptor)
}

