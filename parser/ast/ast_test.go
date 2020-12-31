package ast

import (
	"monkey/ast"
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		statements: []ast.Statement{
			NewLetStatement(
				token.New(token.LET, "let"),
				NewIdentifier(
					token.New(token.IDENT, "myVar"),
					"myVar",
				),
				NewIdentifier(
					token.New(token.IDENT, "anotherVar"),
					"anotherVar",
				),
			),
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
