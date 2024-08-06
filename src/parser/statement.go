package parser

import (
	"errors"
	"fmt"

	"github.com/NickSavage/glox/src/tokens"
)

func (p *Parser) Statement() (*Statement, error) {
	if p.match(tokens.TokenType{Type: "Print"}) {
		return p.PrintStatement()
	}
	return &Statement{}, fmt.Errorf("not implemented yet")
}

func (p *Parser) PrintStatement() (*Statement, error) {
	expr, err := p.Expression()
	if !(p.match(tokens.TokenType{Type: "Semicolon"})) {
		return &Statement{}, errors.New("expecting ';' after expression")
	}
	return &Statement{
		Expression: expr,
	}, err
}
