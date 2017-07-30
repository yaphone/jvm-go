package main

import "fmt"
//import "strings"
import "ch02/classpath"
import (
	"strings"
	"ch03/classfile"
)


func main(){
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	}else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	}else{
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	// fmt.Println("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOpthon, cmd.cpOption)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Printf(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("\n*********************\n")
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile)  {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MajorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fileds()))
	for _, f := range cf.Fileds() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}