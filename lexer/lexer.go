package lexer

import (
	"github.com/MYKatz/PLZ/token"
)

type Lexer struct {
	input        string
	position     int  //index of current character
	readPosition int  //current reading position -> next char to read
	ch           byte //current char to evaluate
}

func NewLexer(inp string) *Lexer {
	l := &Lexer{input: inp}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //sets character to null value
	} else {
		l.ch = l.input[l.readPosition] //gets char at readPosition
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) checkChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition] //next char
	}
}

func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()
	var tok token.Token
	switch l.ch {
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '!':
		if l.checkChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
			l.readChar()
		} else {
			tok = token.NewToken(token.EXCLAMATION, l.ch)
		}
	case '=':
		if l.checkChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch) //this is what allows declarations to optionally use '='
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = token.NewToken(token.LBRACKET, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok //early exit
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNum()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.position    //start position
	for isLetter(l.ch) { //while loop
		l.readChar() //advances lexer position variables till next token
	}
	return l.input[pos:l.position] //slice input from start of current token to 'next' one
}

func (l *Lexer) readNum() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}
