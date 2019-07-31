package parser

import (
	"fmt"
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

func TestIdentifierExpression(t *testing.T) {
	input := "foo plz"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	//checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	idnt, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", stmt.Expression)
	}
	if idnt.Value != "foo" {
		t.Fatalf("Incorrect identifier value, expected 'foo', received %q", idnt.Value)
	}
	if idnt.TokenLiteral() != "foo" {
		t.Fatalf("Incorrect TokenLiteral, expected foo, got %q", idnt.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "13 plz"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	//checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	idnt, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral, is %T", stmt.Expression)
	}
	if idnt.Value != 13 {
		t.Fatalf("Incorrect identifier value, expected 13, received %q", idnt.Value)
	}
	if idnt.TokenLiteral() != "13" {
		t.Fatalf("Incorrect TokenLiteral, expected '13', got %q", idnt.TokenLiteral())
	}
}

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5 plz", "!", 5},
		{"-1 plz", "-", 1},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("Incorrect number of program statements. Expected 1, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Stmt is incorrectly typed. Expected *ast.ExpressionStatement, got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp is incorrectly typed. Expected *ast.PrefixExpression, got %T", stmt)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator not %s. Received %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. Received %T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("Incorrect il value. Expected %d, received %d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d, got %s", value, integ.TokenLiteral())
		return false
	}

	return true
}
