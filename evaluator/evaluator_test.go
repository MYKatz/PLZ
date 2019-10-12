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
		{"let a be 5 plz let a be 5 plz a plz", 5},
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

func TestStringLiteral(t *testing.T) {
	input := `"Foo bar"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Object is incorrect type. Expected string, received %T", evaluated)
	}

	if str.Value != "Foo bar" {
		t.Fatalf("str is incorrect. Expected 'foo bar', received %q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Foo" + " " + "bar"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Object is incorrect type. Expected string, received %T", evaluated)
	}

	if str.Value != "Foo bar" {
		t.Fatalf("str is incorrect. Expected 'foo bar', received %q", str.Value)
	}
}

func TestLenBuiltinFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("foobar")`, 6},
		{`len("foo bar")`, 7},
		{`len("")`, 0},
		{`len(1)`, "Invalid argument to len, received INTEGER"},
		{`len("foo", "bar")`, "Incorrect number of arguments. Expected 1, received 2"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error, received %T", evaluated)
			}

			if errObj.Message != expected {
				t.Errorf("incorrect error message, expected %q, received %q", expected, errObj.Message)
			}
		}
	}
}

func TestBuiltinOverrideErrorHandling(t *testing.T) {
	input := "let len be function(x) please return x thanks plz"
	expected := "Invalid let statement: cannot override builtin function len"

	evaluated := testEval(input)
	errObj, ok := evaluated.(*object.Error)

	if !ok {
		t.Errorf("object is not Error, received %T", evaluated)
	}

	if errObj.Message != expected {
		t.Errorf("incorrect error message, expected %q, received %q", expected, errObj.Message)
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 + 2, 3 * 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Errorf("object is not Array, received %T", evaluated)
	}

	if len(result.Elements) != 3 {
		t.Errorf("incorrect number of array elements. Expected 3, got %d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 9)
}

func TestIndexExpression(t *testing.T) {
	input := "[1, 2, 3][1]"
	expected := 2

	evaluated := testEval(input)
	_, ok := evaluated.(*object.Integer)

	if !ok {
		t.Errorf("result is not integer. Received %T", evaluated)
	}

	testIntegerObject(t, evaluated, int64(expected))
}

func TestHashLiterals(t *testing.T) {
	input := `let two be "two" plz
	{
		"one": 5 - 4,
		two: 2,
		"thre" + "e": 9 / 3,
		4: 4,
		True: 100,
		False: -1
	}`

	evaluated := testEval(input)
	res, ok := evaluated.(*object.HashMap)
	if !ok {
		t.Errorf(evaluated.Inspect())
		t.Fatalf("Eval did not return hashmap, got %T", evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		BOOL_TRUE.HashKey():                        100,
		BOOL_FALSE.HashKey():                       -1,
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := res.Pairs[expectedKey]
		if !ok {
			t.Errorf("No pair")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestArrayAssignFunction(t *testing.T) {
	input := `
	let a be [1, 2, 4] plz
	let a be assign(a, 2, 3) plz
	a
	`

	evaluated := testEval(input)
	res, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("Eval did not return array, got %T", evaluated)
	}

	num, ok := res.Elements[2].(*object.Integer)
	if !ok {
		t.Fatalf("New element is not integer, is type %T", res.Elements[2])
	}
	if num.Value != int64(3) {
		t.Fatalf("Incorrect new array element. Expected 3, got %d", num.Value)
	}
}

func TestHashAssignFunction(t *testing.T) {
	input := `
	let a be {"one": 1, "two": 3} plz
	let a be assign(a, "two", 2) plz
	a
	`

	evaluated := testEval(input)
	res, ok := evaluated.(*object.HashMap)
	if !ok {
		t.Fatalf("Eval did not return hashmap, got %T", evaluated)
	}

	expectedKey := (&object.String{Value: "two"}).HashKey()

	pair, ok := res.Pairs[expectedKey]
	if !ok {
		t.Errorf("No pair found")
	}

	testIntegerObject(t, pair.Value, 2)
}

func TestMapReassign(t *testing.T) {
	input := `
		let m be {"one": 1, "two": 3} plz
		m["two"] = 2 plz
		m
	`

	evaluated := testEval(input)
	res, ok := evaluated.(*object.HashMap)
	if !ok {
		t.Errorf(evaluated.Inspect())
		t.Fatalf("Eval did not return hashmap, got %T", evaluated)
	}

	expectedKey := (&object.String{Value: "two"}).HashKey()

	pair, ok := res.Pairs[expectedKey]
	if !ok {
		t.Errorf("No pair found")
	}

	testIntegerObject(t, pair.Value, 2)
}

func TestMapAccessWithPeriod(t *testing.T) {
	input := `
	let m be {"one": 1, "two": 2} plz
	m.two plz
	`

	evaluated := testEval(input)
	hash, ok := evaluated.(*object.HashObject)

	if !ok {
		t.Errorf(evaluated.Inspect())
		t.Fatalf("Eval did not return hash object, got %T", evaluated)
	}

	integ, ok := hash.Inner.(*object.Integer)
	if !ok {
		t.Errorf(evaluated.Inspect())
		t.Fatalf("Hash object does not contain integer, got %T", integ)
	}

	testIntegerObject(t, integ, 2)
}
