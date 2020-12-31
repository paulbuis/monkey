package javascript

import (
	"fmt"
	"io"
	"monkey/ast"
)

// translate an AST to JavaScript
// if successful, the next normal step would be to send the
// []byte to an io.Writer
func Program(node ast.Node, out io.Writer) error {
	switch node := node.(type) {
	case ast.Expression:
		return Expression(node, out)
	case ast.Statement:
		return Statement(node, out)
	}
	return fmt.Errorf("unknown node type: %T", node)
}

func Expression(expr ast.Expression, out io.Writer) error {
	if expr == nil {
		return fmt.Errorf("nil expression")
	}
	switch expr := expr.(type) {
	case ast.ArrayLiteral:
		return ArrayLiteral(expr, out)

	case ast.Boolean:
		return BooleanLiteral(expr, out)

	case ast.CallExpression:
		return CallExpression(expr, out)

	case ast.FunctionLiteral:
		return FunctionLiteral(expr, out)

	case ast.HashLiteral:
		return HashLiteral(expr, out)

	case ast.Identifier:
		return Identifier(expr, out)

	case ast.IfExpression:
		return IfExpression(expr, out)

	case ast.IndexExpression:
		return IndexExpression(expr, out)

	case ast.InfixExpression:
		return InfixExpression(expr, out)

	case ast.IntegerLiteral:
		return IntegerLiteral(expr, out)

	case ast.PrefixExpression:
		return PrefixExpression(expr, out)

	case ast.StringLiteral:
		return StringLiteral(expr, out)
	}

	return fmt.Errorf("unkown expression type: %T", expr)
}

func ExpressionReturns(expr ast.Expression) bool {
	if expr == nil {
		return false
	}
	switch expr := expr.(type) {
	case ast.ArrayLiteral:
		for _, e := range expr.Elements() {
			if ExpressionReturns(e) {
				return true
			}
		}
		return false

	case ast.Boolean:
		return false

	case ast.CallExpression:
		if ExpressionReturns(expr.Function()) {
			return true
		}
		count := expr.ArgumentCount()
		for index := 0; index < count; index++ {
			if ExpressionReturns(expr.Argument(index)) {
				return true
			}
		}
		return false

	case ast.FunctionLiteral:
		return false

	case ast.HashLiteral:
		count := expr.PairCount()
		for index := 0; index < count; index++ {
			if ExpressionReturns(expr.Keys(index)) || ExpressionReturns(expr.Values(index)) {
				return true
			}
		}
		return false

	case ast.Identifier:
		return false

	case ast.IfExpression:
		return ExpressionReturns(expr.Condition()) ||
			StatementReturns(expr.Consequence()) ||
			StatementReturns(expr.Alternative())

	case ast.IndexExpression:
		return ExpressionReturns(expr.Left()) || ExpressionReturns(expr.Index())

	case ast.InfixExpression:
		return ExpressionReturns(expr.Left()) || ExpressionReturns(expr.Right())

	case ast.IntegerLiteral:
		return false

	case ast.PrefixExpression:
		return ExpressionReturns(expr.Right())

	case ast.StringLiteral:
		return false
	}

	return false
}

func Statement(statement ast.Statement, out io.Writer) error {
	switch statement := statement.(type) {
	case ast.ReturnStatement:
		return ReturnStatement(statement, out)

	case ast.BlockStatement:
		return BlockStatement(statement, out)

	case ast.ExpressionStatement:
		return ExpressionStatement(statement, out)

	case ast.LetStatement:
		return LetStatement(statement, out)
	}
	return fmt.Errorf("unknown statement type: %T", statement)
}

func ArrayLiteral(al ast.ArrayLiteral, out io.Writer) error {
	_, err := fmt.Fprint(out, "[")
	count := len(al.Elements())
	for index := 0; index < count; index++ {
		err = Expression(al.Elements()[index], out)
		if err != nil {
			return err
		}
		if index != count-1 {
			_, err = fmt.Fprint(out, ", ")
		}
	}
	_, err = fmt.Fprint(out, "]")
	return err
}

func BlockStatement(bs ast.BlockStatement, out io.Writer) error {
	_, err := fmt.Fprint(out, "\n{\n")
	count := bs.StatementCount()
	for index := 0; index < count; index++ {
		err = Statement(bs.Statement(index), out)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(out, "\n}\n")
	return err
}

func StatementReturns(statement ast.Statement) bool {
	if statement == nil {
		return false
	}
	switch statement := statement.(type) {
	case ast.ReturnStatement:
		return true

	case ast.BlockStatement:
		count := statement.StatementCount()
		for index := 0; index < count; index++ {
			if StatementReturns(statement.Statement(index)) {
				return true
			}
		}
		return false

	case ast.ExpressionStatement:
		return ExpressionReturns(statement.Expression())

	case ast.LetStatement:
		return false
	}
	return false
}

func BlockExpression(bs ast.BlockStatement, out io.Writer) error {
	_, err := fmt.Fprint(out, "(")

	count := bs.StatementCount()
	for index := 0; index < count; index++ {
		err = Statement(bs.Statement(index), out)
		if err != nil {
			return err
		}
		if index != count-1 {
			_, err = fmt.Fprint(out, ", ")
		}
	}
	_, err = fmt.Fprint(out, ")")
	return err
}

func BooleanLiteral(b ast.Boolean, out io.Writer) error {
	var err error
	if b.Value() {
		_, err = fmt.Fprint(out, "true")
	} else {
		_, err = fmt.Fprint(out, "false")
	}
	return err
}

func CallExpression(ce ast.CallExpression, out io.Writer) error {
	err := Expression(ce.Function(), out)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, "(")
	count := ce.ArgumentCount()

	for index := 0; index < count; index++ {
		err = Expression(ce.Argument(index), out)
		if index != count-1 {
			_, err = fmt.Fprint(out, ", ")
		}
	}
	_, err = fmt.Fprint(out, ")")
	return err
}

func ExpressionStatement(es ast.ExpressionStatement, out io.Writer) error {
	err := Expression(es.Expression(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, ";\n")
	return err
}

func IntegerLiteral(il ast.IntegerLiteral, out io.Writer) error {
	_, err := fmt.Fprintf(out, " %d ", il.Value())
	return err
}

func FunctionLiteral(fl ast.FunctionLiteral, out io.Writer) error {
	_, err := fmt.Fprint(out, "function(")
	if err != nil {
		return err
	}
	count := fl.ParameterCount()
	for index := 0; index < count; index++ {
		_, err = fmt.Fprint(out, fl.Parameter(index).Value())
		if err != nil {
			return err
		}
		if index != count-1 {
			_, err = fmt.Fprint(out, ", ")
		}
	}
	_, err = fmt.Fprint(out, ") {")

	return BlockStatement(fl.Body(), out)
}

func HashLiteral(hl ast.HashLiteral, out io.Writer) error {
	_, err := fmt.Fprint(out, "{")
	if err != nil {
		return err
	}
	count := hl.PairCount()
	for index := 0; index < count; index++ {
		err = Expression(hl.Keys(index), out)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(out, " : ")
		err = Expression(hl.Values(index), out)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(out, "}")
	return err
}

func Identifier(id ast.Identifier, out io.Writer) error {
	_, err := fmt.Fprintf(out, " %s ", id.Value())
	return err
}

func IfExpression(ie ast.IfExpression, out io.Writer) error {
	err := Expression(ie.Condition(), out)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, " ? ")

	err = BlockExpression(ie.Consequence(), out)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, " : ")

	if ie.Alternative() != nil {
		err = BlockExpression(ie.Alternative(), out)
		if err != nil {
			return err
		}
	} else {
		_, err = fmt.Fprint(out, "null")
	}
	return nil
}

func IndexExpression(ie ast.IndexExpression, out io.Writer) error {
	err := Expression(ie.Left(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "[")
	err = Expression(ie.Index(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "]")
	return err
}

func InfixExpression(ie ast.InfixExpression, out io.Writer) error {
	_, err := fmt.Fprint(out, "(")
	err = Expression(ie.Left(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, " %s ", ie.Operator())
	err = Expression(ie.Right(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "(")
	return nil
}

func LetStatement(ls ast.LetStatement, out io.Writer) error {
	_, err := fmt.Fprintf(out, "\nlet %s = ", ls.Name())
	if err != nil {
		return err
	}
	err = Expression(ls.Value(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, ";\n")
	return err
}

func PrefixExpression(pe ast.PrefixExpression, out io.Writer) error {
	_, err := fmt.Fprintf(out, "(%s(", pe.Operator())
	err = Expression(pe.Right(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "))")
	return err
}

func ReturnStatement(rs ast.ReturnStatement, out io.Writer) error {
	_, err := fmt.Fprint(out, "return ")
	err = Expression(rs.ReturnValue(), out)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, ";\n")
	return err
}

func StringLiteral(sl ast.StringLiteral, out io.Writer) error {
	// TODO: handle escapes
	_, err := fmt.Fprint(out, "\""+sl.Value()+"\"")
	return err
}
