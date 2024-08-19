package stdlib

import (
	"fmt"
	"github.com/NickSavage/glox/src/interpreter"
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
	"log"
)

func LenFunction(i *interpreter.Interpreter) *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("len"),
		Parameters: []tokens.Token{
			identifierToken("array"),
		},
		NativeFunction: lenImplementation,
	}
	return result
}

func lenImplementation(i interface{}) (interface{}, error) {
	in, ok := i.(*interpreter.Interpreter)

	if !ok {
		return nil, implWrapperError()
	}
	array, err := in.Memory.Get("array")
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
