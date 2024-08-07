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
