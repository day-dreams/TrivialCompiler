package ast

import (
	"errors"
	"fmt"
	"github.com/day-dreams/TrivialCompiler/token"
	"reflect"
	"runtime"
)

func Error(msg string) error {
	_, f, line, _ := runtime.Caller(1)
	return errors.New(fmt.Sprintf("in %s, line %d, error:%s", f, line, msg))
}

// source code -> lexer -> parser -> type Program
func NewProgram(stats Attrib) (*Program, error) {
	s, ok := stats.([]Statement)
	if !ok {
		return nil, Error("wrong type of stats")
	}
	return &Program{Stats: s}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
}

func AppendStatement(stats, stat Attrib) ([]Statement, error) {
	_stats, ok := stats.([]Statement)
	if !ok {
		return nil, Error("stats is not []Statement")
	}
	_stat, ok := stat.(Statement)
	if !ok {
		fmt.Printf("type of sta:%v\n", reflect.TypeOf(stat))
		return nil, Error("stat is not Statement")
	}
	return append(_stats, _stat), nil
}

func NewIntegerLiteral(integer Attrib) (Expression, error) {
	if integ, ok := integer.(*token.Token); !ok {
		return nil, Error("integer is not token")
	} else {
		return &IntegerLiteral{
			Token: integ,
			Value: string(integ.Lit),
		}, nil
	}
}

func NewInfixExpression(left, right, oper Attrib) (Expression, error) {
	_left, ok := left.(Expression)
	if !ok {
		return nil, Error("left is not expression")
	}
	_oper, ok := oper.(*token.Token)
	if !ok {
		return nil, Error("oper is not expression")
	}
	_right, ok := right.(Expression)
	if !ok {
		return nil, Error("right is not expression")
	}

	return &InfixExpression{
		Token:    _oper,
		Left:     _left,
		Right:    _right,
		Operator: string(_oper.Lit),
	}, nil
}

func NewExpressionStatement(expr Attrib) (Expression, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, Error("expr is no Expression")
	}
	return e, nil
}
