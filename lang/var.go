package lang

type Var struct {
	name Token
	init Exp
}

func (v *Var) Execute() error {
	if v.init != nil {
		value := v.init.Eval()
		current_env.Define(v.name, value)
		return nil
	}
	current_env.Define(v.name, nil)
	return nil
}

//Variable implements Exp
type Variable struct {
	name Token
}

func (v *Variable) Eval() any {
	return current_env.GetValue(v.name)
}

// Assignment
type Assignment struct {
	name Token
	val Exp
}

func (a *Assignment) Eval() any {
	val := a.val.Eval()
	current_env.Assign(a.name, val)
	return val
}