package interpreter

import (
	//	"log"
	"testing"
	//"github.com/NickSavage/glox/src/parser"
	//"github.com/NickSavage/glox/src/tokens"
)

func TestDefineCallFunction(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "var a = 1; func hello(hi) { a = 2;} hello('test');"
	declarations, err := parseDeclarations(t, text)
	if err != nil {
		t.Errorf(err.Error())
	}
	i := Interpreter{
		Expression: declarations[0].Expression,
		Memory:     memory,
	}
	for _, declaration := range declarations {
		i.Expression = declaration.Expression
		rerr := i.Execute(declaration)
		if rerr.HasError {
			t.Errorf(rerr.Message.Error())
		}
	}
	result, err := i.Memory.Get("a")
	if err != nil {
		t.Errorf(err.Error())
	}
	expected := 2
	if result != expected {
		t.Errorf("wrong result, got %v want %v", result, expected)
	}

}
