package ast

import (
	"testing"

	"github.com/MYKatz/PLZ/token"
)

func TestString(t *testing.T) {
	//tests printing 'let myvar be 5 plz'
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myvar"},
					Value: "myvar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.INT, Literal: "5"},
					Value: "5",
				},
			},
		},
	}
	if program.String() != "let myvar be 5 plz" {
		t.Errorf("program.String() incorrect. Expected 'let myvar be 5 plz', received %q", program.String())
	}
}
