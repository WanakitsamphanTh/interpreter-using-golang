package lang

import (
	"bufio"
	"fmt"
	"os"
)

func run(str string) error {
	scanner := NewScanner(str)
	tokens, err := scanner.scanTokens()
	if err != nil {
		return err
	}

	parser := NewParser(tokens)
	statements, err := parser.parse()
	if err != nil {
		return err
	}

	ast := Parenthesize(statements)
	fmt.Println("Abstract Syntax Tree:\n", ast)

	for _, stmt := range statements {
		err := stmt.Resolve()
		if err != nil {
			return err
		}
	}

	_err := interpret(statements)
	if _err != nil {
		return _err.(error)
	}
	return nil
}

func interpret(statements []Statement) disruptive {
	for _, statement := range statements {
		err := statement.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func RunScript(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = run(string(data))
	if err != nil {
		return err
	}
	return nil
}

func RunREPL() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		var input string

		scanner.Scan()
		input = scanner.Text()

		if input == "" {
			break
		}

		fmt.Println("Input: ", input)
		err := run(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
	}
	return nil
}
