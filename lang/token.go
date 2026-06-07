package lang

import (
	"fmt"
)

type Keyword int

const (
	// Single-character tokens.
	LEFT_PAREN Keyword = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	SHARP
	
	// Two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	BREAK

	// EOF
	EOF
)

type Token struct {
	Type    Keyword
	Lexeme  string
	Literal any
	Line int16
}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}

func NewToken(Type Keyword, Lexeme string, Literal any, line int16) Token {
	return Token{Type, Lexeme, Literal, line}
}


func mapKeyword(lexeme string) Keyword {
	switch lexeme {
	case "and":
		return AND
	case "class":
		return CLASS
	case "else":
		return ELSE
	case "false":
		return FALSE
	case "for":
		return FOR
	case "fun":
		return FUN
	case "if":
		return IF
	case "nil":
		return NIL
	case "or":
		return OR
	case "print":
		return PRINT
	case "return":
		return RETURN
	case "super":
		return SUPER
	case "this":
		return THIS
	case "true":
		return TRUE
	case "var":
		return VAR
	case "while":
		return WHILE
	case "break":
		return BREAK
	default:
		return IDENTIFIER
	}
}