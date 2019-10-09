package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/MYKatz/PLZ/ast"
)

type ObjectType string

type BuiltinFunction func(args ...Object) Object

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	HASHOBJ_OBJ      = "HASHOBJ"
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

//function

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (fn *Function) Inspect() string {
	return "" //TODO - function inspect
}

func (fn *Function) Type() string {
	return FUNCTION_OBJ
}

//string

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() string {
	return STRING_OBJ
}

//builtin

type BuiltIn struct {
	Fn BuiltinFunction
}

func (f *BuiltIn) Inspect() string {
	return "builtin function"
}

func (f *BuiltIn) Type() string {
	return BUILTIN_OBJ
}

//array

type Array struct {
	Elements []Object
}

func (a *Array) Inspect() string {
	var output bytes.Buffer

	elements := []string{}

	for _, el := range a.Elements {
		elements = append(elements, el.Inspect())
	}

	output.WriteString("[")
	output.WriteString(strings.Join(elements, ", "))
	output.WriteString("]")

	return output.String()
}

func (a *Array) Type() string {
	return ARRAY_OBJ
}

//hashkey

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

//func to turn objects into hashkey

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: BOOLEAN_OBJ, Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: INTEGER_OBJ, Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: STRING_OBJ, Value: h.Sum64()}
}

//hashes

type HashPair struct {
	Key   Object
	Value Object
}

type HashMap struct {
	Pairs map[HashKey]HashPair
}

func (hm *HashMap) Inspect() string {

	var output bytes.Buffer

	output.WriteString("{")
	for _, pair := range hm.Pairs {
		output.WriteString(pair.Key.Inspect() + ":" + pair.Value.Inspect() + ",")
	}
	output.WriteString("}")

	return output.String()
}

func (hm *HashMap) Type() string {
	return HASH_OBJ
}

//wrapper for objects accessed via Hash indexing
type HashObject struct {
	Hash     *HashMap
	PlainKey Object
	Key      HashKey
	Inner    Object
}

func (ho *HashObject) Type() string {
	return HASHOBJ_OBJ
}

func (ho *HashObject) Inspect() string {
	return ho.Inner.Inspect()
}
