package io

import "fmt"

func Writeln(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}
func Write(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func Debug(format string, v ...interface{}) {
	//Writeln(format, v...)
}
