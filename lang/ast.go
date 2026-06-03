package lang

import (
	"fmt"
)

func Parenthesize(expr Exp) string {
	switch e:= expr.(type) {
	case *BinaryExp:
		return fmt.Sprintf("(%s %s %s)", e.Op.Lexeme, Parenthesize(e.Left), Parenthesize(e.Right))
	case *GroupingExp:
		return fmt.Sprintf("(group %s)", Parenthesize(e.Exp))
	case *LiteralExp:
		return fmt.Sprintf("%v", e.Value)
	case *UnaryExp:
		return fmt.Sprintf("(%s %s)", e.Op.Lexeme, Parenthesize(e.Right))
	default:
		return "Unknown expression type"
	}

}