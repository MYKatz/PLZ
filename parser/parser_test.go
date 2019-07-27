package parser

import (
	"testing"

	"github.com/MYKatz/PLZ/ast"
	"github.com/MYKatz/PLZ/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let a be 10 plz
	let b be 2 plz
	let c be 538 plz
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Error: Parseprogram returned null")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Error: incorrect amount of program statements. Expected 3, received %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"c"},
	}

	for i, te := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, te.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, ident string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Incorrect token literal. Expected 'let', received %s", s.TokenLiteral())
	}
	letStatement, ok := s.(*ast.LetStatement) //check type of s
	if !ok {
		t.Errorf("s is wrongly typed. Expected ast.LetStatement, received %T", s)
		return false
	}

	if letStatement.Name.Value != ident {
		t.Errorf("letStatement.Name.Value incorrect. Expected %s, received %s", ident, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != ident {
		t.Errorf("letStatement.Name.TokenLiteral() incorrect. Expected %s, received %s", ident, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
