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
	output := PrettyPrintExpressionTree(input, "")
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

func TestParserUnary(t *testing.T) {
	p, err := makeParser("!1 -123")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Unary()
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Printf("expr %v", expr)
	if expr.Type != "Unary" {
		t.Errorf("unexpected expression, got %v want Unary", expr.Type)
	}
	if !(expr.Operator.Type.Type == "Bang") {
		t.Errorf("unexpected operator, got %v want Bang", expr.Operator.Type.Type)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Right.Value.Literal == 1) {
		t.Errorf("unexpected expression, got %v want 1", expr.Right.Value.Literal)
	}

	expr, err = p.Unary()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Unary" {
		t.Errorf("unexpected expression, got %v want Unary", expr.Type)
	}
	if !(expr.Operator.Type.Type == "Minus") {
		t.Errorf("unexpected operator, got %v want Minus", expr.Operator.Type.Type)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Right.Value.Literal == 123) {
		t.Errorf("unexpected expression, got %v want 123", expr.Right.Value.Literal)
	}
}
func TestParserFactor(t *testing.T) {
	p, err := makeParser("1 / 2")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Factor()
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Printf("expr %v", expr)
	if expr.Type != "Binary" {
		t.Errorf("unexpected expression, got %v want Unary", expr.Type)
	}
	if !(expr.Operator.Type.Type == "Slash") {
		t.Errorf("unexpected operator, got %v want Slash", expr.Operator.Type.Type)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Right.Value.Literal == 2) {
		t.Errorf("unexpected expression, got %v want 2", expr.Right.Value.Literal)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Left.Value.Literal == 1) {
		t.Errorf("unexpected expression, got %v want 1", expr.Right.Value.Literal)
	}
}

func TestParserComparison(t *testing.T) {
	p, err := makeParser("1 > 2")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Comparison()
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Printf("expr %v", expr)
	if expr.Type != "Binary" {
		t.Errorf("unexpected expression, got %v want Unary", expr.Type)
	}
	if !(expr.Operator.Type.Type == "Greater") {
		t.Errorf("unexpected operator, got %v want Greater", expr.Operator.Type.Type)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Right.Value.Literal == 2) {
		t.Errorf("unexpected expression, got %v want 2", expr.Right.Value.Literal)
	}
	if !(expr.Right.Type == "Literal") {
		t.Errorf("unexpected expression, got %v want Number", expr.Value.Type.Type)
	}
	if !(expr.Left.Value.Literal == 1) {
		t.Errorf("unexpected expression, got %v want 1", expr.Right.Value.Literal)
	}
}

func TestParserExpression(t *testing.T) {
	p, err := makeParser("(1 + 2) != (3 + 4)")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Expression()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Binary" {
		t.Errorf("unexpected expression, got %v want Binary", expr.Type)
	}
	if !(expr.Operator.Type.Type == "BangEqual") {
		t.Errorf("unexpected operator, got %v want BangEqual", expr.Operator.Type.Type)
	}
	if expr.Left.Expression.Type != "Binary" {
		t.Errorf("unexpected left expression type, got %v want Grouping", expr.Left.Expression.Type)
	}
	if expr.Left.Expression.Operator.Type.Type != "Plus" {
		t.Errorf("unexpected left expression operator, got %v want Plus", expr.Left.Expression.Operator.Type.Type)
	}

}

func TestParseIdentifier(t *testing.T) {
	p, err := makeParser("a")
	if err != nil {
		t.Errorf(err.Error())
	}
	expr, err := p.Expression()
	if err != nil {
		t.Errorf(err.Error())
	}
	if expr.Type != "Variable" {
		t.Errorf("unexpected expression, got %v want Variable", expr.Type)
	}
	if expr.Name.Lexeme != "a" {
		t.Errorf("unexpected variable name, got %v want 'a'", expr.Name.Lexeme)

	}

}
