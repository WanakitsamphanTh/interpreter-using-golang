package lang

import "fmt"

type Callable interface {
	call(params []any) (any, disruptive)
	arity() int
}

// Function struct

type Function struct {
	decl    *FnDecl
	closure *Environment
}

func (f *Function) arity() int {
	return len(f.decl.params)
}

func (f *Function) call(params []any) (any, disruptive) {
	//NewNestedEnvironment(true)
	//defer RetractEnvironment()
	prev := current_env
	current_env = NewEnvironment(f.closure, true)
	defer func() {
		current_env = prev
	}()

	for i, param := range f.decl.params {
		current_env.Define(param.Lexeme, params[i])
	}

	err := f.decl.body.Execute()
	if err != nil {
		switch disruption := err.(type) {
		case *Return:
			if disruption.val != nil {
				return disruption.val.Eval(), nil
			}
			return nil, nil
		default:
			return nil, err
		}
	}

	return nil, nil
}

// FnDecl struct

type FnDecl struct {
	name   Token
	params []Token
	body   *Block
}

func (fn *FnDecl) Execute() disruptive {
	current_env.Define(fn.name.Lexeme, &Function{fn, current_env})
	return nil
}

// FnCall struct

type FnCall struct {
	callee Exp
	paren  Token // to report where error occurs
	params []Exp
}

func (fn *FnCall) Eval() any {
	callee := fn.callee.Eval()
	var params []any
	for _, p := range fn.params {
		params = append(params, p.Eval())
	}
	callable, ok := callee.(Callable)
	if !ok {
		msg := fmt.Sprintf("At line %v, this is not a function", fn.paren.Line)
		panic(msg)
	}
	if callable.arity() != -1 && len(params) != callable.arity() {
		msg := fmt.Sprintf("At line %v, expected %v parameters but got %v", fn.paren.Line, callable.arity, len(params))
		panic(msg)
	}
	ret_val, err := callable.call(params)
	if err != nil {
		err := err.(error)
		panic(err.Error())
	}
	return ret_val
}

// Return struct
type Return struct {
	keyword Token
	val     Exp
}

func (r *Return) Execute() disruptive {
	return r
}

// Native functions
type NativeFn struct {
	_arity int
	_fn    func([]any) (any, disruptive)
}

func (fn *NativeFn) arity() int {
	return fn._arity
}

func (fn *NativeFn) call(params []any) (any, disruptive) {
	return fn._fn(params)
}
