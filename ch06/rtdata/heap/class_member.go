package heap

import (
	"gojvm/ch06/classfile"
	"strings"
)

/*
 * 	Field和Method成员的父类
 *	定义通用属性，例如访问权限，名称和描述符
 */
type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}


func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}

	c := self.class
	if self.IsProtected() {
		return d == c || d.isSuperClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}

	if !self.isPrivate() {
		return c.getPackageName() == d.getPackageName()
	}

	return d == c
}


