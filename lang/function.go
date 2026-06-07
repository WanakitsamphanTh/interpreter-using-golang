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
	NewNestedEnvironment()
	defer RetractEnvironment()
	for i, param := range f.decl.params {
		current_env.Define(param, params[i])
	}
	err := f.decl.body.Execute()
	return nil, err
}

// FnDecl struct

type FnDecl struct {
	name Token
	params []Token
	body *Block
}

func (fn *FnDecl) Execute() error {
	current_env.Define(fn.name, &Function{fn})
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
