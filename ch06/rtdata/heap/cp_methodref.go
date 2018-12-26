package heap

import "gojvm/ch06/classfile"

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
