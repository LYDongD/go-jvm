package heap

import (
	"gojvm/ch08/classfile"
	"strings"
)

type Class struct {
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint //实例变量占用空间
	staticSlotCount uint //类变量占用空间
	staticVars Slots //静态变量表
	initStarted bool //类是否已经初始化
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) isPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}


func (self *Class) isFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *Class) isSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}


func (self *Class) isAccessibleTo(other *Class) bool {
	return self.isPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.Index(self.name, "/"); i > 0 {
		return self.name[:i]
	}
	return ""
}


func (self *Class) NewObject() *Object {
	return NewObject(self)
}

func NewObject(class *Class) *Object {
	return &Object{
		class : class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name string, descriptor string) *Method{
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return nil
}

//获取初始化方法
func (self *Class) GetClInitMethod() *Method {
	return self.getStaticMethod("<cinit>", "()V")
}

//根据数组元素类得到数组类，名字映射关系
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

//根据引用类型数组元素的类名获取数组本身的类名，例如java/lang/String -> [java/lang/string
func getArrayClassName(className string) string{
	return "[" + toDescriptor(className)
}
