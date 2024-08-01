package tokens

import (
	"errors"
	"log"
)

type Scanner struct {
	Source  string
	Tokens  []Token
	line    int
	current int
	start   int
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
			result = s.Source[start : current+2]
			break
		}

		current++
	}
	return Token{
		Type:    TokenType{Type: "String"},
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
		case ';':
			s.Tokens = append(s.Tokens, SemicolonToken(s.line))
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
			s.current += len(token.Lexeme)
		case '\'':
			token, err := s.parseString(s.current)
			if err != nil {
				return err
			}
			s.Tokens = append(s.Tokens, token)
			s.current += len(token.Lexeme)

		default:
			log.Fatal("Unexpected character")
			return errors.New("Unexpected character")
		}

		// is at end
		s.current++
		//log.Printf("current %v %v", s.current, s.Tokens)
		if s.current >= len(s.Source) {

			break
		}
	}
	s.Tokens = append(s.Tokens, EOFToken(s.line))

	return nil
}
