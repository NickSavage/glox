package parser

import (
	"errors"

	"github.com/NickSavage/glox/src/tokens"
)

func prettyPrintExpressionTree(input *Expression, result string) string {
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
		result = prettyPrintExpressionTree(input.Expression, result)
	}
	if input.Left != nil && input.Left.Type != "" {
		result = prettyPrintExpressionTree(input.Left, result) + " "
	}
	if input.Right != nil && input.Right.Type != "" {
		result = prettyPrintExpressionTree(input.Right, result)
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

// func (p *Parser) Equality() {
// 	var expr Expression

// 	if p.match(tokens.EqualEqualToken().Type) {

// 	}

// }

func LiteralExpression(token tokens.Token) Expression {
	return Expression{Value: token, Type: "Literal"}
}

func (p *Parser) Primary() (Expression, error) {
	if p.match(tokens.TokenType{Type: "False"}) {
		return LiteralExpression(p.Tokens[p.Current-1]), nil
	}
	if p.match(tokens.TokenType{Type: "True"}) {
		return LiteralExpression(p.Tokens[p.Current-1]), nil
	}
	if p.match(tokens.TokenType{Type: "Nil"}) {
		return LiteralExpression(p.Tokens[p.Current-1]), nil
	}
	if p.match(tokens.TokenType{Type: "Number"}) {
		return LiteralExpression(p.Tokens[p.Current-1]), nil
	}
	if p.match(tokens.TokenType{Type: "String"}) {
		return LiteralExpression(p.Tokens[p.Current-1]), nil
	}

	return Expression{}, errors.New("not implemented yet")

}
