package parser

import "github.com/NickSavage/glox/src/tokens"

type Statement struct {
	Type         tokens.TokenType
	Expression   *Expression
	VariableName tokens.Token
	Initializer  *Expression
}

type Expression struct {
	Expression  *Expression
	Left        *Expression
	Operator    tokens.Token
	Right       *Expression
	Value       tokens.Token
	Name        tokens.Token
	AssignValue *Expression
	Type        string // "binary, unary, literal, grouping, assignment, identifier"
}

type Parser struct {
	Tokens  []tokens.Token
	Current int
}
