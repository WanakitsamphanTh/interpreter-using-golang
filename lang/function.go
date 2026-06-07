package lang

import "fmt"

type Callable interface {
	call(params []any) (any, error)
	arity() int
}

// Function struct

type Function struct {
	decl *FnDecl
}

func (f *Function) arity() int {
	return len(f.decl.params)
}

func (f *Function) call(params []any) (any, error) {
	NewNestedEnvironment(true)
	defer RetractEnvironment()
	for i, param := range f.decl.params {
		current_env.Define(param.Lexeme, params[i])
	}
	f.decl.body.shared = true
	err := f.decl.body.Execute()
	if err != nil {
		return nil, err
	}
	return current_env.GetValue("ret_val"), nil
}

// FnDecl struct

type FnDecl struct {
	name Token
	params []Token
	body *Block
}

func (fn *FnDecl) Execute() error {
	current_env.Define(fn.name.Lexeme, &Function{fn})
	return nil	
}

// FnCall struct

type FnCall struct {
	callee Exp
	paren Token // to report where error occurs
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
	if len(params) != callable.arity() {
		msg := fmt.Sprintf("At line %v, expected %v parameters but got %v", fn.paren.Line, callable.arity, len(params))
		panic(msg)
	}
	ret_val, err := callable.call(params)
	if err != nil {
		panic(err.Error())
	}
	return ret_val
}

// Return struct
type Return struct {
	keyword Token
	val Exp
}

func (r *Return) Execute() error {
	current_env.Assign("ret_val", r.val.Eval())
	return nil
}

// Native functions
type NativeFn struct {
	_arity int
	_fn func([]any) (any, error)
}

func (fn *NativeFn) arity() int {
	return fn._arity
}

func (fn *NativeFn) call(params []any) (any, error) {
	return fn._fn(params)
}

