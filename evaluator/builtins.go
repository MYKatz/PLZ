package evaluator

import (
	"github.com/MYKatz/PLZ/object"
)

var builtins = map[string]*object.BuiltIn{
	"len": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Incorrect number of arguments. Expected 1, received %d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Invalid argument to len, received %s", args[0].Type())
			}
		},
	},
}
