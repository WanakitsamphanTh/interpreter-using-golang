package lang

import (
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
	expr, err := parser.parse()
	if err != nil {
		return err
	}
	
	fmt.Printf("%v\n", Parenthesize(expr))
	fmt.Printf("%v\n", expr.Eval())
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
	for {
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
		if input == "" {
			break
		}
		err = run(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
	}
	return nil
}
