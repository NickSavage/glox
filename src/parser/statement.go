package parser

import (
	"errors"
	//	"fmt"
	"log"

	"github.com/NickSavage/glox/src/tokens"
)

func (p *Parser) Statement() (*Statement, error) {
	if p.match(tokens.TokenType{Type: "Print"}) {
		return p.PrintStatement()
	}
	if p.match(tokens.TokenType{Type: "LeftBrace"}) {
		return p.BlockStatement()
	}
	return p.ExpressionStatement()
}

func (p *Parser) PrintStatement() (*Statement, error) {
	expr, err := p.Expression()
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		return &Statement{}, errors.New("expecting ';' after expression")
	}
	return &Statement{
		Type:       tokens.TokenType{Type: "Print"},
		Expression: expr,
	}, err
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
	log.Printf(" expr statement")
	expr, err := p.Expression()
	if err != nil {
		return &Statement{}, err
	}
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		return &Statement{}, errors.New("expecting ';' after expression")
	}
	return &Statement{
		Type:       tokens.TokenType{Type: "Expression"},
		Expression: expr,
	}, nil
}
