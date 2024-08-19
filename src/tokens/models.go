package tokens

type TokenType struct {
	Type string
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func EOFToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "EOF"},
		Lexeme:  "",
		Literal: nil,
		Line:    line,
	}
}

func LeftBracketToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "LeftBracket"},
		Lexeme:  "[",
		Literal: nil,
		Line:    line,
	}
}

func RightBracketToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "RightBracket"},
		Lexeme:  "]",
		Literal: nil,
		Line:    line,
	}
}
func LeftParenToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "LeftParen"},
		Lexeme:  "(",
		Literal: nil,
		Line:    line,
	}
}

func RightParenToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "RightParen"},
		Lexeme:  ")",
		Literal: nil,
		Line:    line,
	}
}
func LeftBraceToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "LeftBrace"},
		Lexeme:  "{",
		Literal: nil,
		Line:    line,
	}
}

func RightBraceToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "RightBrace"},
		Lexeme:  "}",
		Literal: nil,
		Line:    line,
	}
}
func CommaToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Comma"},
		Lexeme:  ",",
		Literal: nil,
		Line:    line,
	}
}

func DotToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Dot"},
		Lexeme:  ".",
		Literal: nil,
		Line:    line,
	}
}

func MinusToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Minus"},
		Lexeme:  "-",
		Literal: nil,
		Line:    line,
	}
}

func PlusToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Plus"},
		Lexeme:  "+",
		Literal: nil,
		Line:    line,
	}
}
func TildeToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Tilde"},
		Lexeme:  "~",
		Literal: nil,
		Line:    line,
	}
}

func SemicolonToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Semicolon"},
		Lexeme:  ";",
		Literal: nil,
		Line:    line,
	}
}

func ColonToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Colon"},
		Lexeme:  ":",
		Literal: nil,
		Line:    line,
	}
}

func StarToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Star"},
		Lexeme:  "*",
		Literal: nil,
		Line:    line,
	}
}

func SlashToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Slash"},
		Lexeme:  "/",
		Literal: nil,
		Line:    line,
	}
}

func BangToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Bang"},
		Lexeme:  "!",
		Literal: nil,
		Line:    line,
	}
}

func BangEqualToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "BangEqual"},
		Lexeme:  "!=",
		Literal: nil,
		Line:    line,
	}
}
func EqualEqualToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "EqualEqual"},
		Lexeme:  "==",
		Literal: nil,
		Line:    line,
	}
}
func EqualToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Equal"},
		Lexeme:  "=",
		Literal: nil,
		Line:    line,
	}
}

func LessToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Less"},
		Lexeme:  "<",
		Literal: nil,
		Line:    line,
	}
}

func GreaterToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "Greater"},
		Lexeme:  ">",
		Literal: nil,
		Line:    line,
	}
}

func LessEqualToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "LessEqual"},
		Lexeme:  "<=",
		Literal: nil,
		Line:    line,
	}
}

func GreaterEqualToken(line int) Token {
	return Token{
		Type:    TokenType{Type: "GreaterEqual"},
		Lexeme:  ">=",
		Literal: nil,
		Line:    line,
	}
}

func IdentifierToken(name string) Token {
	return Token{
		Type:    TokenType{Type: "Identifier"},
		Lexeme:  name,
		Literal: nil,
		Line:    0,
	}
}
