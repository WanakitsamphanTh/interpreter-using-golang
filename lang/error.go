package lang

import "fmt"

type RuntimeError struct {
	token Token
	error string
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("at line %v, %v", e.token.Line, e.error)
}

type SyntaxError struct {
	line int16
	error string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("at line %v, %v", e.line, e.error)
}