package main

import (
	"bufio"
	"flag"
	interpreter2 "github.com/day-dreams/TrivialCompiler/interpreter"
	"github.com/day-dreams/TrivialCompiler/io"
	"io/ioutil"
	"os"
)

var (
	arg1 *string
)

func init() {
	arg1 = flag.String("src", "", "gen code from src")
	flag.Parse()
}
func cmd() {

	bytes, err := ioutil.ReadFile(*arg1)
	if err != nil {
		io.Writeln("fail to open file: %v", err)
	}
	interpreter := interpreter2.Interpreter{}
	interpreter.Interpret(string(bytes))
}

func interpreter() {
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
func main() {

	if *arg1 == "" {
		interpreter()
	} else {
		cmd()
	}
}
