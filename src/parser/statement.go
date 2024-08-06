package parser

import (
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
	return &Statement{
		Expression: expr,
	}, err
}
