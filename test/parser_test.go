package test

import (
	"github.com/day-dreams/TrivialCompiler/ast"
	"github.com/day-dreams/TrivialCompiler/lexer"
	"github.com/day-dreams/TrivialCompiler/parser"
	"testing"
)

func TestWorld(t *testing.T) {
	in := []byte(`hello gocc`)

	lex := lexer.NewLexer(in)
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}

	if w, ok := st.(*ast.World); !ok {
		t.Fatalf("This is not a world")
	} else if w.Name != `gocc` {
		t.Fatalf("Wrong world %v", w.Name)
	}
}
