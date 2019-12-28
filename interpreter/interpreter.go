package interpreter

import (
	"bytes"
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
		io.Writeln("parse failed:")
		io.Writeln("\t%v", err)
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

	doCommand(&program.Command)
}

const (
	CodeGenGoStruct = "CodeGenGoStruct"
)

var (
	cmd2func = map[string]func(command *ast.Command){
		CodeGenGoStruct: doCodeGenGoStruct,
	}
)

func doCommand(cmd *ast.Command) {
	handle := cmd2func[cmd.Cmd.TokenLit()]
	handle(cmd)
}

// 根据struct定义，生成crud函数的param和pb文件
func doCodeGenGoStruct(cmd *ast.Command) {
	crud := doCodeGenGoStructCrudParam(cmd)
	io.Writeln(string(crud))
}

func doCodeGenGoStructCrudParam(cmd *ast.Command) []byte {
	name := cmd.Param.StructName
	fields := bytes.NewBuffer(nil)
	for _, field := range cmd.Param.Fields {
		fields.WriteString(fmt.Sprintf("\t%s %s `validate:\" \"`\n", field.Ident, field.GoType))
	}

	crud := bytes.NewBuffer(nil)
	crud.WriteString(fmt.Sprintf("type ParamCreate%s {\n", name))
	crud.Write(fields.Bytes())
	crud.WriteString(fmt.Sprintf("}\n"))
	crud.WriteString(fmt.Sprintf("type ParamUpdate%s {\n", name))
	crud.Write(fields.Bytes())
	crud.WriteString(fmt.Sprintf("}\n"))
	crud.WriteString(fmt.Sprintf("type ParamList%s {\n", name))
	crud.Write(fields.Bytes())
	crud.WriteString(fmt.Sprintf("}\n"))
	crud.WriteString(fmt.Sprintf("type ParamDelete%s {\n", name))
	crud.Write(fields.Bytes())
	crud.WriteString(fmt.Sprintf("}\n"))

	return crud.Bytes()
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
