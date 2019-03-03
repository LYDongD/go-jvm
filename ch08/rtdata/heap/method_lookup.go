package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	//从当前类开始不断向上查找，找到为止
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}


func lookupInInterfaces(ifaces []*Class, name, desciptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == desciptor {
				return method
			}
		}

		//继续向上查找
		method := lookupInInterfaces(iface.interfaces, name, desciptor)

		if method != nil {
			return method
		}
	}

	return nil
}