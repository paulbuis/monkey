// package monkey/ast declares the interface types for the nodes of an Abstract Syntax Tree for monkey
// package monkey.parser implements a parser that generates structs conforming to this
// interface in package monkey/parser/ast
package ast

import (
	"fmt"
	"monkey/token"
)

// The base Node interface
// conforms to interface fmt.Stringer
type Node interface {
	fmt.Stringer
	TokenLiteral() string
}

// All statement nodes implement this
type Statement interface {
	Node
	StatementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	ExpressionNode()
}

// Statements

//
type Program interface {
	Statement
	Statements() []Statement
}

//
type ReturnStatement interface {
	Statement
	Token() token.Token // the 'return' token
	ReturnValue() Expression
}

//
type ExpressionStatement interface {
	Statement
	Token() token.Token // the first token of the expression
	Expression() Expression
}

// see also monkey.parser.ast.BlockStatement
type BlockStatement interface {
	Statement
	Token() token.Token // the { token
	Statement(index int) Statement
	StatementCount() int
}

//
type ArrayLiteral interface {
	Token() token.Token // the '[' token
	Elements() []Expression
}

//
type Boolean interface {
	Expression
	Token() token.Token
	BooleanValue() bool
}

//
type CallExpression interface {
	Expression
	Token() token.Token   // The '(' token
	Function() Expression // Identifier or FunctionLiteral
	Argument(index int) Expression
	ArgumentCount() int
}

//
type FunctionLiteral interface {
	Expression
	Token() token.Token // The 'fn' token
	Parameter(index int) Identifier
	ParameterCount() int
	Body() BlockStatement
}

//
type HashLiteral interface {
	Expression
	Token() token.Token // the '{' token
	Keys(index int) Expression
	Values(index int) Expression
	PairCount() int
}

//
type Identifier interface {
	Expression
	Token() token.Token // the token.IDENT token
	IdentifierName() string
}

//
type IfExpression interface {
	Token() token.Token // The 'if' token
	Condition() Expression
	Consequence() BlockStatement
	Alternative() BlockStatement
}

//
type IndexExpression interface {
	Expression
	Token() token.Token // The [ token
	Left() Expression
	Index() Expression
}

//
type InfixExpression interface {
	Expression
	Token() token.Token // The operator token, e.g. +
	Left() Expression
	Operator() string
	Right() Expression
}

//
type IntegerLiteral interface {
	Expression
	Token() token.Token
	IntegerValue() int64
}

//
type LetStatement interface {
	Statement
	Token() token.Token // the token.LET token
	Name() Identifier
	LetValue() Expression
}

//
type MacroLiteral interface {
	Expression
	Token() token.Token // The 'macro' token
	Parameter(index int) Identifier
	ParameterCount() int
	Body() BlockStatement
}

//
type PrefixExpression interface {
	Expression
	Token() token.Token // The prefix token, e.g. !
	PrefixOperator() string
	Operand() Expression
}

//
type StringLiteral interface {
	Expression
	Token() token.Token
	StringValue() string
}
