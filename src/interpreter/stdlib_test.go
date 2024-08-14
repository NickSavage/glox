package interpreter

import (
	"testing"
)

func TestNativeLen(t *testing.T) {
	
	memory := &Storage{
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
	i := Interpreter{
		Expression: declarations[0].Expression,
		Memory:     memory,
	}
	i.LoadNativeFunctions()
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
