package parser

import (
	"log"
	"testing"

	"github.com/NickSavage/glox/src/tokens"
)

func TestPrettyPrintExpressionTree(t *testing.T) {
	expectedOutput := "(* (- 123) (group 45.67))"
	input := &Expression{
		Left: &Expression{
			Operator: tokens.MinusToken(0),
			Right: &Expression{
				Value: tokens.Token{
					Type:    tokens.TokenType{Type: "Number"},
					Lexeme:  "123",
					Literal: 123,
				},
				Type: "Literal",
			},
			Type: "Unary",
		},
		Operator: tokens.StarToken(0),
		Right: &Expression{
			Expression: &Expression{
				Value: tokens.Token{
					Type:    tokens.TokenType{Type: "Number"},
					Lexeme:  "45.67",
					Literal: 45.67,
				},
				Type: "Literal",
			},
			Type: "Grouping",
		},
		Type: "Binary",
	}
	output := prettyPrintExpressionTree(input, "")
	if expectedOutput != output {
		log.Fatalf("output of prettyPrintExpressionTree not correct, got %v want %v", output, expectedOutput)
	}
}

func makeParser(text string) (Parser, error) {
	s := tokens.Scanner{
		Source: text,
		Tokens: make([]tokens.Token, 0),
	}
	err := s.ScanTokens()
	if err != nil {
		return Parser{}, err
	}

	return Parser{
		Tokens:  s.Tokens,
		Current: 0,
	}, nil
}

// func TestParserEquality(t *testing.T) {
// 	p, err := makeParser("1 == 1")
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	expr, err := p.Equality()
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	if expr.Type != "Binary" {
// 		t.Errorf("unexpected expression, got %v want Binary", expr.Type)
// 	}

// }

func TestParserPrimaryFalse(t *testing.T) {

	p, err := makeParser("false")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Primary()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Literal" {
		t.Errorf("unexpected expression, got %v want Literal", expr.Type)
	}
	if !(expr.Value.Lexeme == "false") {
		t.Errorf("unexpected expression, got %v want false", expr.Value.Lexeme)
	}

}

func TestParserPrimaryNumber(t *testing.T) {
	p, err := makeParser("1")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Primary()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Literal" {
		t.Errorf("unexpected expression, got %v want Literal", expr.Type)
	}
	if !(expr.Value.Lexeme == "1") {
		t.Errorf("unexpected expression, got %v want 1", expr.Value.Lexeme)
	}
	if !(expr.Value.Type.Type == "Number") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
}

func TestParserPrimaryNumberString(t *testing.T) {
	p, err := makeParser("1 'hello world'")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Primary()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Literal" {
		t.Errorf("unexpected expression, got %v want Literal", expr.Type)
	}
	if !(expr.Value.Lexeme == "1") {
		t.Errorf("unexpected expression, got %v want 1", expr.Value.Lexeme)
	}
	if !(expr.Value.Type.Type == "Number") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}

	expr, err = p.Primary()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Literal" {
		t.Errorf("unexpected expression, got %v want Literal", expr.Type)
	}
	if !(expr.Value.Lexeme == "'hello world'") {
		t.Errorf("unexpected expression, got %v want 1", expr.Value.Lexeme)
	}
	if !(expr.Value.Type.Type == "String") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
}
