package lang

type IfStatement struct {
	condition Exp
	thenBranch Statement
	elseBranch Statement
}

func (s *IfStatement) Execute() error {
	if isTruthy(s.condition.Eval()) {
		return s.thenBranch.Execute()
	} else {
		return s.elseBranch.Execute()
	}
}

type WhileStatement struct {
	condition Exp
	body Statement
}

func (s *WhileStatement) Execute() error {
	for isTruthy(s.condition.Eval()) {
		err := s.body.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}