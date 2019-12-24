package test

import (
	"fmt"
	"github.com/day-dreams/TrivialCompiler/ast"
	"github.com/day-dreams/TrivialCompiler/helper"
	"github.com/day-dreams/TrivialCompiler/lexer"
	"github.com/day-dreams/TrivialCompiler/parser"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {

	const input = `
			(1+2)  *	3/
	4-5;
	`
	output := &ast.Program{Stats: []ast.Statement{
		ast.InfixExpression{
			Left: ast.InfixExpression{
				Left: ast.InfixExpression{
					Left: ast.InfixExpression{
						Left:     ast.IntegerLiteral{Value: "1"},
						Right:    ast.IntegerLiteral{Value: "2"},
						Operator: "+",
					},
					Right:    ast.IntegerLiteral{Value: "3"},
					Operator: "*",
				},
				Right:    ast.IntegerLiteral{Value: "4"},
				Operator: "/",
			},
			Right:    ast.IntegerLiteral{Value: "5"},
			Operator: "-",
		},
	}}

	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	node, err := p.Parse(l)
	if err != nil {
		t.Fatal(err)
	}
	program, ok := node.(*ast.Program)
	if !ok {
		t.Fatal("TestParser failed.")
	}

	if !reflect.DeepEqual(helper.ToPrettyJson(program), helper.ToPrettyJson(output)) {
		t.Fatal("input != output.")
	}
	fmt.Printf("%s\n", helper.ToPrettyJson(program))
}
