package integer

import (
	"math/big"
	"monkey/object"
	"monkey/object/boolean"
	objectError "monkey/object/error"
)

func ApplyIntegerInfixOperator(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*Integer).value
	rightVal := right.(*Integer).value
	result := big.NewInt(0)
	switch operator {
	case "+":
		result = result.Add(leftVal, rightVal)
		return &Integer{value: result}
	case "-":
		result = result.Sub(leftVal, rightVal)
		return &Integer{value: result}
	case "*":
		result = result.Mul(leftVal, rightVal)
		return &Integer{value: result}
	case "/":
		result = result.Div(leftVal, rightVal)
		return &Integer{value: result}
	case "<":
		return boolean.GetBoolean(leftVal.Cmp(rightVal) < 0)
	case ">":
		return boolean.GetBoolean(leftVal.Cmp(rightVal) > 0)
	case "==":
		return boolean.GetBoolean(leftVal.Cmp(rightVal) == 0)
	case "!=":
		return boolean.GetBoolean(leftVal.Cmp(rightVal) != 0)
	default:
		return objectError.New3("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func ApplyMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return objectError.New("unknown operator: -%s", right.Type())
	}

	value := right.(*Integer).value
	result := new(big.Int)
	result = result.Neg(value)
	return &Integer{value: result}
}
