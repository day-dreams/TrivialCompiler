package main

import (
	"bufio"
	interpreter2 "github.com/day-dreams/TrivialCompiler/interpreter"
	"github.com/day-dreams/TrivialCompiler/io"
	"os"
)

func main() {
	interpreter := interpreter2.Interpreter{}
	reader := bufio.NewReader(os.Stdin)
	for {
		io.Write("$ ")
		x, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		interpreter.Interpret(x)
	}
}
