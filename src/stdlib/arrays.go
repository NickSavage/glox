package stdlib

import (
	//	"errors"
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

func MapFunction(i *interpreter.Interpreter) *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("len"),
		Parameters: []tokens.Token{
			identifierToken("array"),
			identifierToken("func"),
		},
		NativeFunction: mapImplementation,
	}
	return result

}
func mapImplementation(i interface{}) (interface{}, error) {
	in, ok := i.(*interpreter.Interpreter)

	if !ok {
		return nil, implWrapperError()
	}
	a, _ := in.Memory.Get("array")
	array, ok := a.(*parser.Array)
	if !ok {
		log.Printf("error")
		return nil, fmt.Errorf("value is not an array")
	}
	f, _ := in.Memory.Get("func")
	function, ok := f.(*parser.Statement)
	if !ok {
		log.Printf("error")
		return nil, fmt.Errorf("function needs to be a function")
	}

	var result []tokens.Token
	var elementResult interface{}
	var rerr interpreter.RuntimeError

	log.Printf("asdsa %v", function.Expression)
	var arguments []interface{}
	for _, e := range array.Elements {
		// need to evaluate here
		arguments = make([]interface{}, 0)
		arguments = append(arguments, e)

		elementResult, rerr = in.FunctionCall(function.Expression, arguments)
		if rerr.HasError {
			log.Printf("map rerr %v", rerr.Message)
			return nil, rerr.Message
		}
		result = append(result, tokens.Token{
			Type:    e.Type,
			Lexeme:  fmt.Sprintf("%v", elementResult),
			Line:    e.Line,
			Literal: elementResult,
		})

	}

	return array, nil
}
