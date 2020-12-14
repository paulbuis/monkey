package static_analysis

import (
	"fmt"
	"monkey/ast"
	"reflect"
)

type Inspector struct {
	stack    [1024]ast.Node
	stackTop int
}

func NewInspector() *Inspector {
	return &Inspector{}
}

func FindLets(root ast.Node) {
	visit := func(node ast.Node) {
		let, ok := node.(*ast.LetStatement)
		if !ok {
			return
		}
		fmt.Println(let.Name().Value())
	}
	inspector := NewInspector()
	inspector.Inspect(root, visit, nil)
}

func (in *Inspector) ScanToRoot(p reflect.Type) ast.Node {
	for i := in.stackTop - 1; i >= 0; i-- {
		node := in.stack[i]
		t := reflect.TypeOf(in.stack[i])
		if t == p {
			return node
		}

	}
	return nil
}

func (in *Inspector) Inspect(node ast.Node, before func(n ast.Node), after func(n ast.Node)) {
	in.stack[in.stackTop] = node
	in.stackTop++

	if before != nil {
		before(node)
	}
	switch node := node.(type) {
	case *ast.ArrayLiteral:
		for _, e := range node.Elements() {
			in.Inspect(e, before, after)
		}

	case *ast.BlockStatement:
		for _, s := range node.Statements() {
			in.Inspect(s, before, after)
		}

	case *ast.Boolean:
		return

	case *ast.CallExpression:
		for _, a := range node.Arguments() {
			in.Inspect(a, before, after)
		}
		in.Inspect(node.Function(), before, after)

	case *ast.ExpressionStatement:
		in.Inspect(node.Expression(), before, after)

	case *ast.FunctionLiteral:
		for _, p := range node.Parameters() {
			in.Inspect(p, before, after)
		}
		in.Inspect(node.Body(), before, after)

	case *ast.HashLiteral:
		for i, k := range node.Keys() {
			in.Inspect(k, before, after)
			in.Inspect(node.Values()[i], before, after)
		}

	case *ast.Identifier:
		return

	case *ast.IfExpression:
		in.Inspect(node.Condition(), before, after)
		in.Inspect(node.Consequence(), before, after)
		if node.Alternative() != nil {
			in.Inspect(node.Alternative(), before, after)
		}

	case *ast.IndexExpression:
		in.Inspect(node.Left(), before, after)
		in.Inspect(node.Index(), before, after)

	case *ast.InfixExpression:
		in.Inspect(node.Left(), before, after)
		in.Inspect(node.Right(), before, after)

	case *ast.IntegerLiteral:
		return

	case *ast.LetStatement:
		in.Inspect(node.Value(), before, after)

	case *ast.PrefixExpression:
		in.Inspect(node.Right(), before, after)

	case *ast.Program:
		for _, s := range node.Statements {
			in.Inspect(s, before, after)
		}

	case *ast.ReturnStatement:
		in.Inspect(node.ReturnValue(), before, after)

	case *ast.StringLiteral:
		return

	}

	if after != nil {
		after(node)
	}
}
