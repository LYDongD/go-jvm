package heap

import "gojvm/ch06/classfile"

type Field struct {
	ClassMember
	SlotId uint
	constantValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAtrributes(cfField)
	}
	
	return fields
}


func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "F"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttribute := cfField.ConstantValueAttribute(); valAttribute != nil {
		self.constantValueIndex = uint(valAttribute.ConstantValueIndex)
	}
}