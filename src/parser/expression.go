package parser

import (
	"errors"
	"github.com/NickSavage/glox/src/tokens"
	"log"
)

func LiteralExpression(token tokens.Token) *Expression {
	return &Expression{Value: token, Type: "Literal"}
}

func (p *Parser) Expression() (*Expression, error) {
	return p.Assignment()
}

func (p *Parser) Assignment() (*Expression, error) {
	expr, err := p.Equality()
	//	return expr, err

	if p.match(tokens.TokenType{Type: "Equal"}) {
		//		equals := p.Tokens[p.Current-1]
		value, err := p.Assignment()
		log.Printf("value %v", value)
		if err != nil {
			return &Expression{}, nil
		}
		return &Expression{
			Type:        "Assignment",
			AssignValue: value,
		}, nil

	}
	return expr, err
}

func (p *Parser) Equality() (*Expression, error) {
	var err error
	result := &Expression{}
	left, err := p.Comparison()
	if err != nil {
		return &Expression{}, err
	}

	if p.match(tokens.TokenType{Type: "EqualEqual"}) ||
		p.match(tokens.TokenType{Type: "BangEqual"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Right, err = p.Comparison()
		if err != nil {
			return &Expression{}, err
		}
		result.Type = "Binary"
		result.Left = left

		return result, err

	}
	return left, err

}

func (p *Parser) Comparison() (*Expression, error) {

	var err error
	result := &Expression{}
	left, err := p.Term()
	if err != nil {
		return &Expression{}, err
	}

	if p.match(tokens.TokenType{Type: "Greater"}) ||
		p.match(tokens.TokenType{Type: "Less"}) ||
		p.match(tokens.TokenType{Type: "GreaterEqual"}) ||
		p.match(tokens.TokenType{Type: "LessEqual"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Right, err = p.Term()
		if err != nil {
			return &Expression{}, err
		}
		result.Type = "Binary"
		result.Left = left

		return result, err

	}
	return left, err

}

func (p *Parser) Term() (*Expression, error) {

	var err error
	result := &Expression{}
	left, err := p.Factor()
	if err != nil {
		return &Expression{}, err
	}

	if p.match(tokens.TokenType{Type: "Plus"}) || p.match(tokens.TokenType{Type: "Minus"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Right, err = p.Factor()
		if err != nil {
			return &Expression{}, err
		}
		result.Type = "Binary"
		result.Left = left

		return result, err

	}
	return left, err

}

func (p *Parser) Factor() (*Expression, error) {
	var err error
	result := &Expression{}
	left, err := p.Unary()
	if err != nil {
		return &Expression{}, err
	}

	if p.match(tokens.TokenType{Type: "Star"}) || p.match(tokens.TokenType{Type: "Slash"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Right, err = p.Unary()
		if err != nil {
			return &Expression{}, err
		}
		result.Type = "Binary"
		result.Left = left

		return result, err

	}
	return left, err
}

func (p *Parser) Unary() (*Expression, error) {
	var err error
	result := &Expression{}
	if p.match(tokens.TokenType{Type: "Bang"}) || p.match(tokens.TokenType{Type: "Minus"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Right, err = p.Unary()
		result.Type = "Unary"
		//log.Printf("%v", result.Right)
		if err != nil {
			return &Expression{}, err
		}
		return result, nil
	}

	return p.Primary()
}

func (p *Parser) Primary() (*Expression, error) {
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

	// TODO: add expression()
	if p.match(tokens.TokenType{Type: "LeftParen"}) {
		expr, err := p.Expression()
		if err != nil {
			return &Expression{}, nil
		}
		if !(p.match(tokens.TokenType{Type: "RightParen"})) {
			return &Expression{}, errors.New("expecting ')' after expression")
		}
		return &Expression{
			Expression: expr,
			Type:       "Grouping",
		}, nil
	}

	return &Expression{}, errors.New("not implemented yet")

}
