package object

type BuiltinFunction func(args ...Object) Object

type ObjectType string

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"

	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ  = "STRING"

	RETURN_VALUE_OBJ = "RETURN_VALUE"

	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"

	ARRAY_OBJ = "ARRAY"
	HASH_OBJ  = "HASH"

	QUOTE_OBJ = "QUOTE"
	MACRO_OBJ = "MACRO"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	Object
	Inspect() string
	HashKey() HashKey
}

func IsError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}

func IsInteger(obj Object) bool {
	if obj != nil {
		return obj.Type() == INTEGER_OBJ
	}
	return false
}

func IsBoolean(obj Object) bool {
	if obj != nil {
		return obj.Type() == BOOLEAN_OBJ
	}
	return false
}

func IsArray(obj Object) bool {
	if obj != nil {
		return obj.Type() == ARRAY_OBJ
	}
	return false
}

func IsFunction(obj Object) bool {
	if obj != nil {
		return obj.Type() == FUNCTION_OBJ
	}
	return false
}

func IsHash(obj Object) bool {
	if obj != nil {
		return obj.Type() == HASH_OBJ
	}
	return false
}

func IsString(obj Object) bool {
	if obj != nil {
		return obj.Type() == STRING_OBJ
	}
	return false
}
