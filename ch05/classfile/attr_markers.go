package classfile

type DeprecatedAttribute struct {
	MarkerAtrribute
}

type SyntheticAttribute struct {
	MarkerAtrribute
}

type MarkerAtrribute struct {

}

func (self *MarkerAtrribute) readInfo(reader *ClassReader) {
	
}
