package stdlib

import (
	"errors"
	"fmt"
	"github.com/NickSavage/glox/src/interpreter"
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
	"log"
)

func identifierToken(name string) tokens.Token {
	return tokens.Token{
		Type:    tokens.TokenType{Type: "Identifier"},
		Lexeme:  name,
		Literal: nil,
		Line:    0,
	}
}

func LenFunction(i *interpreter.Interpreter) *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("len"),
		Parameters: []tokens.Token{
			identifierToken("array"),
		},
		NativeFunction: lenImplWrapper,
	}
	return result
}

func lenImplWrapper(i interface{}) (interface{}, error) {
	in, ok := i.(*interpreter.Interpreter)

	if !ok {
		return nil, errors.New("sdasd")
	}
	return lenImplementation(in)

}

func lenImplementation(i *interpreter.Interpreter) (interface{}, error) {
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
