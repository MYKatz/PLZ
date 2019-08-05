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
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		if node.Value {
			return BOOL_TRUE
		}
		return BOOL_FALSE
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
