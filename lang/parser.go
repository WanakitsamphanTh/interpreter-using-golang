package lang

import "fmt"

type Parser struct {
	tokens  []Token
	current int
	env *Environment
}

func NewParser(tokens []Token, env *Environment) *Parser {
	return &Parser{
		tokens: tokens,
		current: 0,
		env: env,
	}
}

func (p *Parser) parse() ([]Statement, error) {
	var statements []Statement
	for !p.isAtEnd() {
		stmt, err := p.declaration()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}

	return statements, nil
}

func (p *Parser) declaration() (Statement, error) {
	if p.match(VAR) {
		return p.varDecl()
	}
	return p.statement()
}

func (p *Parser) expression() (Exp, error) {
	return p.assignment()
}

func (p *Parser) varDecl() (Statement, error) {
	name, err := p.consume(IDENTIFIER, "Expected variable name")
	var init Exp
	if p.match(EQUAL) {
		init, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	_, err = p.consume(SEMICOLON, "Expected ; after variable declaration.")
	return &Var{name, init, p.env}, err
}

func (p *Parser) statement() (Statement, error) {
	if p.match(PRINT) {
		return p.newPrintStatement()
	}
	if p.match(LEFT_BRACE) {
		return p.newBlock()
	}
	return p.newExpressionStatement()
}

func (p *Parser) newBlock() (Statement, error) {
	var statements []Statement
	prev := p.env
	current_env := NewNestedEnvironment(prev)
	p.env = &current_env
	for !p.check(RIGHT_BRACE) && !p.isAtEnd() {
		statement, err :=  p.declaration()
		if err != nil {
			p.env = prev
			return nil, err
		}
		statements = append(statements, statement)
	}
	_, err := p.consume(RIGHT_BRACE, "Expect '}' after block.");
	p.env = prev
	return NewBlock(statements, &current_env), err
}

func (p *Parser) newExpressionStatement() (Statement, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expected ; after expression.")
	return NewExpressionStatement(expr), err
}

func (p *Parser) newPrintStatement() (Statement, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expected ; after value.")
	return NewPrintStatement(expr), err
}

func (p *Parser) equality() (Exp, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(BANG_EQUAL, EQUAL_EQUAL){
		op := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExp{expr, op, right}
	}

	return expr, nil
}

func (p *Parser) assignment() (Exp, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	if p.match(EQUAL) {
		equals := p.previous()
		val, err := p.assignment()
		if err != nil {
			return nil, err
		}
		var_expr, ok := expr.(*Variable)
		if !ok {
			return nil, raiseError(equals, "Invalid assignment target")
		}
		return &Assignment{var_expr.name, val, p.env}, nil
	}
	return expr, nil
}

func (p *Parser) comparison() (Exp, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		op := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExp{expr, op, right}
	}

	return expr, nil
}

func (p *Parser) term() (Exp, error){
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(PLUS, MINUS) {
		op := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExp{expr, op, right}
	}
	return expr, nil
}

func (p *Parser) factor() (Exp, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(STAR, SLASH) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &BinaryExp{expr, op, right}
	}

	return expr, nil
}

func (p *Parser) unary() (Exp, error) {
	if p.match(BANG,MINUS) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &UnaryExp{op, right}, nil
	}
	return p.primary()
}

func (p *Parser) primary() (Exp, error) {
	if p.match(FALSE) {
		return &LiteralExp{false}, nil
	}
	if p.match(TRUE) {
		return &LiteralExp{true}, nil
	}
	if p.match(NIL){
		return &LiteralExp{nil}, nil
	}
	if p.match(NUMBER,STRING) {
		return &LiteralExp{p.previous().Literal}, nil
	}
	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return &GroupingExp{expr}, nil
	}
	if p.match(IDENTIFIER) {
		return &Variable{p.previous(), p.env}, nil
	}
	return nil, raiseError(p.peek(), "Expect expression")
}

func (p *Parser) match(tokenType ...Keyword) bool {
	for _, t := range tokenType {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType Keyword) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current - 1]
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Type == EOF
}

func (p *Parser) consume(tokenType Keyword, msg string) (Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return p.peek(), raiseError(p.peek(), msg) // ?
}


func raiseError(token Token, msg string) error {
	msg = fmt.Sprintf("%v : %v", token.Lexeme, msg)
	return &SyntaxError{token.Line, msg}	
}
