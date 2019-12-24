package ast

import "github.com/day-dreams/TrivialCompiler/token"

type Attrib interface{}

type World struct {
	Name string
}

func (w *World) String() string {
	return "hello " + w.Name
}

func NewWorld(id Attrib) (*World, error) {
	return &World{Name: string(id.(*token.Token).Lit)}, nil
}
