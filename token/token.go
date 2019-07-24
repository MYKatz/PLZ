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
