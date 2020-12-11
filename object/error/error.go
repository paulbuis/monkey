package error

import (
	"fmt"
	"monkey/object"
)

type Error struct {
	message string
}

func New(format string, a ...interface{}) *Error {
	return &Error{message: fmt.Sprintf(format, a...)}
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Type() object.ObjectType {
	return object.ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.message
}
