package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/NickSavage/glox/src/interpreter"
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

var Memory *interpreter.Storage

func report(line int, where string, message string) {
	log.Printf("Error: line %v, token %v: %v", line, where, message)
}

func printError(line int, message string) {
	report(line, "", message)
}

func run(source string) error {
	s := tokens.Scanner{
		Source: source,
		Tokens: make([]tokens.Token, 0),
	}
	err := s.ScanTokens()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	p := parser.Parser{
		Tokens:  s.Tokens,
		Current: 0,
	}
	declarations, err := p.Parse()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	i := interpreter.Interpreter{
		Memory: Memory,
	}
	for _, declaration := range declarations {

		rerr := i.Execute(declaration)
		if rerr.HasError {
			return rerr.Message
		}

	}

	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	Memory = &interpreter.Storage{
		Memory: make(map[string]interface{}),
	}
	var err error
	for {
		print("> ")
		if !scanner.Scan() {
			// If there's an error, print it and exit the program
			fmt.Println(scanner.Err())
			return
		}
		input := scanner.Text()

		switch strings.ToLower(input) {
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			err = run(input)
			if err != nil {
				printError(0, err.Error())
			}

		}
	}

}

func runFile(path string) error {
	Memory = &interpreter.Storage{
		Memory: make(map[string]interface{}),
	}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = run(string(contents))
	if err != nil {
		return err
	}
	return nil

}

// complete main by taking in command line arguments and returning
func main() {
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else if len(os.Args) < 2 {
		runPrompt()
	} else {
		print("Usage: glox [script]\n")
		return
	}

}
