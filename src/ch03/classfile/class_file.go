package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16 //新生代？
	majorVersion uint16
	constantPool ConstantPool //常量池
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields [] *MemberInfo
	methods [] *MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error){
	defer func(){
		if r := recover(); r != nil{
			var ok bool;
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	//return self.magic
}

func (sele *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (self *ClassFile) MinorVersion() uint16{
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Filed() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

//从常量池中查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

//从常量池中查找超类名
func (self *ClassFile) superClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""  //java.lang.Object 没有超类
}

//从常量池中查找接口名
func (self *ClassFile) interfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}


