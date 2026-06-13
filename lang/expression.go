package lang

import "fmt"

type Exp interface {
	Resolveable
	Eval() any
}

// UnaryExp
type UnaryExp struct {
	Op    Token
	Right Exp
}

func (u UnaryExp) Eval() any {
	switch u.Op.Type {
	case MINUS:
		return -u.Right.Eval().(float64)
	case BANG:
		return isTruthy(u.Right.Eval())
	}
	return u.Right.Eval()
}

// BinaryExp
type BinaryExp struct {
	Left  Exp
	Op    Token
	Right Exp
}

func (b BinaryExp) Eval() any {
	left := b.Left.Eval()
	right := b.Right.Eval()

	switch b.Op.Type {
	case PLUS:
		switch l := left.(type) {
		case float64:
			r, ok := right.(float64)
			if !ok {
				panic("Unable to convert variable to float64")
			}
			return l + r
		case string:
			/*r, ok := right.(string)
			if !ok {
				panic("Unable to convert variable to string")
			}
			return l + r*/
			return fmt.Sprintf("%s%v", l, right)
		}
	case MINUS:
		return left.(float64) - right.(float64)
	case STAR:
		return left.(float64) * right.(float64)
	case SLASH:
		if right.(float64) == 0.0 {
			panic("Can't divide by zero")
		}
		return left.(float64) / right.(float64)
	case GREATER:
		return left.(float64) > right.(float64)
	case GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case LESS:
		return left.(float64) < right.(float64)
	case LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case EQUAL_EQUAL:
		return isEqual(left, right)
	case BANG_EQUAL:
		return !isEqual(left, right)
	default:
		panic("unknown operator")
	}
	return nil
}

// GroupingExp
type GroupingExp struct {
	Exp Exp
}

func (g GroupingExp) Eval() any {
	return g.Exp.Eval()
}

// LiteralExp
type LiteralExp struct {
	Value any
}

func (l LiteralExp) Eval() any {
	return l.Value
}

// Miscellaneous

func isTruthy(obj any) bool {
	val, ok := obj.(bool)
	if !ok {
		return false
	}
	return val
}

func isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

type LogicalExpression struct {
	Op    Token
	left  Exp
	right Exp
}

func (l *LogicalExpression) Eval() any {
	if l.Op.Type == OR {
		if isTruthy(l.left.Eval()) {
			return l.left
		}
	} else {
		if !isTruthy(l.left.Eval()) {
			return l.left
		}
	}
	return l.right
}
