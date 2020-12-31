package json

import (
	"encoding/json"
	"monkey/ast"
)

// Used to Marshal an AST to JSON
// if successful, the next normal step would be to send the
// []byte to an io.Writer
func Program(node ast.Node) ([]byte, error) {
	var m map[string]interface{}
	switch node := node.(type) {
	case ast.Expression:
		m = Expression(node)
	case ast.Statement:
		m = Statement(node)
	}
	return json.Marshal(m)

}

func Expression(expr ast.Expression) map[string]interface{} {
	switch expr := expr.(type) {
	case ast.ArrayLiteral:
		return ArrayLiteral(expr)

	case ast.Boolean:
		return BooleanLiteral(expr)

	case ast.CallExpression:
		return CallExpression(expr)

	case ast.FunctionLiteral:
		return FunctionLiteral(expr)

	case ast.HashLiteral:
		return HashLiteral(expr)

	case ast.Identifier:
		return Identifier(expr)

	case ast.IfExpression:
		return IfExpression(expr)

	case ast.IndexExpression:
		return IndexExpression(expr)

	case ast.InfixExpression:
		return InfixExpression(expr)

	case ast.IntegerLiteral:
		return IntegerLiteral(expr)

	case ast.PrefixExpression:
		return PrefixExpression(expr)

	case ast.StringLiteral:
		return StringLiteral(expr)
	}

	return nil
}

func Statement(statement ast.Statement) map[string]interface{} {
	switch statement := statement.(type) {
	case ast.ReturnStatement:
		return ReturnStatement(statement)

	case ast.BlockStatement:
		return BlockStatement(statement)

	case ast.ExpressionStatement:
		return ExpressionStatement(statement)

	case ast.LetStatement:
		return LetStatement(statement)
	}
	return nil
}

func ArrayLiteral(al ast.ArrayLiteral) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "ArrayLiteral"
	m["token"] = al.Token()
	count := len(al.Elements())
	elements := make([]interface{}, count)
	for index := 0; index < count; index++ {
		elements[index] = Expression(al.Elements()[index])
	}
	m["elements"] = elements

	return m
}

func BlockStatement(bs ast.BlockStatement) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "BlockStatement"
	m[" token"] = bs.Token()
	count := bs.StatementCount()
	statements := make([]interface{}, count)
	for index := 0; index < count; index++ {
		statements[index] = bs.Statement(index)
	}
	m["statements"] = statements
	return m
}

func BooleanLiteral(b ast.Boolean) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "Boolean"
	m[" token"] = b.Token()
	m["value"] = b.Value()
	return m
}

func CallExpression(ce ast.CallExpression) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "CallExpression"
	m[" token"] = ce.Token()
	m["function"] = Expression(ce.Function())
	count := ce.ArgumentCount()
	args := make([]interface{}, count)
	for index := 0; index < count; index++ {
		args[index] = Expression(ce.Argument(index))
	}
	m["function_arguments"] = args

	return m
}

func ExpressionStatement(es ast.ExpressionStatement) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "ExpressionStatement"
	m[" token"] = es.Token()
	m["expression"] = Expression(es.Expression())
	return m
}

func IntegerLiteral(il ast.IntegerLiteral) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "IntegerLiteral"
	m[" token"] = il.Token()
	return m
}

func FunctionLiteral(fl ast.FunctionLiteral) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "FunctionLiteral"
	m[" token"] = fl.Token()
	count := fl.ParameterCount()
	parameters := make([]string, count)
	for index := 0; index < count; index++ {
		parameters[index] = fl.Parameter(index).Value()
	}
	m["arguments"] = parameters
	m["body"] = BlockStatement(fl.Body())
	return m
}

func HashLiteral(hl ast.HashLiteral) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "HashLiteral"
	m[" token"] = hl.Token()
	count := hl.PairCount()
	keys := make([]interface{}, count)
	values := make([]interface{}, count)
	for index := 0; index < count; index++ {
		keys[index] = Expression(hl.Keys(index))
		values[index] = Expression(hl.Values(index))
	}
	m["keys"] = keys
	m["values"] = values
	return m
}

func Identifier(id ast.Identifier) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "Identifier"
	m[" token"] = id.Token()
	m["value"] = id.Value()
	return m
}

func IfExpression(ie ast.IfExpression) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "IfExpression"
	m[" token"] = ie.Token()
	m["condition"] = Expression(ie.Condition())
	m["consequence"] = Statement(ie.Consequence())
	if ie.Alternative() != nil {
		m["alternative"] = Statement(ie.Alternative())
	}
	return m
}

func IndexExpression(ie ast.IndexExpression) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "IndexExpression"
	m[" token"] = ie.Token()
	m["expression"] = Expression(ie.Left())
	m["index"] = Expression(ie.Index())
	return m
}

func InfixExpression(ie ast.InfixExpression) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "InfixExpression"
	m[" token"] = ie.Token()
	m["left"] = Expression(ie.Left())
	m["operator"] = ie.Operator()
	m["right"] = Expression(ie.Right())
	return m
}

func LetStatement(ls ast.LetStatement) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "LetStatement"
	m[" token"] = ls.Token()
	m["name"] = Identifier(ls.Name())
	m["value"] = Expression(ls.Value())
	return m
}

func PrefixExpression(pe ast.PrefixExpression) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "PrefixExpression"
	m[" token"] = pe.Token()
	m["operator"] = pe.Operator()
	m["right"] = Expression(pe.Right())
	return m
}

func ReturnStatement(rs ast.ReturnStatement) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "ReturnStatement"
	m[" token"] = rs.Token()
	m["returnValue"] = Expression(rs.ReturnValue())
	return m
}

func StringLiteral(sl ast.StringLiteral) map[string]interface{} {
	m := make(map[string]interface{})
	m["NodeType"] = "StringLiteral"
	m[" token"] = sl.Token()
	m["value"] = sl.Value()
	return m
}
