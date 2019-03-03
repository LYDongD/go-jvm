package heap

//基本类型的描述符表
var primitiveTypes = map[string]string{
	"void" : "V",
	"boolean" : "Z",
	"byte" : "B",
	"short" : "S",
	"int" : "I",
	"long" : "L",
	"float" : "F",
	"double" : "D",
}

//获取类名的描述符形式
func toDescriptor(className string) string {

	//如果数组类名，描述符即本身
	if className[0] == '[' {
		return className
	}

	//如果是基本类型，则查表返回
	if d, ok := primitiveTypes[className]; ok {
		return d
	}

	return "L" + className + ";"
}

//获取描述符对应的类名
func toClassName(descriptor string) string {
	//数组的描述符和类名是一致的
	if descriptor[0] == '[' {
		return descriptor
	}

	//引用类型描述符，去掉开头的L和末尾的分号
	if descriptor[0] == 'L' {
		return descriptor[1:len(descriptor) - 1]
	}

	//基本类型描述符
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}

	panic("invalid descriptor: " + descriptor)
}
