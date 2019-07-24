package lexer

import (
	"testing"

	"github.com/MYKatz/PLZ/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){}, plz"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.TERMINATOR, "plz"},
	}

	l := NewLexer(input) //eventually will be the lexer
	for i, exp := range tests {
		t := l.NextToken()

		if t.Type != exp.expectedType {
			t.Fatalf("Test #%d: incorrect token type. Expected %s, received %s.", i, exp.expectedType, t.Type)
		}

		if t.Literal != exp.expectedLiteral {
			t.Fatalf("Test #%d: incorrect token literal. Expected %s, received %s", i, exp.expectedLiteral, t.Literal)
		}
	}
}
