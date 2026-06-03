package lang

import (
	"fmt"
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int16
	current int16
	line    int16
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokens: []Token{},
		start: int16(0),
		current: int16(0),
		line: int16(1),
	}
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return nil, err
		}
	}
	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: 0})
	return s.tokens, nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= int16(len(s.source))
}

func (s *Scanner) scanToken() error {
	var err error
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '/':
		s.addToken(SLASH, nil)
	case '*':
		s.addToken(STAR, nil)
	case '#': {
		for s.peek() != '\n' && !s.isAtEnd() {
			s.advance()
		}
	}
	case '!':
		if s.matchNext('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.matchNext('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '<':
		if s.matchNext('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.matchNext('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case ' ', '\r', '\t':
		// Ignore whitespace.
	case '\n':
		s.line++
	case '"':
		err = s.stringLiteral()
	default:
		// parse numbers
		if isDigit(c) {
			err = s.parseNumber()
		} else if isAlpha(c) {
			err = s.identifier()
		} else {
			err = fmt.Errorf("Unexpected character: %c", c)
		}
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) addToken(t Keyword, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: t, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) matchNext(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.source[s.current + 1]
}

func (s *Scanner) stringLiteral() error {
	for s.peek() != '"' && !s.isAtEnd() {
      if s.peek() == '\n' {
		s.line++;
	  }
      s.advance();
    }
	if s.isAtEnd() {
		err := fmt.Errorf("Unterminated string.")
		return err
	}
	s.advance() // The closing ".
	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
	return nil
}

func (s *Scanner) parseNumber() error {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return fmt.Errorf("Invalid number: %s", s.source[s.start:s.current])
	}
	s.addToken(NUMBER, value)

	return nil
}

func (s *Scanner) identifier() error {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	lexeme := s.source[s.start:s.current]
	keyword := mapKeyword(lexeme)
	if keyword != IDENTIFIER {
		s.addToken(keyword, nil)
	} else {
		s.addToken(IDENTIFIER, lexeme)
	}
	return nil
}
