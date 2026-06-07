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
		if s.elseBranch != nil {
			return s.elseBranch.Execute()
		}
		return nil
	}
}

type WhileStatement struct {
	condition Exp
	body Statement
}

func (s *WhileStatement) Execute() error {
	current_env.Define("@looping", true)
	defer current_env.Assign("@looping", false)
	
	for isTruthy(s.condition.Eval()) {
		err := s.body.Execute()
		if current_env.GetValue("@terminated").(bool) {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Break & Continue
type BreakStmt struct {
	keyword Token
}

func (s *BreakStmt) Execute() error {
	return current_env.TerminateLoop()
}

type SkipStmt struct {
	keyword Token
}

func (s *SkipStmt) Execute() error {
	return nil
}