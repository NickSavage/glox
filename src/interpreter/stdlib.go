package interpreter

import (
	"errors"
	"fmt"
	"os"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func PrintFunction(i *Interpreter) *parser.Statement {
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: tokens.IdentifierToken("print"),
		Parameters: []tokens.Token{
			tokens.IdentifierToken("arg"),
		},
		NativeFunction: printImplementation,
	}
	return result
}

func printImplementation(i interface{}) (interface{}, error) {

	in, ok := i.(*Interpreter)

	if !ok {
		return nil, errors.New("sdasd")
	}
	arg, _ := in.Memory.Get("arg")
	//	print(arg)
	fmt.Println(arg)
	//log.Printf("")
	return nil, nil
}

func ExitFunction(i *Interpreter) *parser.Statement {
	result := &parser.Statement{
		Type:           tokens.TokenType{Type: "NativeFunction"},
		FunctionName:   tokens.IdentifierToken("exit"),
		Parameters:     []tokens.Token{},
		NativeFunction: exitImplementation,
	}
	return result

}

func exitImplementation(i interface{}) (interface{}, error) {
	os.Exit(0)
	return nil, nil
}
