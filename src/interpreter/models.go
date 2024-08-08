package interpreter

import (
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

type Interpreter struct {
	Expression     *parser.Expression
	Memory         *Storage
	InLoop         bool
	BreakTriggered bool
}

type RuntimeError struct {
	Message  error
	HasError bool
	Token    tokens.Token
}

type Storage struct {
	Memory       map[string]interface{}
	Enclosing    *Storage
	HasEnclosing bool
}
