package interpreter

import (
	"fmt"
	"github.com/day-dreams/TrivialCompiler/ast"
	"github.com/day-dreams/TrivialCompiler/io"
	"github.com/day-dreams/TrivialCompiler/lexer"
	parser2 "github.com/day-dreams/TrivialCompiler/parser"
	"strconv"
)

type Interpreter struct {
}

func (i *Interpreter) Interpret(source string) {

	lex := lexer.NewLexer([]byte(source))
	parser := parser2.NewParser()
	node, err := parser.Parse(lex)
	if err != nil {
		io.Writeln("parse failed.")
		return
	}

	program, ok := node.(*ast.Program)
	if !ok {
		io.Writeln("parse failed.")
		return
	}

	for _, stat := range program.Stats {
		value := getValue(stat)
		io.Writeln("%s = %s", stat.TokenLit(), value)
	}

}

func doOperation(left, right, op string) string {
	a, err := strconv.ParseFloat(left, 64)
	if err != nil {
		return "invalid"
	}
	b, err := strconv.ParseFloat(right, 64)
	if err != nil {
		return "invalid"
	}
	switch op {
	case "+":
		return fmt.Sprintf("%.8f", a+b)
	case "-":
		return fmt.Sprintf("%.8f", a-b)
	case "*":
		return fmt.Sprintf("%.8f", a*b)
	case "/":
		return fmt.Sprintf("%.8f", a/b)
	default:
		return "invalid"
	}

}

func getValue(node ast.Node) string {
	switch node.(type) {
	case *ast.IntegerLiteral:
		return node.TokenLit()
	case *ast.InfixExpression:
		left := getValue(node.(*ast.InfixExpression).Left)
		right := getValue(node.(*ast.InfixExpression).Right)
		return doOperation(left, right, node.(*ast.InfixExpression).Operator)
	default:
		return "invalid node"
	}
}
