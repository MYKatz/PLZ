package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	//operators
	ASSIGN = "BE"
	PLUS   = "+"

	//delimiters
	COMMA      = ","
	TERMINATOR = "PLZ"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"be":       ASSIGN,
	"plz":      TERMINATOR,
	"please":   LBRACE,
	"thanks":   RBRACE,
}

func NewToken(tt TokenType, literal byte) Token {
	t := Token{Type: tt, Literal: string(literal)}
	return t
}

func LookupIdent(identifier string) TokenType {
	tok, ok := keywords[identifier]
	if ok {
		return tok
	}
	return IDENT
}
