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
	output := prettyPrintExpressionTree(input)
	if expectedOutput != output {
		log.Fatalf("output of prettyPrintExpressionTree not correct, got %v want %v", output, expectedOutput)
	}
}
