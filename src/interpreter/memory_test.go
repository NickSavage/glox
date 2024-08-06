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
