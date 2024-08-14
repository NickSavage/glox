package tokens

import (
	"testing"
)

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

	text := "'hello world';"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 3 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 3)
	}
	if s.Tokens[0].Lexeme != "hello world" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "hello world")
	}

}

func TestScanDoubleQuoteStrings(t *testing.T) {
	text := "\"hello world\";"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 3 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 3)
	}
	if s.Tokens[0].Lexeme != "hello world" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "hello world")
	}

}

func TestIsDigit(t *testing.T) {
	digits := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	not_digits := [10]rune{'a', 'A', 'z', 'Z', '>', ' ', '\n', '\t'}
	for _, digit := range digits {
		if !isDigit(digit) {
			t.Errorf("wrong result from isDigit for %v, got false want true", digit)
		}
	}
	for _, not_digit := range not_digits {
		if isDigit(not_digit) {
			t.Errorf("wrong result from isDigit for %v, got true want false", not_digit)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	alphas := [6]rune{'a', 'm', 'z', 'A', 'Z', '_'}
	not_alphas := [6]rune{'1', '0', '>', '\n', '\t'}

	for _, alpha := range alphas {
		if !isAlpha(alpha) {
			t.Errorf("wrong result from isalpha for %v, got false want true", alpha)
		}
	}
	for _, not_alpha := range not_alphas {
		if isAlpha(not_alpha) {
			t.Errorf("wrong result from isalpha for %v, got true want false", not_alpha)
		}
	}
}

func TestParseInteger(t *testing.T) {

	text := "1234"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Literal != 1234 {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Literal, 1234)
	}
}

func TestParseFloat(t *testing.T) {

	text := "1234.56"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Literal != 1234.56 {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Literal, 1234.56)
	}
}

func TestParseIdentifier(t *testing.T) {

	text := "abcd"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Type.Type != "Identifier" {
		t.Errorf("wrong token type, got %v want Identifier", s.Tokens[0].Type)
	}
	if s.Tokens[0].Lexeme != "abcd" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "abcd")
	}
}
func TestParseKeywordIdentifier(t *testing.T) {

	text := "and"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Type.Type != "And" {
		t.Errorf("wrong token type, got %v want Identifier", s.Tokens[0].Type)
	}
	if s.Tokens[0].Lexeme != "and" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "and")
	}
}

func TestParseComment(t *testing.T) {
	text := "//test\nand"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 2 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 2)
	}
	if s.Tokens[0].Type.Type != "And" {
		t.Errorf("wrong token type, got %v want Identifier", s.Tokens[0].Type)
	}
	if s.Tokens[0].Lexeme != "and" {
		t.Errorf("wrong result, got %v want %v", s.Tokens[0].Lexeme, "and")
	}

}

func TestScanGroup(t *testing.T) {
	text := "(1 + 1)"
	s := makeScanner(text)
	err := s.ScanTokens()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(s.Tokens) != 6 {
		t.Errorf("wrong number of tokens returned. got %v want %v", len(s.Tokens), 6)
	}
}
