package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func report(line int, where string, message string) {

}

func error(line int, message string) {
	report(line, "", message)
}

func run(source string) {
	tokens := strings.Split(source, " ")
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
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
			run(input)
		}
	}

}

func runFile(path string) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	run(string(contents))

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
