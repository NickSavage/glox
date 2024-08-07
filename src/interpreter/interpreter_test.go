package interpreter

import (
	"log"
	"testing"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func parseSource(t *testing.T, text string) (Interpreter, error) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}
	s := tokens.Scanner{
		Source: text,
		Tokens: make([]tokens.Token, 0),
	}
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
		return Interpreter{}, nil
	}
	p := parser.Parser{Tokens: s.Tokens, Current: 0}
	expr, _ := p.Expression()
	i := Interpreter{
		Expression: expr,
		Memory:     memory,
	}
	return i, nil
}

func parseDeclarations(t *testing.T, text string) ([]*parser.Statement, error) {
	s := tokens.Scanner{
		Source: text,
		Tokens: make([]tokens.Token, 0),
	}
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	p := parser.Parser{Tokens: s.Tokens, Current: 0}
	declarations, err := p.Parse()
	return declarations, err

}

func TestInterpretLiteral(t *testing.T) {
	text := "1"
	i, _ := parseSource(t, text)
	result, err := i.evaluateLiteral(i.Expression)
	if err.HasError {
		t.Errorf(err.Message.Error())
	}
	if _, ok := result.(int); !ok {
		t.Errorf("expected int, got %T", result)
	}
	if result != 1 {
		t.Errorf("incorrect result, got %v want %v", result, 1)
	}

}

func TestInterpretGrouping(t *testing.T) {
	text := "(1 + 1) + (4 / 2)"
	i, _ := parseSource(t, text)
	result, err := i.Evaluate(i.Expression)
	if err.HasError {
		t.Errorf(err.Message.Error())
	}
	if _, ok := result.(int); !ok {
		t.Errorf("expected int, got %T", result)
	}
	if result != 4 {
		t.Errorf("incorrect result, got %v want %v", result, 4)
	}

}

func TestInterpretComparison(t *testing.T) {
	text := "(1 + 1) > (4 / 2)"
	i, _ := parseSource(t, text)
	result, err := i.Evaluate(i.Expression)
	if err.HasError {
		t.Errorf(err.Message.Error())
	}
	if _, ok := result.(bool); !ok {
		t.Errorf("expected bool, got %T", result)
	}
	if result != false {
		t.Errorf("incorrect result, got %v want %v", result, false)
	}
}

func TestInterpretEquality(t *testing.T) {
	text := "(1 + 1) == (6 / 2)"
	i, _ := parseSource(t, text)
	result, err := i.Evaluate(i.Expression)
	if err.HasError {
		t.Errorf(err.Message.Error())
	}
	if _, ok := result.(bool); !ok {
		t.Errorf("expected bool, got %T", result)
	}
	if result != false {
		t.Errorf("incorrect result, got %v want %v", result, false)
	}
}

func TestInterpretNotEquality(t *testing.T) {
	text := "(1 + 1) != (2 / 2)"
	i, _ := parseSource(t, text)
	result, err := i.Evaluate(i.Expression)
	if err.HasError {
		t.Errorf(err.Message.Error())
	}
	if _, ok := result.(bool); !ok {
		t.Errorf("expected bool, got %T", result)
	}
	if result != true {
		t.Errorf("incorrect result, got %v want %v", result, true)
	}
}

func TestInterpretDivideZeroEquality(t *testing.T) {
	text := "(0 / 0) != (0 / 0)"
	i, _ := parseSource(t, text)
	result, err := i.Evaluate(i.Expression)
	if !err.HasError {
		t.Errorf(err.Message.Error())
	}
	if result != nil {
		t.Errorf("result should have been null after error")
	}
}

func TestInterpretVariables(t *testing.T) {
	text := "var a = 1 + 1; print a;"
	declarations, err := parseDeclarations(t, text)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(declarations) != 2 {
		t.Errorf("wrong number of declarations, got %v want %v", len(declarations), 2)
	}
	log.Printf("dec %v", declarations[1])
	if declarations[1].Expression.Name.Lexeme != "a" {
		t.Errorf("wrong variable returned, got %v want %v", declarations[1].Expression.Name.Lexeme, "a")
	}

}
