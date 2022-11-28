package main

import (
	"awesomeProject/ch04/rtda"
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)

	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetInt(1))
	fmt.Println(vars.GetLong(2))
	fmt.Println(vars.GetLong(4))
	fmt.Println(vars.GetFloat(6))
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)

	fmt.Println(ops.PopRef())
	fmt.Println(ops.PopDouble())
	fmt.Println(ops.PopFloat())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopInt())
	fmt.Println(ops.PopInt())
}
