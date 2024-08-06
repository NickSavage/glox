package parser

import "github.com/NickSavage/glox/src/tokens"

type Statement struct {
	Type         tokens.TokenType
	Expression   *Expression
	VariableName tokens.Token
	Initializer  *Expression
}

type Expression struct {
	Expression *Expression
	Left       *Expression
	Operator   tokens.Token
	Right      *Expression
	Value      tokens.Token
	Type       string // "binary, unary, literal, grouping"
}

type Parser struct {
	Tokens  []tokens.Token
	Current int
}
