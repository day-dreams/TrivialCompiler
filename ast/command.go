package ast

import (
	"fmt"
	"github.com/day-dreams/TrivialCompiler/io"
	"github.com/day-dreams/TrivialCompiler/token"
	"reflect"
)

type Command struct {
	Cmd   Cmd
	Param Param
}

func (c Command) TokenLit() string {
	return fmt.Sprintf("cmd:%s\nparam:%s\n", c.Cmd.TokenLit(), c.Param.TokenLit())
}

type Cmd struct {
	Cmd string
}

func (c Cmd) TokenLit() string {
	return c.Cmd
}

type Param struct {
	StructName string          `json:"struct_name"`
	Fields     []GoStructField `json:"fields"`
}

func (p Param) TokenLit() string {
	s := ""
	for _, field := range p.Fields {
		s += fmt.Sprintf("\t%s\n", field.TokenLit())
	}
	return s[0 : len(s)-1]
}

type GoStructField struct {
	Ident  string
	GoType string
	GoTag  string
}

func (g GoStructField) TokenLit() string {
	return fmt.Sprintf("%s\t%s", g.Ident, g.GoType)
}

func NewGoStructField(ident, gotype, gotag Attrib) (*GoStructField, error) {
	i := string(ident.(*token.Token).Lit)
	gt := string(gotype.(*token.Token).Lit)
	tag := string(gotag.(*token.Token).Lit)
	io.Debug("new gostructfield...")
	return &GoStructField{
		Ident:  i,
		GoType: gt,
		GoTag:  tag,
	}, nil
}

func NewGoStructFieldList() ([]GoStructField, error) {
	io.Debug("new gostructfieldlist")
	return []GoStructField{}, nil
}

func AppendGoStructField(fields, field Attrib) ([]GoStructField, error) {
	io.Debug("appending gostructfield")
	_fields, ok := fields.([]GoStructField)
	if !ok {
		return nil, Error("fields is not []GoStructField")
	}
	_field, ok := field.(*GoStructField)
	if !ok {
		return nil, Error("field is not *GoStructField")
	}
	return append(_fields, *_field), nil
}

func NewCommand(cmd Attrib, param Attrib) (*Command, error) {
	io.Debug("new command...")
	return &Command{
		Cmd:   *cmd.(*Cmd),
		Param: *param.(*Param),
	}, nil
}

func NewCmd(cmd Attrib) (*Cmd, error) {
	io.Debug("new cmd... type%v", reflect.TypeOf(cmd))
	return &Cmd{Cmd: string(cmd.(*token.Token).Lit)}, nil
}

func NewParam(ident Attrib, fields Attrib) (*Param, error) {
	io.Debug("new param...")
	name := string(ident.(*token.Token).Lit)
	_fields, ok := fields.([]GoStructField)
	if !ok {
		return nil, Error("fields is not []GoStructField")
	}
	return &Param{
		StructName: name,
		Fields:     _fields,
	}, nil
}
