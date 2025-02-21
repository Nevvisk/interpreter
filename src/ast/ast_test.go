package ast

import (
	"interpreter/src/token"
	"testing"
)

func TestString(t *testing.T) {
	TestLetStatement(t)
}

func TestLetStatement(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anoterVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar" {
		t.Errorf("program.String() wrong. got = %q", program.String())
	}
}
