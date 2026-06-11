package lang

type Var struct {
	name Token
	init Exp
}

func (v *Var) Execute() disruptive {
	if v.init != nil {
		value := v.init.Eval()
		current_env.Define(v.name.Lexeme, value)
		return nil
	}
	current_env.Define(v.name.Lexeme, nil)
	return nil
}

// Variable implements Exp
type Variable struct {
	name Token
}

func (v *Variable) Eval() any {
	return LookUpVariable(v.name.Lexeme, v)
}

// Assignment
type Assignment struct {
	name Token
	val  Exp
}

func (a *Assignment) Eval() any {
	val := a.val.Eval()

	distance, ok := locals[a]
	if !ok {
		err := current_env.assignAt(distance, a.name.Lexeme, val)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := global.Assign(a.name.Lexeme, val)
		if err != nil {
			panic(err.Error())
		}
	}

	return val
}
