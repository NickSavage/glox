package parser

import (
	"log"
	"testing"
	//	"github.com/NickSavage/glox/src/tokens"
)

func TestParseStatements(t *testing.T) {
	p, err := makeParser("print 1; print 2; print 3;")
	if err != nil {
		t.Errorf(err.Error())
	}
	statements, err := p.Parse()
	log.Printf("err %v", err)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(statements) != 3 {
		t.Errorf("not the correct number of statements, got %v want %v", 3, len(statements))

	}
}

func TestParseStatementString(t *testing.T) {
	p, err := makeParser("print 'hello world';")
	if err != nil {
		t.Errorf(err.Error())
	}
	statements, err := p.Parse()
	log.Printf("err %v", err)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(statements) != 1 {
		t.Errorf("not the correct number of statements, got %v want %v", 1, len(statements))

	}

}

func TestParseBlock(t *testing.T) {
	p, err := makeParser("{ var a = 1; a = 2;}")
	if err != nil {
		t.Errorf(err.Error())
	}
	statements, err := p.Parse()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(statements) != 1 {
		t.Errorf("not the correct number of statements, got %v want %v", 1, len(statements))

	}
	if len(statements[0].Statements) != 2 {
		t.Errorf("not the correct number of statements, got %v want %v", 2, len(statements[0].Statements))

	}

}
