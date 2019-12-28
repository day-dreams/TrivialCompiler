package interpreter

import (
	"bytes"
	"fmt"
	"github.com/day-dreams/TrivialCompiler/ast"
	"github.com/day-dreams/TrivialCompiler/helper"
	"github.com/day-dreams/TrivialCompiler/io"
	"github.com/day-dreams/TrivialCompiler/lexer"
	parser2 "github.com/day-dreams/TrivialCompiler/parser"
	"io/ioutil"
	"strconv"
	"strings"
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
	pb := doCodeGenGoStructCrudProtoBuffer(cmd)

	if err := ioutil.WriteFile(strings.ToLower(cmd.Param.StructName)+".go", crud, 0666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(strings.ToLower(cmd.Param.StructName)+".proto", pb, 0666); err != nil {
		panic(err)
	}
}

func doCodeGenGoStructCrudProtoBuffer(cmd *ast.Command) []byte {
	name := cmd.Param.StructName
	fields := bytes.NewBuffer(nil)
	for _, field := range cmd.Param.Fields {
		fields.WriteString(fmt.Sprintf("\t%s %s [ json_name = \"%s\" ];\n",
			helper.GoType2pbType(field.GoType), field.Ident, helper.Underlined(field.Ident)))
	}

	// header
	header := bytes.NewBuffer(nil)
	header.WriteString(fmt.Sprintf(`
syntax = "proto3";
package %s;
option go_package = "pb%s";
`, strings.ToLower(name), strings.ToLower(name)))

	// service
	service := bytes.NewBuffer(nil)
	service.WriteString(fmt.Sprintf(`
service %sService {

	rpc List%s (ReqList%s) returns(ResList%s);
	rpc Create%s (ReqCreate%s) returns(ResCreate%s);
	rpc Update%s (ReqUpdate%s) returns(ResUpate%s);
	rpc Delete%s (ReqDelete%s) returns(ResDelete%s);

}
`, name, name, name, name, name, name, name, name, name, name, name, name, name))

	// object
	object := bytes.NewBuffer(nil)
	object.WriteString(fmt.Sprintf(`message %s {\n`, name))
	for _, field := range cmd.Param.Fields {
		object.WriteString(
			fmt.Sprintf(`\t%s %s [ json_name = "%s" ];\n`,
				helper.GoType2pbType(field.GoType),
				field.Ident,
				helper.Underlined(field.Ident)))
	}
	object.WriteString(`}\n`)

	// list
	list := bytes.NewBuffer(nil)
	list.WriteString(fmt.Sprintf(`

message ReqList%s {
%s
}

message ResList%s {
	repeated %s data [ json_tag = "data" ];
}

`, name, string(fields.Bytes()), name, name))

	// create
	create := bytes.NewBuffer(nil)
	create.WriteString(fmt.Sprintf(`

message ReqCreate%s {
%s
}

message ResCreate%s {

}

`, name, string(fields.Bytes()), name))

	// update
	update := bytes.NewBuffer(nil)
	update.WriteString(fmt.Sprintf(`

message ReqUpdate%s {
%s
}

message ResUpdate%s {

}

`, name, string(fields.Bytes()), name))

	// delete
	del := bytes.NewBuffer(nil)
	del.WriteString(fmt.Sprintf(`

message ReqDelete%s {
%s
}

message ResDelete%s {

}

`, name, string(fields.Bytes()), name))

	return helper.Contact(header.Bytes(), service.Bytes(), list.Bytes(), create.Bytes(), update.Bytes(), del.Bytes())
}

func doCodeGenGoStructCrudParam(cmd *ast.Command) []byte {
	name := cmd.Param.StructName
	fields := bytes.NewBuffer(nil)
	for _, field := range cmd.Param.Fields {
		fields.WriteString(fmt.Sprintf("\t%s %s `validate:\" \"`\n", field.Ident, field.GoType))
	}

	crud := bytes.NewBuffer(nil)
	crud.WriteString(fmt.Sprintf("package %s\n\n", strings.ToLower(name)))
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
