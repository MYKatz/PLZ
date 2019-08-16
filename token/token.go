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
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	//operators
	ASSIGN      = "BE"
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	EXCLAMATION = "!"
	COLON       = ":"

	//comparison
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	//delimiters
	COMMA      = ","
	TERMINATOR = "PLZ"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"be":       ASSIGN,
	"plz":      TERMINATOR,
	"please":   LBRACE,
	"thanks":   RBRACE,
	"True":     TRUE,
	"False":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
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
