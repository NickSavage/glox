package interpreter

import (
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

type Interpreter struct {
	Expression *parser.Expression
}

type RuntimeError struct {
	Message  error
	HasError bool
	Token    tokens.Token
}
