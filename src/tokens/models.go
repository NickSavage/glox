package tokens

type TokenType struct {
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}
