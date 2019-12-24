package ast

import "github.com/day-dreams/TrivialCompiler/token"

type Attrib interface{}

type Program struct {
	Stats []Statement `json:"stats"`
}

type Node interface {
	TokenLit() string
}

type Statement interface {
	Node
	statNode()
}

type Expression interface {
	Node
	exprNode()
}

type IntegerLiteral struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

func (i IntegerLiteral) TokenLit() string {
	return string(i.Token.Lit)
}

func (i IntegerLiteral) exprNode() {
}

type InfixExpression struct {
	Token    *token.Token `json:"-"`
	Type     string       `json:"-"`
	Left     Expression   `json:"left"`
	Right    Expression   `json:"right"`
	Operator string       `json:"operator"`
}

func (i InfixExpression) statNode() {
}

func (i InfixExpression) TokenLit() string {
	return string(i.Token.Lit)
}

func (i InfixExpression) exprNode() {
}
