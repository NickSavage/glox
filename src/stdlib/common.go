package stdlib

import (
	"errors"
	"github.com/NickSavage/glox/src/tokens"
)

func implWrapperError() error {
	return errors.New("Wrong type given to native function")
}

func identifierToken(name string) tokens.Token {
	return tokens.Token{
		Type:    tokens.TokenType{Type: "Identifier"},
		Lexeme:  name,
		Literal: nil,
		Line:    0,
	}
}
