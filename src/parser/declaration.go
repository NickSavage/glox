package parser

import "github.com/NickSavage/glox/src/tokens"

func (p *Parser) Declaration() (*Statement, error) {
	if p.match(tokens.TokenType{Type: "Var"}) {
		return p.varDeclaration()
	}
	// TODO we need to be handling syncronizing here
	return p.Statement()

}

func (p *Parser) varDeclaration() (*Statement, error) {
	var statement Statement

	statement.Type = tokens.TokenType{Type: "Variable"}
	var init *Expression
	name, err := p.consume(tokens.TokenType{Type: "Identifier"}, "expecting variable name.")
	if err != nil {
		return &statement, err
	}
	statement.VariableName = name
	if p.match(tokens.TokenType{Type: "Equal"}) {
		p.Current--
		init, err = p.Expression()
		if err != nil {
			return &statement, err
		}
		statement.Initializer = init
	}

	_, err = p.consume(tokens.TokenType{Type: "Semicolon"}, "expecting ; after variable declaration.")
	if err != nil {
		return &statement, err
	}
	return &statement, nil

}
