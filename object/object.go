package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() string
	Inspect() string
}

//integer

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() string {
	return INTEGER_OBJ
}

//boolean

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() string {
	return BOOLEAN_OBJ
}

//null

type Null struct{}

func (n *Null) Inspect() string {
	return "Null"
}

func (n *Null) Type() string {
	return NULL_OBJ
}

//return

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() string {
	return RETURN_VALUE_OBJ
}

//error

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

func (e *Error) Type() string {
	return ERROR_OBJ
}
