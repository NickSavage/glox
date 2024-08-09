package interpreter

import (
	"fmt"
	//	"log"
	"os"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func identifierToken(name string) tokens.Token {
	return tokens.Token{
		Type:    tokens.TokenType{Type: "Identifier"},
		Lexeme:  name,
		Literal: nil,
		Line:    0,
	}
}

func (i *Interpreter) PrintFunction() *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("print"),
		Parameters: []tokens.Token{
			identifierToken("arg"),
		},
		NativeFunction: i.printImplementation,
	}
	return result
}

func (i *Interpreter) printImplementation() (interface{}, error) {
	arg, _ := i.Memory.Get("arg")
	//	print(arg)
	fmt.Println(arg)
	//log.Printf("")
	return nil, nil
}

func (i *Interpreter) ExitFunction() *parser.Statement {
	result := &parser.Statement{
		Type:           tokens.TokenType{Type: "NativeFunction"},
		FunctionName:   identifierToken("exit"),
		Parameters:     []tokens.Token{},
		NativeFunction: i.exitImplementation,
	}
	return result

}

func (i *Interpreter) exitImplementation() (interface{}, error) {
	os.Exit(0)
	return nil, nil
}
