package lang

type IfStatement struct {
	condition  Exp
	thenBranch Statement
	elseBranch Statement
}

func (s *IfStatement) Execute() disruptive {
	if isTruthy(s.condition.Eval()) {
		return s.thenBranch.Execute()
	} else {
		if s.elseBranch != nil {
			return s.elseBranch.Execute()
		}
		return nil
	}
}

type WhileStatement struct {
	condition Exp
	body      Statement
	increment Statement
}

func (s *WhileStatement) Execute() disruptive {
	for isTruthy(s.condition.Eval()) {
		err := s.body.Execute()
		if err != nil {
			switch disruption := err.(type) {
			case *BreakStmt:
				return nil
			case *SkipStmt:
				if s.increment != nil {
					err := s.increment.Execute()
					if err != nil {
						return err
					}
				}
				continue
			default:
				return disruption
			}
		}
		if s.increment != nil {
			err := s.increment.Execute()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Break & Continue
type BreakStmt struct {
	keyword Token
}

func (s *BreakStmt) Execute() disruptive {
	return s
}

type SkipStmt struct {
	keyword Token
}

func (s *SkipStmt) Execute() disruptive {
	return s
}
