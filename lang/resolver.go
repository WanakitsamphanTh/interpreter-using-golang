package lang

import "fmt"

type Resolveable interface {
	Resolve() error
}

type ScopeStack struct {
	elements []map[string]bool
}

func (s *ScopeStack) size() int {
	return len(s.elements)
}

func (s *ScopeStack) get(i int) map[string]bool {
	return s.elements[i]
}

func (s *ScopeStack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *ScopeStack) push(scope map[string]bool) {
	s.elements = append(s.elements, scope)
}

func (s *ScopeStack) pop() map[string]bool {
	index := len(s.elements) - 1
	final := s.elements[index]
	s.elements = s.elements[:index]
	return final
}

func (s *ScopeStack) peek() (map[string]bool, error) {
	index := len(s.elements) - 1
	if index == -1 {
		return nil, fmt.Errorf("The scope stack is empty")
	}
	return s.elements[index], nil
}

var scopes ScopeStack = ScopeStack{}
var locals map[Exp]int = make(map[Exp]int)

func beginScope() {
	scopes.push(make(map[string]bool))
}

func endScope() {
	scopes.pop()
}

func declare(name string) error {
	if scopes.isEmpty() {
		return nil
	}
	scope, err := scopes.peek()
	if err != nil {
		return err
	}
	scope[name] = false
	return nil
}

func define(name string) error {
	if scopes.isEmpty() {
		return nil
	}
	scope, err := scopes.peek()
	if err != nil {
		return err
	}
	scope[name] = true
	return nil
}

func resolveLocal(expr Exp, name string) error {
	for i := scopes.size() - 1; i >= 0; i++ {
		_, ok := scopes.get(i)[name]
		if ok {
			return resolveDepth(expr, scopes.size()-1-i)
		}
	}
	return nil
}

func resolveDepth(expr Exp, depth int) error {
	locals[expr] = depth
	return nil
}

func (block *Block) Resolve() error {
	return nil
}

func (expr *UnaryExp) Resolve() error {
	return expr.Right.Resolve()
}

func (expr *BinaryExp) Resolve() error {
	err := expr.Left.Resolve()
	if err != nil {
		return err
	}
	err = expr.Right.Resolve()
	if err != nil {
		return err
	}
	return nil
}

func (expr *LiteralExp) Resolve() error {
	return nil
}

func (expr *GroupingExp) Resolve() error {
	return expr.Exp.Resolve()
}

func (expr *LogicalExpression) Resolve() error {
	err := expr.left.Resolve()
	if err != nil {
		return err
	}
	err = expr.right.Resolve()
	if err != nil {
		return err
	}
	return nil
}

func (v *Variable) Resolve() error {
	if !scopes.isEmpty() {
		scopes, _ := scopes.peek()
		if scopes[v.name.Lexeme] == false {
			return &SyntaxError{v.name.Line, "Can't read local variable in its own initializer."}
		}
	}
	resolveLocal(v, v.name.Lexeme)
	return nil
}

func LookUpVariable(name string, expr Exp) any {
	distance, ok := locals[expr]
	if !ok {
		return current_env.GetAt(distance, name)
	} else {
		return global.GetValue(name)
	}
}

func (v *Assignment) Resolve() error {
	err := v.val.Resolve()
	if err != nil {
		return err
	}
	err = v.Resolve()
	return err
}

func (e *ExpressionStatement) Resolve() error {
	return e.expr.Resolve()
}

func (f *FnCall) Resolve() error {
	err := f.callee.Resolve()
	if err != nil {
		return err
	}
	for _, arg := range f.params {
		err := arg.Resolve()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FnDecl) Resolve() error {
	declare(f.name.Lexeme)
	define(f.name.Lexeme)
	resolveFunction(f)
	return nil
}

// ??????????????
func resolveFunction(f *FnDecl) error {
	beginScope()
	defer endScope()
	for _, param := range f.params {
		err := declare(param.Lexeme)
		if err != nil {
			return err
		}
		err = define(param.Lexeme)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *IfStatement) Resolve() error {
	err := f.condition.Resolve()
	if err != nil {
		return err
	}
	err = f.thenBranch.Resolve()
	if err != nil {
		return err
	}
	if f.elseBranch != nil {
		err = f.elseBranch.Resolve()
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Var) Resolve() error {
	declare(v.name.Lexeme)
	if v.init != nil {
		err := v.init.Resolve()
		if err != nil {
			return err
		}
	}
	define(v.name.Lexeme)
	return nil
}

func (v *Return) Resolve() error {
	if v.val != nil {
		err := v.val.Resolve()
		return err
	}
	return nil
}

func (v *WhileStatement) Resolve() error {
	err := v.condition.Resolve()
	if err != nil {
		return err
	}
	err = v.body.Resolve()
	if err != nil {
		return err
	}
	return nil
}

func (v *SkipStmt) Resolve() error {
	return nil
}

func (v *BreakStmt) Resolve() error {
	return nil
}
