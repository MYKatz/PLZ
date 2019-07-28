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
	if len(p.Errors()) > 0 {
		for i, e := range p.Errors() {
			t.Errorf("Parser error #%d: %s", i+1, e)
		}
		t.FailNow()
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

func TestReturnStatements(t *testing.T) {
	input := `
	return 1 plz
	return 10 plz
	return 3141 plz
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()

	if prog == nil {
		t.Fatalf("Error: parseprogram returned nil")
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("Error: parseprogram returned incorrect amount of statements. Expected %d, received %d", 3, len(prog.Statements))
	}

	if len(p.Errors()) > 0 {
		for i, err := range p.Errors() {
			t.Errorf("Parser error #%d: %s", i+1, err)
		}
		t.FailNow()
	}

	for _, stmt := range prog.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("Incorrect type, stmt not *ast.ReturnStatement, instead %T", stmt)
		}
		if retStmt.TokenLiteral() != "return" {
			t.Errorf("retStmt.TokenLiteral() not 'return', got %q", retStmt.TokenLiteral())
		}
	}
}
