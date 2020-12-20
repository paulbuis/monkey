package environment

import (
	"monkey/context"
	"monkey/object"
	"monkey/stdlib"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := New(outer.ctx)
	env.outer = outer
	return env
}

func New(ctx *context.Context) *Environment {
	s := stdlib.New(ctx)
	return &Environment{store: s, outer: nil, ctx: ctx}
}

type Environment struct {
	store map[string]object.Object
	outer *Environment
	ctx   *context.Context
}

func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Context() *context.Context {
	return e.ctx
}

func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}
