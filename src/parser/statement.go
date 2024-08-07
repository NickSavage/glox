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
