package stdlib

import (
	"github.com/NickSavage/glox/src/interpreter"
	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
	"testing"
)

func parseSource(t *testing.T, text string) (interpreter.Interpreter, error) {
	memory := &interpreter.Storage{
		Memory: make(map[string]interface{}),
	}
	s := tokens.Scanner{
		Source: text,
		Tokens: make([]tokens.Token, 0),
	}
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
		return interpreter.Interpreter{}, nil
	}
	p := parser.Parser{Tokens: s.Tokens, Current: 0}
	expr, _ := p.Expression()
	i := interpreter.Interpreter{
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
func TestNativeLen(t *testing.T) {

	memory := &interpreter.Storage{
		Memory: make(map[string]interface{}),
	}

	text := " var a = [1,2,3,4]; var b = len(a);"
	declarations, err := parseDeclarations(t, text)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(declarations) != 2 {
		t.Errorf("wrong number of declarations, got %v want %v", len(declarations), 2)
	}
	i := interpreter.Interpreter{
		Expression: declarations[0].Expression,
		Memory:     memory,
	}
	i.Memory.Define("len", LenFunction(&i))
	for _, declaration := range declarations {
		i.Expression = declaration.Expression
		rerr := i.Execute(declaration)
		if rerr.HasError {
			t.Errorf(rerr.Message.Error())
		}
	}
	result, err := i.Memory.Get("b")
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != 4 {
		t.Errorf("wrong result, got %v want %v", result, 4)
	}
}
