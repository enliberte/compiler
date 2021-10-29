package evaluator

import "compiler/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{Fn: lenBuiltinFn},
	"first": &object.Builtin{Fn: firstBuiltinFn},
	"last": &object.Builtin{Fn: lastBuiltinFn},
	"rest": &object.Builtin{Fn: restBuiltinFn},
	"push": &object.Builtin{Fn: pushBuiltinFn},
}

func lenBuiltinFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func firstBuiltinFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY. got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func lastBuiltinFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY. got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	l := len(arr.Elements)
	if l > 0 {
		return arr.Elements[l - 1]
	}

	return NULL
}

func restBuiltinFn(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY. got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	l := len(arr.Elements)
	if l > 0 {
		newElements := make([]object.Object, l - 1, l - 1)
		copy(newElements, arr.Elements[1:l])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func pushBuiltinFn(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, expected=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY. got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	l := len(arr.Elements)
	newElements := make([]object.Object, l + 1, l + 1)
	copy(newElements, arr.Elements)
	newElements[l] = args[1]

	return &object.Array{Elements: newElements}
}

