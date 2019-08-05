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
