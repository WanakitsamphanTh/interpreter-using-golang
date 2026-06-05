package lang

import (
	"fmt"
)

func Parenthesize(node any) string {
	switch n:= node.(type) {
	case *BinaryExp:
		return fmt.Sprintf("(%s %s %s)", n.Op.Lexeme, Parenthesize(n.Left), Parenthesize(n.Right))
	case *GroupingExp:
		return fmt.Sprintf("(group %s)", Parenthesize(n.Exp))
	case *LiteralExp:
		return fmt.Sprintf("%v", n.Value)
	case *UnaryExp:
		return fmt.Sprintf("(%s %s)", n.Op.Lexeme, Parenthesize(n.Right))
	case *ExpressionStatement:
		return Parenthesize(n.expr)
	case *PrintStatement:
		return fmt.Sprintf("(Print %s)", Parenthesize(n.expr))
	case *Var:
		return fmt.Sprintf("(VarDecl %s %s)", n.name.Lexeme, Parenthesize(n.init))
	case *Assignment:
		return fmt.Sprintf("(Assign %s %s)", n.name.Lexeme, Parenthesize(n.val))
	case *Variable:
		return fmt.Sprintf("(Var %s)", n.name.Lexeme)
	case []Statement:
		return fmt.Sprintf("(%s)", listOfStatementsString(n))
	case Block:
		return fmt.Sprintf("(%s)", listOfStatementsString(n.statements))
	default:
		return "Unknown expression type"
	}

}

func listOfStatementsString(statements []Statement) string{
	str := "\n"
	for _, stmt := range statements {
		str = fmt.Sprintf("%s  %s\n", str,Parenthesize(stmt))
	}
	return str + ""
}