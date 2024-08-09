package stdlib

import (
	"fmt"
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

func PrintFunction() *parser.Statement {
	funcImpl := printImplementation
	result := &parser.Statement{
		Type:         tokens.TokenType{Type: "NativeFunction"},
		FunctionName: identifierToken("print"),
		Parameters: []tokens.Token{
			identifierToken("arg"),
		},
		NativeFunction: funcImpl,
	}
	return result
}

func printImplementation(arg ...interface{}) (interface{}, error) {
	fmt.Printf("%v", arg)
	return nil, nil
}
