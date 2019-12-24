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
	const input = `
		(1+2)  *	3/
4-5;
`
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
