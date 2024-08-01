package tokens

import "testing"

func makeScanner(source string) Scanner {
	return Scanner{
		Source: source,
		Tokens: make([]Token, 0),
	}
}

func TestScanTokens(t *testing.T) {
	text := "(){}!= \r\t==\n>="
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 8 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 8)
	}
}

func TestScanSingleQuoteStrings(t *testing.T) {

	text := "'hello world'"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Lexeme != "'hello world'" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "'hello world'")
	}

}

func TestScanDoubleQuoteStrings(t *testing.T) {
	text := "\"hello world\""
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Lexeme != "\"hello world\"" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "\"hello world\"")
	}

}
