package parser

import (
	//	"log"

	"github.com/NickSavage/glox/src/tokens"
)

func PrettyPrintExpressionTree(input *Expression, result string) string {
	if input.Value.Lexeme != "" {
		result += input.Value.Lexeme
		return result
	}

	result += "("
	if input.Operator.Lexeme != "" {
		result += input.Operator.Lexeme + " "
	} else if input.Type == "Grouping" {
		result += "group" + " "
	}
	if input.Expression != nil {
		result = PrettyPrintExpressionTree(input.Expression, result)
	}
	if input.Left != nil && input.Left.Type != "" {
		result = PrettyPrintExpressionTree(input.Left, result) + " "
	}
	if input.Right != nil && input.Right.Type != "" {
		result = PrettyPrintExpressionTree(input.Right, result)
	}
	result += ")"
	return result
}

func (p *Parser) match(tokenType tokens.TokenType) bool {
	result := p.Tokens[p.Current].Type.Type == tokenType.Type
	if result {
		p.Current++
	}
	return result

}

func (p *Parser) Parse() ([]*Statement, error) {
	statements := make([]*Statement, 0)
	for {
		statement, err := p.Statement()
		if err != nil {
			return statements, err
		}
		statements = append(statements, statement)

		if p.Tokens[p.Current].Type.Type == "EOF" {
			break
		}
	}
	return statements, nil
}
