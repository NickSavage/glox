package interpreter

import (
	"testing"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func parseSource(t *testing.T, text string) (Interpreter, error) {
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
		Memory:     make(map[string]interface{}),
	}
	return i, nil
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
