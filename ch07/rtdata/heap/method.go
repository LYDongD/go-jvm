package heap

import (
	"gojvm/ch07/classfile"
)

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calArgSlotCount()
	}

	return methods
}


func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAtrr := cfMethod.CodeAttribute(); codeAtrr != nil {
		self.maxLocals = codeAtrr.MaxLocals()
		self.maxStack = codeAtrr.MaxStack()
		self.code = codeAtrr.Code()
	}
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Method) calArgSlotCount() uint {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, parameterType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++

		//long或double类型占用2个slot
		if parameterType == "J" || parameterType == "D" {
			self.argSlotCount++
		}
	}

	//实例方法额外增加this参数
	if !self.IsStatic() {
		self.argSlotCount++
	}

	return self.argSlotCount
}

