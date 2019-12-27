package test

import (
	"github.com/day-dreams/TrivialCompiler/lexer"
	"github.com/day-dreams/TrivialCompiler/token"
	"testing"
)

func TestLexer(t *testing.T) {
	type Case struct {
		expectedType    token.Type
		expectedLiteral string
	}
	var input = `
		(1+2)  *	3/
4-5;
	"

CodeGenGoStruct
type
struct
int
int64
bool
float64
`
	input += "`gorm:\"tag name\"`"

	tokMap := token.TokMap
	cases := []Case{
		{expectedType: tokMap.Type("lparen"), expectedLiteral: "("},
		{expectedType: tokMap.Type("int"), expectedLiteral: "1"},
		{expectedType: tokMap.Type("plus"), expectedLiteral: "+"},
		{expectedType: tokMap.Type("int"), expectedLiteral: "2"},
		{expectedType: tokMap.Type("rparen"), expectedLiteral: ")"},
		{expectedType: tokMap.Type("mul"), expectedLiteral: "*"},
		{expectedType: tokMap.Type("int"), expectedLiteral: "3"},
		{expectedType: tokMap.Type("div"), expectedLiteral: "/"},
		{expectedType: tokMap.Type("int"), expectedLiteral: "4"},
		{expectedType: tokMap.Type("minus"), expectedLiteral: "-"},
		{expectedType: tokMap.Type("int"), expectedLiteral: "5"},
		{expectedType: tokMap.Type("semicolon"), expectedLiteral: ";"},
		{expectedType: tokMap.Type("dquote"), expectedLiteral: `"`},

		// go code gen
		{expectedType: tokMap.Type("cmdcodegengostruct"), expectedLiteral: "CodeGenGoStruct"},
		{expectedType: tokMap.Type("gotypeof"), expectedLiteral: "type"},
		{expectedType: tokMap.Type("gostructdef"), expectedLiteral: "struct"},
		{expectedType: tokMap.Type("goint"), expectedLiteral: "int"},
		{expectedType: tokMap.Type("goint64"), expectedLiteral: "int64"},
		{expectedType: tokMap.Type("gobool"), expectedLiteral: "bool"},
		{expectedType: tokMap.Type("gofloat64"), expectedLiteral: "float64"},
		{expectedType: tokMap.Type("gotag"), expectedLiteral: "`gorm:\"tag name\"`"},
	}

	l := lexer.NewLexer([]byte(input))
	for index, tcase := range cases {
		tok := l.Scan()
		if tok.Type != tcase.expectedType || string(tok.Lit) != tcase.expectedLiteral {
			t.Fatalf("case %d failed. have {%v,%v}, expected {%v,%v}",
				index, tok.Type, string(tok.Lit), tcase.expectedType, tcase.expectedLiteral)
		} else {
			t.Logf("case %d passed. have {%v,%v}, expected {%v,%v}",
				index, tok.Type, string(tok.Lit), tcase.expectedType, tcase.expectedLiteral)
		}
	}

}
