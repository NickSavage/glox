package parser

import (
	"errors"
	"log"

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

func (p *Parser) consume(tokenType tokens.TokenType, message string) (tokens.Token, error) {
	next := p.Tokens[p.Current]

	if next.Type.Type == tokenType.Type {
		p.Current++
		return next, nil
	}
	return tokens.Token{}, errors.New(message)
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
		statement, err := p.Declaration()
		if err != nil {
			log.Printf("?")
			return statements, err
		}
		statements = append(statements, statement)

		if p.Tokens[p.Current].Type.Type == "EOF" {
			break
		}
	}
	return statements, nil
}
