package main

import (
	"fmt"
	"go-jvm/ch03/classpath"
	"go-jvm/ch03/classfile"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		StartJVM(cmd)
	}
}

func StartJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *ClassPath) *classfile.ClassFile {
	classData, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Println("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Println("constant count: %v\n", len(cf.ConstantPool()))
	fmt.Println("access flags: 0x%x\n", cd.AccessFlags())
	fmt.Println("this class: %v\n", cf.ClassName())
	fmt.Println("super class: %v\n", cf.SuperClassName())
	fmt.Println("interfaces: %v\n", cf.InterfacesNames())
	fmt.Println("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields {
		fmt.Println("  %s\n", field.Name)
	}
	fmt.Println("methods count: %v\n", len(cf.methods))

}
