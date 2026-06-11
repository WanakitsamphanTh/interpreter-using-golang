package lang

import "fmt"

// Environment: mapping variables into variable statements (like Java)
type Environment struct {
	values        map[string]any
	enclosing     *Environment
	functionBound bool
}

func NewEnvironment(enclosing *Environment, functionBound bool) *Environment {
	var e Environment
	e.values = make(map[string]any)
	e.enclosing = enclosing
	e.functionBound = functionBound

	return &e
}

func (e *Environment) Define(name string, val any) {
	e.values[name] = val
}

func (e *Environment) Assign(name string, val any) error {

	_, ok := e.values[name]
	if ok {
		e.values[name] = val
		return nil
	}

	if !ok && e.enclosing != nil {
		err := e.enclosing.Assign(name, val)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%s : Undefined variable", name)
}

func (e *Environment) GetValue(name string) any {

	val, ok := e.values[name]
	if !ok {
		if e.enclosing != nil {
			val = e.enclosing.GetValue(name)
			return val
		}
		panic("Undefined variable " + name)
	}
	return val
}

func (e *Environment) GetAt(distance int, name string) any {
	return e.ancestor(distance).values[name]
}

func (e *Environment) ancestor(distance int) *Environment {
	env := e
	for i := 0; i < distance; i++ {
		env = env.enclosing
	}
	return env
}

var global *Environment = NewEnvironment(nil, false)
var current_env *Environment = global

func NewNestedEnvironment(functionBound bool) *Environment {
	e := NewEnvironment(current_env, functionBound)
	current_env = e
	return current_env
}

func RetractEnvironment() (*Environment, error) {
	if current_env == global {
		return nil, fmt.Errorf("The current environment is the global environment.")
	}
	current_env = current_env.enclosing
	return current_env, nil
}
