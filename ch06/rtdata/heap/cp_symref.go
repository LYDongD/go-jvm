package heap


type SymRef struct {
	cp *ConstantPool
	className string //完全限定类名
	class *Class
}