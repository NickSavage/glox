package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/NickSavage/glox/src/tokens"
)

func report(line int, where string, message string) {

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

	for _, token := range s.Tokens {
		fmt.Println(token)
	}
	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
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
