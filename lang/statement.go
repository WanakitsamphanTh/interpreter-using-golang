package lang

import "fmt"

type Statement interface {
	Execute() error
}

type PrintStatement struct {
	expr Exp
}

func NewPrintStatement(expr Exp) Statement {
	return &PrintStatement{expr}
}

func (s *PrintStatement) Execute() error {
	val := s.expr.Eval()
	if val == nil {
		return fmt.Errorf("Error")
	}
	fmt.Printf("%v\n", val)
	return nil
}

type ExpressionStatement struct {
	expr Exp
}

func NewExpressionStatement(expr Exp) Statement {
	return &ExpressionStatement{expr}
}

func (s *ExpressionStatement) Execute() error {
	s.expr.Eval()
	return nil
}