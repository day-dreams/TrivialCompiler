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

func TestParserStatement(t *testing.T) {

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
		fmt.Printf("::%s\n", helper.ToPrettyJson(program))
		fmt.Printf("::%s\n", helper.ToPrettyJson(output))
		t.Fatal("input != output.")
	}
}

func TestParserCodeGenGoStruct(t *testing.T) {

	//const input = "CodeGenGoStruct type User struct {Id int `gorm:\"id\"`}"
	const input = "CodeGenGoStruct type User struct {" +
		"Id int `gorm:\"id\" json:\"-\"` " +
		"Name string `gorm:\"username\"` " +
		"Int64 int64 `gorm:\"int64\"` " +
		"Float64 float64 `gorm:\"float64\"` " +
		"Bool bool `gorm:\"bool\"` " +
		"NoTag string" +
		"}"
	output := &ast.Program{
		Stats: nil,
		Command: ast.Command{
			Cmd: ast.Cmd{Cmd: "CodeGenGoStruct"},
			Param: ast.Param{
				StructName: "User",
				Fields: []ast.GoStructField{
					{Ident: "Id", GoType: "int", GoTag: ast.GoTag{Tag: "`gorm:\"id\" json:\"-\"`"}},
					{Ident: "Name", GoType: "string", GoTag: ast.GoTag{"`gorm:\"username\"`"}},
					{Ident: "Int64", GoType: "int64", GoTag: ast.GoTag{"`gorm:\"int64\"`"}},
					{Ident: "Float64", GoType: "float64", GoTag: ast.GoTag{"`gorm:\"float64\"`"}},
					{Ident: "Bool", GoType: "bool", GoTag: ast.GoTag{"`gorm:\"bool\"`"}},
					{Ident: "NoTag", GoType: "string"},
				},
			},
		},
	}

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
		fmt.Printf("::%s\n", helper.ToPrettyJson(program))
		fmt.Printf("::%s\n", helper.ToPrettyJson(output))
		t.Fatal("input != output.")
	}
}
