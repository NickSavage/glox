package interpreter

import (
//	"errors"
	"fmt"
	"log"
	"os"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) LoadNativeFunctions() {
	i.Memory.Define("exit", i.ExitFunction())
	i.Memory.Define("len", i.LenFunction())
	i.Memory.Define("print", i.PrintFunction())
}

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

func (i *Interpreter) LenFunction() *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("len"),
		Parameters: []tokens.Token{
			identifierToken("array"),
		},
		NativeFunction: i.lenImplementation,
	}
	return result
}

func (i *Interpreter) lenImplementation() (interface{}, error) {
	array, err := i.Memory.Get("array")
	log.Printf("input")
	if err != nil {
		return nil, err
	}
	a, ok := array.(*parser.Array)
	if !ok {
		log.Printf("error")
		return nil, fmt.Errorf("value is not an array")
	}
	log.Printf("a %v", a)
	return a.Length, nil
}
