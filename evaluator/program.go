package evaluator

import (
	"monkey/ast"
	"monkey/object"
	objectError "monkey/object/error"
	"monkey/object/function/environment"
	objectReturnValue "monkey/object/return_value"
)

// mutually recursive with Eval
func evalProgram(program *ast.Program, env *environment.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *objectReturnValue.ReturnValue:
			return result.Value()
		case *objectError.Error:
			return result
		}
	}

	return result
}
