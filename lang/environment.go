package lang

import "fmt"

// Environment: mapping variables into variable statements (like Java)
type Environment struct {
	values map[string]any
	enclosing *Environment
	functionBound bool
}

func NewEnvironment(enclosing *Environment, functionBound bool) *Environment {
	var e Environment
	e.values = make(map[string]any)
	e.enclosing = enclosing
	e.functionBound = functionBound
	if e.functionBound {
		e.Define("ret_val", nil)
	}
	return &e
}

func (e *Environment) Define(name string, val any) {
	e.values[name] = val
}

func (e *Environment) Assign(name string, val any) error {
	if name == "ret_val" && !e.functionBound {
		if e == global{
			panic("Cannot assign return value outside a function")
		}
		return e.enclosing.Assign(name, val)
	}

	_, ok := e.values[name]
	if ok {
		e.values[name] = val
		return nil
	}

	if !ok && e.enclosing != nil {
		fmt.Printf("Retracted to %p", e.enclosing)
		err := e.enclosing.Assign(name,val)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%s : Undefined variable", name)
}

func (e *Environment) GetValue(name string) any {
	if name == "ret_val" && !e.functionBound {
		if e == global{
			panic("Cannot access return value outside a function")
		}
		return e.enclosing.GetValue(name)
	}

	val, ok := e.values[name]
	if !ok {
		if e.enclosing != nil {
			fmt.Printf("Retracted to %p", e.enclosing)
			val = e.enclosing.GetValue(name)
			return val
		}
		panic("Undefined variable " + name)
	}
	return val
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
