package lexer

import (
	"testing"

	"github.com/MYKatz/PLZ/token"
)

func TestNextToken(test *testing.T) {
	input := `
	let five be 5 plz
	let ten be 10 plz
	let add be function(a, c) please
		a + c plz
	thanks
	
	let result be add(five, ten) plz
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "be"},
		{token.INT, "5"},
		{token.TERMINATOR, "plz"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "be"},
		{token.INT, "10"},
		{token.TERMINATOR, "plz"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "be"},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "c"},
		{token.RPAREN, ")"},
		{token.LBRACE, "please"}, //please as lbrace and thanks as rbrace
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "c"},
		{token.TERMINATOR, "plz"},
		{token.RBRACE, "thanks"},
		//should maybe have a 'semicolon'/terminator here..
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "be"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.COMMA, "ten"},
		{token.RPAREN, ")"},
		{token.TERMINATOR, "plz"},
		{token.EOF, ""},
	}

	l := NewLexer(input) //eventually will be the lexer
	for i, exp := range tests {
		t := l.NextToken()

		if t.Type != exp.expectedType {
			test.Fatalf("Test #%d: incorrect token type. Expected %s, received %s.", i, exp.expectedType, t.Type)
		}

		if t.Literal != exp.expectedLiteral {
			test.Fatalf("Test #%d: incorrect token literal. Expected %s, received %s", i, exp.expectedLiteral, t.Literal)
		}
	}
}
