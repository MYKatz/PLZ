package evaluator

import (
	"testing"

	"github.com/MYKatz/PLZ/lexer"
	"github.com/MYKatz/PLZ/object"
	"github.com/MYKatz/PLZ/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"--10", 10},
		{"2 + 2 + 2", 6},
		{"5 - 3", 2},
		{"5 + 7 * 2", 19},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (1 + 3)", 8},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"False", false},
		{"True", true},
		{"5 == 5", true},
		{"3 > 10", false},
		{"3 < 10", true},
		{"5 != 5", false},
		{"7 != 2", true},
		{"True == True", true},
		{"True == False", false},
		{"(3 > 10) != False", false},
		{"True == (5 == 5)", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(inp string) object.Object {
	l := lexer.NewLexer(inp)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, evaluated object.Object, expected int64) bool {
	result, ok := evaluated.(*object.Integer)
	if !ok {
		t.Errorf("Object is not Integer. Received %T", evaluated)
		return false
	}
	if result.Value != expected {
		t.Errorf("Object has wrong value. Expected %d, received %d", expected, result.Value)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, evaluated object.Object, expected bool) bool {
	result, ok := evaluated.(*object.Boolean)
	if !ok {
		t.Errorf("Object is not Boolean. Received %T", evaluated)
		return false
	}
	if result.Value != expected {
		t.Errorf("Object has wrong value. Expected %t, received %t", expected, result.Value)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!True", false},
		{"!False", true},
		{"!5", false},
		{"!!True", true},
		{"!!False", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (True) please 1 thanks", 1},
		{"if (10) please 10 thanks", 10},
		{"if (False) { 50 }", nil},
		{"if (10 > 2) { 25 }", 25},
		{"if (10 < 2) { 10 }", nil},
		{"if (10 < 2) { 1 } else { 5 }", 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("Object is not NULL, got %T", obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10 plz", 10},
		{"return 15 plz", 15},
		{"return 2 + 8 plz 9 plz", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + True plz", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + True plz 5 plz", "type mismatch: INTEGER + BOOLEAN"},
		{"-True plz", "unknown operator: -BOOLEAN"},
		{"True + False plz", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foo plz", "identifier not found: foo"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		err, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. Received %T", evaluated)
			continue
		}

		if err.Message != tt.expected {
			t.Errorf("wrong error message. Expected %q, received %q", tt.expected, err.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a be 5 plz a plz", 5},
		{"let a be 5 * 5 plz a plz", 25},
		{"let a be 5 plz let b be a plz b plz", 5},
		{"let a be 5 plz let b be a plz let c be a + b + 5 plz c plz", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "function(x, y) please x + y plz thanks plz"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("Object is not a function. Received type %T", evaluated)
	}

	if len(fn.Parameters) != 2 {
		t.Fatalf("Function has wrong number of parameters. Expected 2, received %d", len(fn.Parameters))
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("First parameter is not 'x', got %q", fn.Parameters[0].String())
	}

	if fn.Parameters[1].String() != "y" {
		t.Fatalf("First parameter is not 'y', got %q", fn.Parameters[0].String())
	}

	if fn.Body.String() != "(x+y)" {
		t.Fatalf("Incorrect function body. Expected %s, received %s", "(x+y)", fn.Body.String())
	}
}

func TestFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let id be function(x) please x plz thanks plz id(5) plz", 5},
		{"let id be function(x) please return x plz thanks plz id(5) plz", 5},
		{"let double be function(x) please return x*2 plz thanks plz double(double(2))", 8},
		{"function(x) please return x+1 plz thanks(5) plz", 6},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
