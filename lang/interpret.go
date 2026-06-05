package lang

import (
	"fmt"
	"os"
	"bufio"
)

func run(str string, env *Environment) error {
	scanner := NewScanner(str)
	tokens, err := scanner.scanTokens()
	if err != nil {
		return err
	}

	parser := NewParser(tokens, env)
	statements, err := parser.parse()
	if err != nil {
		return err
	}
	//ast := Parenthesize(statements)
	//fmt.Println("Abstract Syntax Tree:\n", ast)
	err = interpret(statements)
	if err != nil {
		return err
	}
	return nil
}

func interpret(statements []Statement) error {
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
	env := NewEnvironment()
	if err != nil {
		return err
	}
	err = run(string(data), &env)
	if err != nil {
		return err
	}
	return nil
}

func RunREPL() error {
	env := NewEnvironment()
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
		err := run(input, &env)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
	}
	return nil
}
