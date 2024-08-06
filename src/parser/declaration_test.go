package parser

import (
	"log"
	"testing"
)

func TestParseDeclaration(t *testing.T) {
	p, err := makeParser("var hi = 1;")
	if err != nil {
		t.Errorf(err.Error())
	}
	statements, err := p.Parse()
	log.Printf("err %v", err)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Printf("%v", statements)
	log.Printf("%v", statements[0].VariableName)
	log.Printf("%v", statements[0].Initializer.Value)

	PrettyPrintExpressionTree(statements[0].Initializer, "")
	if len(statements) != 1 {
		t.Errorf("not the correct number of statements, got %v want %v", 1, len(statements))

	}
}
