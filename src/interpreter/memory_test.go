package interpreter

import (
	"testing"
)

func TestPutGetMemoryData(t *testing.T) {
	text := "'hello world';"
	i, _ := parseSource(t, text)
	i.Put("hello", "world")
	result, err := i.Get("hello")
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != "world" {
		t.Errorf("wrong result, got %v want %v", result, "world")
	}

}

func TestGetUndefined(t *testing.T) {
	text := "'hello world';"
	i, _ := parseSource(t, text)
	_, err := i.Get("hello")
	if err == nil {
		t.Errorf("expected an error and did not receive one")
	}
	expectedMessage := "undefined variable: hello"
	if err.Error() != expectedMessage {
		t.Errorf("got wrong error message, got %v want %v", err.Error(), expectedMessage)
	}

}

func TestVarInit(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}
	var i Interpreter

	text := "var a = 1 + 1;"
	declarations, err := parseDeclarations(t, text)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, declaration := range declarations {
		i = Interpreter{
			Expression: declaration.Expression,
			Memory:     memory,
		}
		rerr := i.Execute(declaration)
		if rerr.HasError {
			t.Errorf(err.Error())
		}
	}

	result, err := i.Get("a")
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != 2 {
		t.Errorf("wrong result, got %v want %v", result, 2)
	}

}

func TestVarInitAssignment(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}
	var i Interpreter

	text := "var a = 1 + 1; a = 3;"
	declarations, err := parseDeclarations(t, text)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, declaration := range declarations {
		i = Interpreter{
			Expression: declaration.Expression,
			Memory:     memory,
		}
		rerr := i.Execute(declaration)
		if rerr.HasError {
			t.Errorf(err.Error())
		}
	}

	result, err := i.Get("a")
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != 3 {
		t.Errorf("wrong result, got %v want %v", result, 3)
	}

}
