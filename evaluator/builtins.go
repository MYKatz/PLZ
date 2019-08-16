package evaluator

import (
	"fmt"

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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("Invalid argument to len, received %s", args[0].Type())
			}
		},
	},
	"append": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Incorrect number of arguments. Expected 2, received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Array{Elements: append(arg.Elements, args[1])}
			default:
				return newError("Invalid first argument to append, received %s", args[0].Type())
			}
		},
	},
	"peek": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Incorrect number of arguments. Expected 1, received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return arg.Elements[len(arg.Elements)-1]
			default:
				return newError("Invalid argument to peek, received %s", args[0].Type())
			}
		},
	},
	"first": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Incorrect number of arguments. Expected 1, received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return arg.Elements[0]
			default:
				return newError("Invalid argument to peek, received %s", args[0].Type())
			}
		},
	},
	"rest": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Incorrect number of arguments. Expected 1, received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Array{Elements: arg.Elements[1:]}
			default:
				return newError("Invalid argument to peek, received %s", args[0].Type())
			}
		},
	},
	"print": &object.BuiltIn{
		Fn: func(args ...object.Object) object.Object {
			for _, obj := range args {
				fmt.Println(obj.Inspect())
			}

			return NULL
		},
	},
}
