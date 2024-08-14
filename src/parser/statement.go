package parser

import (
	"errors"
	//	"fmt"
	"log"

	"github.com/NickSavage/glox/src/tokens"
)

func (p *Parser) Statement() (*Statement, ParseError) {
	if p.match(tokens.TokenType{Type: "If"}) {
		return p.IfStatement()
	}
	if p.match(tokens.TokenType{Type: "Break"}) {
		return p.BreakStatement()
	}
	if p.match(tokens.TokenType{Type: "Continue"}) {
		return p.ContinueStatement()
	}
	if p.match(tokens.TokenType{Type: "For"}) {
		return p.ForStatement()
	}
	if p.match(tokens.TokenType{Type: "Function"}) {
		return p.FunctionStatement()
	}
	if p.match(tokens.TokenType{Type: "LeftBrace"}) {
		return p.BlockStatement()
	}
	if p.match(tokens.TokenType{Type: "Return"}) {
		return p.ReturnStatement()
	}
	return p.ExpressionStatement()
}

func (p *Parser) BreakStatement() (*Statement, ParseError) {
	return &Statement{
		Type: tokens.TokenType{Type: "Break"},
	}, ParseError{}
}

func (p *Parser) ContinueStatement() (*Statement, ParseError) {
	return &Statement{
		Type: tokens.TokenType{Type: "Continue"},
	}, ParseError{}

}

func (p *Parser) IfStatement() (*Statement, ParseError) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "If"},
	}

	condition, _ := p.Equality()
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return statement, ParseError{
			HasError: true,
			Message:  errors.New("expecting '{' after condition"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	block, perr := p.BlockStatement()
	if perr.HasError {
		return statement, perr
	}
	statement.Condition = condition
	statement.Statements = block.Statements

	if p.match(tokens.TokenType{Type: "Else"}) {
		if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
			return statement, ParseError{
				HasError: true,
				Message:  errors.New("expecting '{' after else"),
				Token:    p.Tokens[p.Current-1],
			}
		}

		elseBlock, perr := p.BlockStatement()
		if perr.HasError {
			return statement, perr
		}
		statement.ElseStatements = elseBlock.Statements
	}

	return statement, ParseError{}
}

func (p *Parser) ForStatement() (*Statement, ParseError) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "For"},
	}
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return statement, ParseError{
			HasError: true,
			Message:  errors.New("expecting '{' after for"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	block, perr := p.BlockStatement()
	log.Printf("block %v", block)
	if perr.HasError {
		return statement, perr
	}
	statement.Statements = block.Statements
	return statement, ParseError{}
}

func (p *Parser) BlockStatement() (*Statement, ParseError) {
	statement := &Statement{}
	statement.Type = tokens.TokenType{Type: "Block"}
	statements := make([]*Statement, 0)

	for {
		statement, perr := p.Declaration()
		if perr.HasError {
			return statement, perr
		}
		statements = append(statements, statement)
		next := p.Tokens[p.Current]
		if next.Type.Type == "EOF" || next.Type.Type == "RightBrace" {
			break
		}
	}

	if !(p.match(tokens.TokenType{Type: "RightBrace"})) {
		return statement, ParseError{
			HasError: true,
			Message:  errors.New("expecting '}' after block"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	statement.IsBlock = true
	statement.Statements = statements
	return statement, ParseError{}
}

func (p *Parser) ExpressionStatement() (*Statement, ParseError) {
	expr, err := p.Expression()
	if err != nil {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  err,
			Token:    p.Tokens[p.Current-1],
		}
	}
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		log.Printf("expr %v", expr)
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting ';' after expression"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	return &Statement{
		Type:       tokens.TokenType{Type: "Expression"},
		Expression: expr,
	}, ParseError{}
}

func (p *Parser) ReturnStatement() (*Statement, ParseError) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "Return"},
	}
	if p.Tokens[p.Current].Type.Type != "Semicolon" {
		expr, err := p.Expression()
		if err != nil {
			return &Statement{}, ParseError{
				HasError: true,
				Message:  err,
				Token:    p.Tokens[p.Current],
			}
		}
		statement.Expression = expr

	}
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting ';' after expression"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	return statement, ParseError{}
}

func (p *Parser) FunctionStatement() (*Statement, ParseError) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "Function"},
	}
	name := p.Tokens[p.Current]
	if name.Type.Type != "Identifier" {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting function name"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	p.Current++
	statement.FunctionName = name

	if !(p.match(tokens.TokenType{Type: "LeftParen"})) {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting '(' after parameters"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	parameters := make([]tokens.Token, 0)
	for {
		token := p.Tokens[p.Current]
		if token.Type == (tokens.TokenType{Type: "RightParen"}) {
			break
		}
		if token.Type.Type != "Identifier" {
			return statement, ParseError{
				HasError: true,
				Message:  errors.New("Expecting parameter name"),
				Token:    token,
			}

		}
		parameters = append(parameters, token)
		p.Current++
		if !p.match(tokens.TokenType{Type: "Comma"}) {
			break
		}

	}
	statement.Parameters = parameters

	if !(p.match(tokens.TokenType{Type: "RightParen"})) {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting ')' after parameters"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return &Statement{}, ParseError{
			HasError: true,
			Message:  errors.New("expecting '{' after parameters"),
			Token:    p.Tokens[p.Current-1],
		}
	}
	body, perr := p.BlockStatement()
	if perr.HasError {
		return statement, perr
	}
	statement.Statements = body.Statements
	return statement, ParseError{}

}
