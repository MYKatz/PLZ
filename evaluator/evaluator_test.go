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

	return Eval(program)
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
		{"if (True) { 1 }", 1},
		{"if (10) { 10 }", 10},
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
