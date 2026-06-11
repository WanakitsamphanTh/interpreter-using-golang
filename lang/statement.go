package lang

type Statement interface {
	Resolveable
	Execute() disruptive
}

/*type PrintStatement struct {
	expr Exp
}

func NewPrintStatement(expr Exp) Statement {
	return &PrintStatement{expr}
}*/

/*func (s *PrintStatement) Execute() disruptive {
	val := s.expr.Eval()
	if val == nil {
		return fmt.Errorf("Undefined value")
	}
	fmt.Printf("%v\n", val)
	return nil
}*/

type ExpressionStatement struct {
	expr Exp
}

func NewExpressionStatement(expr Exp) Statement {
	return &ExpressionStatement{expr}
}

func (s *ExpressionStatement) Execute() disruptive {
	s.expr.Eval()
	return nil
}
