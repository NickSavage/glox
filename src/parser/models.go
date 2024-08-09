package parser

import "github.com/NickSavage/glox/src/tokens"

type Statement struct {
	Type           tokens.TokenType
	Expression     *Expression
	VariableName   tokens.Token
	FunctionName   tokens.Token
	Initializer    *Expression
	Statements     []*Statement
	IsBlock        bool
	Condition      *Expression
	ElseStatements []*Statement
	Parameters     []tokens.Token
}

type Expression struct {
	Expression    *Expression
	Left          *Expression
	Operator      tokens.Token
	Right         *Expression
	Value         tokens.Token
	Name          tokens.Token
	AssignValue   *Expression
	Type          string // "binary, unary, literal, grouping, assignment, identifier, function"
	IsFunction    bool
	Arguments     []*Expression
	FunctionName  tokens.Token
	FunctionParen tokens.Token
	Lambda        *Statement
}

type Parser struct {
	Tokens  []tokens.Token
	Current int
}
