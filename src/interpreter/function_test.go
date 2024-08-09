package interpreter

import (
	//	"log"
	"testing"
	//"github.com/NickSavage/glox/src/parser"
	//"github.com/NickSavage/glox/src/tokens"
)

func TestDefineFunctionNoParameters(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "func hello() { print 2 + 2;}"
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

}

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

func TestCallReturnFunction(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "func add2(x) { return x + 2; } var a = add2(1);"
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
	expected := 3
	if result != expected {
		t.Errorf("wrong result, got %v want %v", result, expected)
	}

}

func TestIfReturnStatements(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "func a(x) { if x == 1 {return 2;} return 3; } var a = a(1);"
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

func TestFunctionWrongNumberArguments(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "func a(x) { if x == 1 {return 2;} return 3; } var a = a(1);"
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

}
func TestLambdaFunction(t *testing.T) {
	memory := &Storage{
		Memory: make(map[string]interface{}),
	}

	text := "var a = 1; var b = lambda x: return x + 1;; a = b(5);"
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
	expected := 6
	if result != expected {
		t.Errorf("wrong result, got %v want %v", result, expected)
	}

}
