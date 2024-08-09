package parser

import (
	"errors"
	//	"fmt"
	"log"

	"github.com/NickSavage/glox/src/tokens"
)

func (p *Parser) Statement() (*Statement, error) {
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

func (p *Parser) BreakStatement() (*Statement, error) {
	return &Statement{
		Type: tokens.TokenType{Type: "Break"},
	}, nil
}

func (p *Parser) ContinueStatement() (*Statement, error) {
	return &Statement{
		Type: tokens.TokenType{Type: "Continue"},
	}, nil

}

func (p *Parser) IfStatement() (*Statement, error) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "If"},
	}

	condition, err := p.Equality()
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return statement, errors.New("expecting '{' after condition")
	}
	block, err := p.BlockStatement()
	if err != nil {
		return statement, err
	}
	statement.Condition = condition
	statement.Statements = block.Statements

	if p.match(tokens.TokenType{Type: "Else"}) {
		if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
			return statement, errors.New("expecting '{' after else")
		}

		elseBlock, err := p.BlockStatement()
		if err != nil {
			return statement, err
		}
		statement.ElseStatements = elseBlock.Statements
	}

	return statement, nil
}

func (p *Parser) ForStatement() (*Statement, error) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "For"},
	}
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return statement, errors.New("expecting '{' after for")
	}
	block, err := p.BlockStatement()
	log.Printf("block %v", block)
	if err != nil {
		return statement, err
	}
	statement.Statements = block.Statements
	return statement, nil
}

func (p *Parser) BlockStatement() (*Statement, error) {
	statement := &Statement{}
	statement.Type = tokens.TokenType{Type: "Block"}
	statements := make([]*Statement, 0)

	for {
		statement, err := p.Declaration()
		if err != nil {
			return statement, err
		}
		statements = append(statements, statement)
		next := p.Tokens[p.Current]
		if next.Type.Type == "EOF" || next.Type.Type == "RightBrace" {
			break
		}
	}

	if !(p.match(tokens.TokenType{Type: "RightBrace"})) {
		return statement, errors.New("expecting '}' after block")
	}
	statement.IsBlock = true
	statement.Statements = statements
	return statement, nil
}

func (p *Parser) ExpressionStatement() (*Statement, error) {
	expr, err := p.Expression()
	if err != nil {
		return &Statement{}, err
	}
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		log.Printf("expr %v", expr)
		return &Statement{}, errors.New("expecting ';' after expression")
	}
	return &Statement{
		Type:       tokens.TokenType{Type: "Expression"},
		Expression: expr,
	}, nil
}

func (p *Parser) ReturnStatement() (*Statement, error) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "Return"},
	}
	if p.Tokens[p.Current].Type.Type != "Semicolon" {
		expr, err := p.Expression()
		if err != nil {
			return statement, err
		}
		statement.Expression = expr

	}
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		return &Statement{}, errors.New("expecting ';' after expression")
	}
	return statement, nil
}

func (p *Parser) FunctionStatement() (*Statement, error) {
	statement := &Statement{
		Type: tokens.TokenType{Type: "Function"},
	}
	name := p.Tokens[p.Current]
	if name.Type.Type != "Identifier" {
		return statement, errors.New("expecting function name")
	}
	p.Current++
	statement.FunctionName = name

	if !(p.match(tokens.TokenType{Type: "LeftParen"})) {
		return &Statement{}, errors.New("expecting '(' after function name")
	}
	parameters := make([]tokens.Token, 0)
	for {
		token := p.Tokens[p.Current]
		if token.Type == (tokens.TokenType{Type: "RightParen"}) {
			break
		}
		if token.Type.Type != "Identifier" {
			return statement, errors.New("Expecting parameter name")
		}
		parameters = append(parameters, token)
		p.Current++
		if !p.match(tokens.TokenType{Type: "Comma"}) {
			break
		}

	}
	statement.Parameters = parameters

	if !(p.match(tokens.TokenType{Type: "RightParen"})) {
		return &Statement{}, errors.New("expecting ')' after parameters")
	}
	if !(p.match(tokens.TokenType{Type: "LeftBrace"})) {
		return &Statement{}, errors.New("expecting '{' before function body")
	}
	body, err := p.BlockStatement()
	if err != nil {
		return statement, err
	}
	statement.Statements = body.Statements
	return statement, nil

}
