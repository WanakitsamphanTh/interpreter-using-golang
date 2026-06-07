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
		val, ok := n.Value.(string)
		if ok {
			return fmt.Sprintf("\"%v\"", val)
		} else {
			return fmt.Sprintf("%v", n.Value)
		}
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
		return fmt.Sprintf("%s", n.name.Lexeme)
	case []Statement:
		return fmt.Sprintf("(%s)", listOfStatementsString(n))
	case *Block:
		return fmt.Sprintf("(%s)", listOfStatementsString(n.statements))
	case *IfStatement:
		return fmt.Sprintf("(If %s %s %s)", Parenthesize(n.condition), Parenthesize(n.thenBranch), Parenthesize(n.elseBranch))
	case *WhileStatement:
		return fmt.Sprintf("(While %s %s)", Parenthesize(n.condition), Parenthesize(n.body))
	case *FnDecl:
		return fmt.Sprintf("(Fun %s(%v) %s)", n.name.Lexeme, listParamsString(n.params), Parenthesize(n.body))
	case *FnCall:
		return fmt.Sprintf("(%s%s)", Parenthesize(n.callee), listParamsString(n.params))
	case *Return:
		return fmt.Sprintf("(Return %s)", Parenthesize(n.val))
	default:
		return fmt.Sprintf("(%T)", n)
	}

}

func listOfStatementsString(statements []Statement) string{
	str := ""
	for _, stmt := range statements {
		str = str + Parenthesize(stmt)
	}
	return str + ""
}

func listParamsString(param_list any) string {
	str := ""
	switch params := param_list.(type) {
	case []Token:
		for i, param := range params {
			if i != 0 {
				str = str + ", "
			}
			str = str + param.Lexeme
		}
	case []Exp:
		for _, param := range params {
			str = str + " " + Parenthesize(param)
		}
	default:
		return "unknown type";
	}
	return str
}