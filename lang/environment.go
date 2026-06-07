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
		e.Define("@ret_val", nil)
	}
	e.Define("@terminated", false)

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
		err := e.enclosing.Assign(name,val)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%s : Undefined variable", name)
}

func (e *Environment) GetValue(name string) any {
	if name == "@ret_val" && !e.functionBound {
		if e == global{
			panic("Cannot access return value outside a function")
		}
		return e.enclosing.GetValue(name)
	}

	if name == "@terminated"  {
		if e.functionBound {
			return e.values["@terminated"]
		} 
		isLoop, ok := e.values["@looping"].(bool)
		if ok {
			if isLoop {
				return e.values["@terminated"]
			}
		}
		if e == global {
			return e.values["@terminated"]
		}
		return e.enclosing.GetValue(name)
	}

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

func (e *Environment) TerminateLoop() error {
	
	_, ok := e.values["@looping"]
	if !ok {
		return e.enclosing.TerminateLoop()
	}

	e.Assign("@terminated", true)
	
	return nil
}

func (e *Environment) TerminateFunction(val any) error {
	if !e.functionBound {
		if e == global{
			panic("Cannot termination outside a function")
		}
		e.values["@terminated"] = true
		return e.enclosing.TerminateFunction(val)
	}
	e.values["@terminated"] = true
	e.values["@ret_val"] = val
	return nil
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
