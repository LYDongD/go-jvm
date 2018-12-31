//Package heap, 运行时数据区，包括方法区和堆
package heap

import (
	"gojvm/ch06/classpath"
	"fmt"
	"gojvm/ch06/classfile"
)

type ClassLoader struct {
	cp *classpath.Classpath
	classMap map[string]*Class  //cache loaded classes, 可以看成是方法区的具体实现
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp: cp,
		classMap: make(map[string] *Class),
	}
}

//将类数据加载到方法区
func (self *ClassLoader) LoadClass(className string) *Class{
	if class, ok := self.classMap[className]; ok {
		return class
	}

	//非数组类，数组类是虚拟机直接生产，不是直接来自于类文件，另外单独处理
	return self.loadNonArrayClass(className)
}


func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	//查找类文件，文件IO读取数据到内存
	data, entry := self.readClass(className)
	//解析生成类数据，并存入方法区
	class := self.defineClass(data)
	//链接
	link(class)
	fmt.Printf("load %s from %s\n", className, entry)
	return class
}


//用ClassPath加载指定类
//Entry: 类路径项接口
func (self *ClassLoader) readClass(className string) ([]byte, classpath.Entry){
	data, entry, err := self.cp.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + className)
	}

	return data, entry
}

//解析内存字节流为类结构
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}

	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	//虚拟机规范验证
	verify(class)
	//给类变量分配空间
	prepare(class)
}

func verify(class *Class) {}

//给静态成员变量和实例成员变量分配空间
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//计算实例成员变量的个数并进行编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}

	class.instanceSlotCount = slotId
}

//计算静态成员变量的个数并进行编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}

	class.staticSlotCount = slotId
}

//给静态成员变量分配空间并初始化
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.constValueIndex
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z","B","C","S","I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")

		}
	}
}

