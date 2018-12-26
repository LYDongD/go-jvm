package heap

import "gojvm/ch06/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
}

func newMethod(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}

	return methods
}


func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAtrr := cfMethod.CodeAttribute(); codeAtrr != nil {
		self.maxLocals = codeAtrr.MaxLocals()
		self.maxLocals = codeAtrr.MaxStack()
		self.code = codeAtrr.Code()
	}
}