package parser

import (
	"errors"
	"fmt"
	"log"

	"github.com/NickSavage/glox/src/tokens"
)

func (p *Parser) Expression() (*Expression, error) {
	return p.Assignment()
}

func (p *Parser) Assignment() (*Expression, error) {
	expr, err := p.Lambda()
	//	return expr, err

	if p.match(tokens.TokenType{Type: "Equal"}) {
		//		equals := p.Tokens[p.Current-1]
		value, err := p.Assignment()
		if err != nil {
			return &Expression{}, nil
		}
		return &Expression{
			Type:        "Assignment",
			Name:        expr.Name,
			AssignValue: value,
		}, nil

	}
	return expr, err
}

func (p *Parser) Lambda() (*Expression, error) {
	expr, err := p.Or()
	if p.match(tokens.TokenType{Type: "Lambda"}) {
		parameters := make([]tokens.Token, 0)
		for {
			token := p.Tokens[p.Current]
			if token.Type.Type != "Identifier" {
				return &Expression{}, errors.New("Expecting parameter name")
			}
			parameters = append(parameters, token)
			p.Current++
			if p.Tokens[p.Current].Type.Type == "Colon" {
				break
			}
			if !p.match(tokens.TokenType{Type: "Comma"}) {
				return &Expression{}, errors.New("Expecting , after parameter name")
			}

		}
		if !(p.match(tokens.TokenType{Type: "Colon"})) {
			return &Expression{}, errors.New("expecting ':' in lambda")
		}

		var statement *Statement
		var perr ParseError
		if p.match(tokens.TokenType{Type: "LeftBrace"}) {
			statement, perr = p.BlockStatement()
		} else {
			statement, perr = p.ReturnStatement()
		}
		if perr.HasError {
			return &Expression{}, perr.Message
		}
		statements := make([]*Statement, 0)
		statements = append(statements, statement)
		lambda := &Statement{
			Parameters: parameters,
			Statements: statements,
		}

		// backing up, without this the lambda needs two ;
		p.Current--

		return &Expression{
			Type:   "Lambda",
			Lambda: lambda,
		}, nil

	}
	return expr, err
}

func (p *Parser) Or() (*Expression, error) {
	expr, err := p.And()
	result := &Expression{}
	if p.match(tokens.TokenType{Type: "Or"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Type = "Logical"
		result.Left = expr
		result.Right, err = p.And()
		return result, nil
	}
	return expr, err

}

func (p *Parser) And() (*Expression, error) {
	expr, err := p.Equality()
	result := &Expression{}
	if p.match(tokens.TokenType{Type: "And"}) {
		result.Operator = p.Tokens[p.Current-1]
		result.Type = "Logical"
		result.Left = expr
		result.Right, err = p.Equality()
		return result, nil
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

	if p.match(tokens.TokenType{Type: "Plus"}) || p.match(tokens.TokenType{Type: "Minus"}) || p.match(tokens.TokenType{Type: "Tilde"}) {
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

	return p.Call()
}

func (p *Parser) Call() (*Expression, error) {
	expr, err := p.Primary()
	for {
		if p.match(tokens.TokenType{Type: "LeftParen"}) {
			expr, err = p.FinishCall(expr)
		} else {
			break
		}
	}

	return expr, err
}

func (p *Parser) FinishCall(expr *Expression) (*Expression, error) {
	arguments := make([]*Expression, 0)
	for {
		if !(p.Tokens[p.Current].Type.Type == "LeftParen") {
			if len(arguments) >= 255 {
				log.Printf("cannot have more than 255 arguments")
			}
			if p.Tokens[p.Current].Type.Type == "RightParen" {
				break
			}
			argument, err := p.Expression()
			if err != nil {
				return &Expression{}, err
			}
			arguments = append(arguments, argument)
			if !p.match(tokens.TokenType{Type: "Comma"}) {
				break
			}
		}
	}
	if !(p.match(tokens.TokenType{Type: "RightParen"})) {
		return &Expression{}, errors.New("expecting ) after function arguments")
	}
	result := &Expression{
		Type:          "Function",
		IsFunction:    true,
		Arguments:     arguments,
		FunctionParen: p.Tokens[p.Current-1],
		FunctionName:  expr.Name,
	}
	return result, nil
}

func LiteralExpression(token tokens.Token) *Expression {
	return &Expression{Value: token, Type: "Literal"}
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
	if p.match(tokens.TokenType{Type: "Identifier"}) {
		name := p.Tokens[p.Current-1]
		if p.match(tokens.TokenType{Type: "LeftBracket"}) {

			token := p.Tokens[p.Current]
			if token.Type.Type != "Number" {
				return &Expression{}, fmt.Errorf("arrays can only be indexed by numbers, not %v", token.Type.Type)
			}
			p.Current++
			if !p.match(tokens.TokenType{Type: "RightBracket"}) {
				return &Expression{}, errors.New("Expecting ] after array element")
			}
			return &Expression{
				Name:  name,
				Index: token.Literal.(int),
				Type:  "Element",
			}, nil

		}
		return &Expression{Name: name, Type: "Variable"}, nil
	}
	if p.match(tokens.TokenType{Type: "LeftBracket"}) {
		elements := make([]tokens.Token, 0)
		for {
			token := p.Tokens[p.Current]
			elements = append(elements, token)
			p.Current++
			if p.match(tokens.TokenType{Type: "RightBracket"}) {
				break
			}
			if !p.match(tokens.TokenType{Type: "Comma"}) {
				return &Expression{}, errors.New("Expecting , after array element")
			}
		}
		array := Array{
			Elements: elements,
			Length:   len(elements),
		}
		return &Expression{Type: "Array", Array: &array}, nil
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
