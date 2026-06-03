package lang

import "fmt"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		current: 0,
	}
}

func (p *Parser) parse() (Exp, error) {
	return p.expression()
}

func (p *Parser) expression() (Exp, error) {
	return p.equality()
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
	if token.Type == EOF {
		return fmt.Errorf("%v at end : %v", token.Line, msg)
	} 
	return fmt.Errorf("%v at %v : %v", token.Line, token.Lexeme, msg)
}