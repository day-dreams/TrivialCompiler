package test

import (
	"github.com/day-dreams/TrivialCompiler/interpreter"
	"testing"
)

func TestInterpreter(t *testing.T) {
	const input = interpreter.CodeGenGoStruct + " " +
		"type User struct {" +
		"	Id int64 `json:\"id\" gorm:\"id\"`" +
		"	Name string `json:\"name\" gorm:\"name\"`" +
		"	PhoneNumber string `json:\"phone\" gorm:\"phone\"`" +
		"	PhoneNumberAtHome string `json:\"phone_at_home\" gorm:\"phone\"`" +
		"	Graduated bool `json:\"graduated\" gorm:\"graduated\"`" +
		"	Weight float64 `json:\"weight\" gorm:\"weight\"`" +
		"}"

	intp := interpreter.Interpreter{}
	intp.Interpret(input)
}
