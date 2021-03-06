package main

import (
	"fmt"
	"strings"
	"gojvm/ch06/classpath"
	"gojvm/ch06/rtdata/heap"
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
	className := strings.Replace(cmd.class, ".", "/", -1)
	classLoader := heap.NewClassLoader(cp)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	}else {
		fmt.Printf("main method not found in class :%s\n", className)
	}
}




