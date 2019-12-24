// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/day-dreams/TrivialCompiler/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 10,
		Ignore: "",
	},
}