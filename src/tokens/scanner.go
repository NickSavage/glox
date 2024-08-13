package tokens

import (
	"errors"
	//	"fmt"
	"log"
	"strconv"
)

var keywords = map[string]TokenType{
	"and":      TokenType{Type: "And"},
	"break":    TokenType{Type: "Break"},
	"continue": TokenType{Type: "Continue"},
	"class":    TokenType{Type: "Class"},
	"else":     TokenType{Type: "Else"},
	"false":    TokenType{Type: "False"},
	"for":      TokenType{Type: "For"},
	"func":     TokenType{Type: "Function"},
	"if":       TokenType{Type: "If"},
	"lambda":   TokenType{Type: "Lambda"},
	"nil":      TokenType{Type: "Nil"},
	"or":       TokenType{Type: "Or"},
	"return":   TokenType{Type: "Return"},
	"super":    TokenType{Type: "Super"},
	"this":     TokenType{Type: "This"},
	"true":     TokenType{Type: "True"},
	"var":      TokenType{Type: "Var"},
	"while":    TokenType{Type: "While"},
}

type Scanner struct {
	Source  string
	Tokens  []Token
	line    int
	current int
	start   int
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		char == '_'

}

func isAlphaNumeric(char rune) bool {
	return isDigit(char) || isAlpha(char)
}

func (s *Scanner) next(current int) rune {
	if current >= len(s.Source)-1 {
		return 0 // or any other sentinel value to indicate EOF
	}
	return rune(s.Source[current+1])
}

func (s *Scanner) parseString(start int) (Token, error) {
	var result string
	current := start
	for {
		next := s.next(current)
		if next == 0 {
			return Token{}, errors.New("unterminated string")
		}
		if next == '\n' {
			s.line++
		}
		if next == rune(s.Source[start]) {
			result = s.Source[start + 1 : current+1]
			log.Printf("%s", result)
			break
		}

		current++
	}
	return Token{
		Type:    TokenType{Type: "String"},
		Lexeme:  result,
		Literal: result,
		Line:    s.line,
	}, nil

}

func (s *Scanner) parseNumber(start int) (Token, error) {
	var numberString string
	double := false

	current := start
	for {
		next := s.next(current)
		if !isDigit(next) && !(next == '.') {
			break
		}
		if next == '.' {
			double = true
		}
		if next == 0 {
			break
		}
		if next == '\n' {
			break
		}
		current++
	}
	// TODO: this seems messy and inefficient
	numberString = s.Source[start : current+1]
	var i interface{}
	var err error
	if double {
		i, err = strconv.ParseFloat(numberString, 64)
	} else {
		i, err = strconv.Atoi(numberString)

	}
	if err != nil {
		return Token{}, err
	}
	return Token{
		Type:    TokenType{Type: "Number"},
		Lexeme:  numberString,
		Literal: i,
		Line:    s.line,
	}, nil
}

func (s *Scanner) parseIdentifier(start int) (Token, error) {

	var result string
	current := start
	for {
		next := s.next(current)
		if !isAlphaNumeric(next) {
			break
		}
		current++
	}
	result = s.Source[start : current+1]
	tokenType := keywords[result]
	if tokenType.Type == "" {
		tokenType = TokenType{Type: "Identifier"}
	}
	return Token{
		Type:    tokenType,
		Lexeme:  result,
		Literal: nil,
		Line:    s.line,
	}, nil

}

func (s *Scanner) ScanTokens() error {
	//start := 0
	s.current = 0
	s.line = 1
	var c rune

	for {
		//start = current
		c = rune(s.Source[s.current])

		switch c {
		case '[':
			s.Tokens = append(s.Tokens, LeftBracketToken(s.line))
		case ']':
			s.Tokens = append(s.Tokens, RightBracketToken(s.line))
		case '(':
			s.Tokens = append(s.Tokens, LeftParenToken(s.line))
		case ')':
			s.Tokens = append(s.Tokens, RightParenToken(s.line))
		case '{':
			s.Tokens = append(s.Tokens, LeftBraceToken(s.line))
		case '}':
			s.Tokens = append(s.Tokens, RightBraceToken(s.line))
		case ',':
			s.Tokens = append(s.Tokens, CommaToken(s.line))
		case '.':
			s.Tokens = append(s.Tokens, DotToken(s.line))
		case '-':
			s.Tokens = append(s.Tokens, MinusToken(s.line))
		case '+':
			s.Tokens = append(s.Tokens, PlusToken(s.line))
		case '~':
			s.Tokens = append(s.Tokens, TildeToken(s.line))
		case ';':
			s.Tokens = append(s.Tokens, SemicolonToken(s.line))
		case ':':
			s.Tokens = append(s.Tokens, ColonToken(s.line))
		case '*':
			s.Tokens = append(s.Tokens, StarToken(s.line))
		case '!':
			if s.next(s.current) == '=' {
				s.Tokens = append(s.Tokens, BangEqualToken(s.line))
				s.current++
			} else {
				s.Tokens = append(s.Tokens, BangToken(s.line))
			}
		case '=':
			if s.next(s.current) == '=' {
				s.Tokens = append(s.Tokens, EqualEqualToken(s.line))
				s.current++
			} else {
				s.Tokens = append(s.Tokens, EqualToken(s.line))
			}

		case '<':
			if s.next(s.current) == '=' {
				s.Tokens = append(s.Tokens, LessEqualToken(s.line))
				s.current++
			} else {
				s.Tokens = append(s.Tokens, LessToken(s.line))
			}

		case '>':
			if s.next(s.current) == '=' {
				s.Tokens = append(s.Tokens, GreaterEqualToken(s.line))
				s.current++
			} else {
				s.Tokens = append(s.Tokens, GreaterToken(s.line))
			}
		case '/':
			if s.next(s.current) == '/' {
				// this is a comment, we want to ignore
				for {
					s.current++
					if s.next(s.current) == '\n' {
						break
					}
				}

			} else {
				s.Tokens = append(s.Tokens, SlashToken(s.line))
			}
		case ' ':
		case '\r':
		case '\t':
			// ignore
		case '\n':
			s.line++
		case '"':
			token, err := s.parseString(s.current)
			if err != nil {
				return err
			}
			s.Tokens = append(s.Tokens, token)
			s.current += len(token.Lexeme) +1
		case '\'':
			token, err := s.parseString(s.current)
			if err != nil {
				return err
			}
			s.Tokens = append(s.Tokens, token)
			s.current += len(token.Lexeme) +1

		default:
			if isDigit(c) {
				token, err := s.parseNumber(s.current)
				if err != nil {
					return err
				}
				s.Tokens = append(s.Tokens, token)
				s.current += len(token.Lexeme) - 1
			} else if isAlpha(c) {
				token, err := s.parseIdentifier(s.current)
				if err != nil {
					return err
				}
				s.Tokens = append(s.Tokens, token)
				s.current += len(token.Lexeme) - 1

			} else {
				log.Fatal("Unexpected character")
				return errors.New("Unexpected character")

			}
		}

		// is at end
		if s.next(s.current) == 0 {
			break
		}
		s.current++
	}
	s.Tokens = append(s.Tokens, EOFToken(s.line))

	return nil
}
