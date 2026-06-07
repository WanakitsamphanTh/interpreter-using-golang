package lang

import "fmt"

// Environment: mapping variables into variable statements (like Java)
type Environment struct {
	values map[string]any
	enclosing *Environment
}

func NewEnvironment(enclosing *Environment) *Environment {
	var e Environment
	e.values = make(map[string]any)
	e.enclosing = enclosing
	return &e
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

var current_env *Environment = NewEnvironment(nil)

func NewNestedEnvironment() *Environment {
	e := NewEnvironment(current_env)
	current_env = e
	return current_env
}

func RetractEnvironment() (*Environment, error) {
	prev := current_env.enclosing
	if prev == nil {
		return nil, fmt.Errorf("The current environment is the global environment.")
	}
	current_env = prev
	return current_env, nil
}