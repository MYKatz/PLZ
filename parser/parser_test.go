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
	let b be True plz
	let c be foo plz
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
		input              string
		expectedValue      interface{}
	}{
		{"a", "let a be 10 plz", 5},
		{"b", "let b be True plz", true},
		{"c", "let c be foo plz", "foo"},
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

func TestBooleanExpression(t *testing.T) {
	input := "True plz"
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

	idnt, ok := stmt.Expression.(*ast.Boolean)
	if !ok {
		t.Fatalf("expression not *ast.Boolean, is %T", stmt.Expression)
	}
	if idnt.Value != true {
		t.Fatalf("Incorrect identifier value, expected true, received %t", idnt.Value)
	}
	if idnt.TokenLiteral() != "True" {
		t.Fatalf("Incorrect TokenLiteral, expected 'True', got %q", idnt.TokenLiteral())
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

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. Received %T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("Incorrect exp value. Expected %s, received %s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("integ.TokenLiteral not %s, got %s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}

	t.Errorf("Type of exp not handled. Received type %T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not *ast.InfixExpression got %T", exp)
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not %s, got %q", operator, opExp.Operator)
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		leftVal  int64
		operator string
		rightVal int64
	}{
		{"5 + 5 plz", 5, "+", 5},
		{"5 - 5 plz", 5, "-", 5},
		{"5 * 5 plz", 5, "*", 5},
		{"5 / 5 plz", 5, "/", 5},
		{"5 > 5 plz", 5, ">", 5},
		{"5 < 5 plz", 5, "<", 5},
		{"5 == 5 plz", 5, "==", 5},
		{"5 != 5 plz", 5, "!=", 5},
	}

	for _, tt := range infixTests {
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

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is incorrectly typed. Expected *ast.PrefixExpression, got %T", stmt)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftVal) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator not %s. Received %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightVal) {
			return
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("expression not *ast.IfExpression, is %T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Fatalf("Consequence length is not 1, got %d", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("consequence statement not *ast.ExpressionStatement, is %T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil, was %+v", exp.Alternative)
	}

}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("expression not *ast.IfExpression, is %T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Fatalf("Consequence length is not 1, got %d", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("consequence statement not *ast.ExpressionStatement, is %T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Fatalf("Consequence length is not 1, got %d", len(exp.Consequence.Statements))
	}

	alti, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("alternative statement not *ast.ExpressionStatement, is %T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, alti.Expression, "y") {
		return
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `function(x, y) {x + y}`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	fn, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression not *ast.FunctionLiteral, is %T", stmt.Expression)
	}

	if len(fn.Parameters) != 2 {
		t.Fatalf("Number of function parameters incorrect. Expected 2, received %d", len(fn.Parameters))
	}

	testLiteralExpression(t, fn.Parameters[0], "x")
	testLiteralExpression(t, fn.Parameters[1], "y")

	if len(fn.Body.Statements) != 1 {
		t.Fatalf("Number of function body statements incorrected. Expected 1, received %d", len(fn.Body.Statements))
	}

	bodyStmt, ok := fn.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Function body statement is not ast.ExpressionStatement, got %T", fn.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestCallExpression(t *testing.T) {
	input := `add(2 + 3, 6 * 2, 5) plz`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	fn, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression not *ast.FunctionLiteral, is %T", stmt.Expression)
	}

	if !testIdentifier(t, fn.Function, "add") {
		return
	}

	if len(fn.Arguments) != 3 {
		t.Fatalf("Number of function arguments incorrect. Expected 3, received %d", len(fn.Arguments))
	}

	testInfixExpression(t, fn.Arguments[0], 2, "+", 3)
	testInfixExpression(t, fn.Arguments[1], 6, "*", 2)
	testLiteralExpression(t, fn.Arguments[2], 5)

}

func TestStringExpression(t *testing.T) {
	input := `"plz is awesome!" plz`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has incorrect number of statements. Got %d, expected 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression not *ast.ExpressionStatement, is %T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("expression not *ast.StringLiteral, is %T", stmt.Expression)
	}
	if literal.Value != "plz is awesome!" {
		t.Fatalf("Incorrect identifier value, expected %q, received %q", "plz is awesome!", literal.Value)
	}
}
