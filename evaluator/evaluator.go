package evaluator

import (
	"github.com/MYKatz/PLZ/ast"
	"github.com/MYKatz/PLZ/object"
)

var (
	BOOL_TRUE  = &object.Boolean{Value: true}
	BOOL_FALSE = &object.Boolean{Value: false}
	NULL       = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		if node.Value {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node)
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	}
	return nil
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case BOOL_TRUE:
		return BOOL_FALSE
	case BOOL_FALSE:
		return BOOL_TRUE
	case NULL:
		return BOOL_TRUE
	default:
		return BOOL_FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		if left == right {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	case operator == "!=":
		if left != right {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	default:
		return nil
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "==":
		if leftVal == rightVal {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	case ">":
		if leftVal > rightVal {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	case "<":
		if leftVal < rightVal {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	case "!=":
		if leftVal != rightVal {
			return BOOL_TRUE
		}
		return BOOL_FALSE
	default:
		return NULL
	}
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

//TODO: make certain other things truthy ie 0, empty string, etc.
func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case BOOL_TRUE:
		return true
	case BOOL_FALSE:
		return false
	default:
		return true
	}
}
