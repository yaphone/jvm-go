package classfile

type DeprecatedAttrbute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {

}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}
