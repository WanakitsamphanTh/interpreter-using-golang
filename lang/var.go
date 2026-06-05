package lang

// Var implements Statement
type Var struct {
	name Token
	init Exp
	env *Environment
}

func (v *Var) Execute() error {
	if v.init != nil {
		value := v.init.Eval()
		v.env.Define(v.name, value)
		return nil
	}
	v.env.Define(v.name, nil)
	return nil
}

//Variable implements Exp
type Variable struct {
	name Token
	env *Environment
}

func (v *Variable) Eval() any {
	return v.env.GetValue(v.name)
}

// Environment: mapping variables into variable statements (like Java)
type Environment struct {
	values map[string]any
	enclosing *Environment
}

func NewEnvironment() Environment {
	var e Environment
	e.values = make(map[string]any)
	return e
}

func NewNestedEnvironment(enclosing *Environment) Environment {
	e := NewEnvironment()
	e.enclosing = enclosing
	return e
}

func (e *Environment) Define(name Token, val any) {
	e.values[name.Lexeme] = val
}

func (e *Environment) Assign(name Token, val any) error {
	_, ok := e.values[name.Lexeme]
	if ok {
		e.values[name.Lexeme] = val
		return nil
	}

	if !ok && e.enclosing != nil {
		err := e.enclosing.Assign(name,val)
		if err == nil {
			return nil
		}
	}

	return &RuntimeError{name, "Undefined variable"}
}

func (e *Environment) GetValue(name Token) any {
	val, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing != nil {
			val = e.enclosing.GetValue(name)
			return val
		}
		panic("Undefined variable " + name.Lexeme)
	}
	return val
}

// Assignment
type Assignment struct {
	name Token
	val Exp
	env *Environment
}

func (a *Assignment) Eval() any {
	val := a.val.Eval()
	a.env.Assign(a.name, val)
	return val
}