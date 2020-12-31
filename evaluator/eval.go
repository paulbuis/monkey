// package evaluator is responsible to traverse the AST to evaluate programs / statements / expressions
package evaluator

import (
	"monkey/ast"
	"monkey/evaluator/identifier"
	"monkey/object"
	"monkey/object/boolean"
	objectFunction "monkey/object/function"
	"monkey/object/function/environment"
	objectInteger "monkey/object/integer"
	objectString "monkey/object/string"
)

// Eval is the entry point for recursive traversal of an AST
func Eval(node ast.Node, env *environment.Environment) object.Object {
	switch node := node.(type) {

	// statements
	case ast.Program:
		return evalProgram(node, env) // in program.go

	case ast.BlockStatement:
		return evalBlockStatement(node, env) // in block_statement.go

	case ast.ExpressionStatement:
		return Eval(node.Expression(), env)

	case ast.ReturnStatement:
		return evalReturnStatement(node, env) // in return_statement.go

	case ast.LetStatement:
		val := Eval(node.LetValue(), env)
		if object.IsError(val) {
			return val
		}
		env.Set(node.Name().IdentifierName(), val)

	// Atomic Literal Expressions
	case ast.IntegerLiteral:
		return objectInteger.New(node.IntegerValue())

	case ast.StringLiteral:
		return objectString.New(node.StringValue())

	case ast.Boolean:
		return boolean.GetBoolean(node.BooleanValue())

	case ast.InfixExpression: // eval *before* prefix, Infix conforms to Prefix because of naming
		left := Eval(node.Left(), env)
		if object.IsError(left) {
			return left
		}
		right := Eval(node.Right(), env)
		if object.IsError(right) {
			return right
		}
		return applyInfixOperator(node.Operator(), left, right)

	case ast.PrefixExpression:
		operand := Eval(node.Operand(), env)
		if object.IsError(operand) {
			return operand
		}
		return applyPrefixOperator(node.PrefixOperator(), operand) // in prefix_operator.go

	case ast.IfExpression:
		return evalIfExpression(node, env) // mutually recursive with Eval

	// Atomic Expression
	case ast.Identifier:
		return identifier.EvalIdentifier(node, env) //

	case ast.FunctionLiteral:
		count := node.ParameterCount()
		parameters := make([]ast.Identifier, count)
		for index := 0; index < count; index++ {
			parameters[index] = node.Parameter(index)
		}
		body := node.Body()
		return objectFunction.New(parameters, body, env)

	case ast.CallExpression:
		//macro expansion disabled, hopefully temporary
		//if node.Function().TokenLiteral() == "quote" {
		//	return quote(node.Arguments()[0], env)
		//}

		function := Eval(node.Function(), env)
		if object.IsError(function) {
			return function
		}
		count := node.ArgumentCount()
		expressions := make([]ast.Expression, count)
		for index := 0; index < count; index++ {
			expressions[index] = node.Argument(index)
		}
		args := evalExpressions(expressions, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	// Composite Literals
	case ast.ArrayLiteral:
		return evalArrayLiteral(node, env)

	case ast.IndexExpression:
		left := Eval(node.Left(), env)
		if object.IsError(left) {
			return left
		}
		index := Eval(node.Index(), env)
		if object.IsError(index) {
			return index
		}
		return applyIndexOperator(left, index)

	case ast.HashLiteral:
		return evalHashLiteral(node, env)
	}

	return nil
}

// mutually recursive with Eval
// only used to evaluate an ArrayLiteral and CallExpression
func evalExpressions(
	exps []ast.Expression,
	env *environment.Environment,
) []object.Object {
	result := make([]object.Object, len(exps))

	for i, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result[i] = evaluated
	}

	return result
}
