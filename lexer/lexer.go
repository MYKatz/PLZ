package lexer

type Lexer struct {
	input        string
	position     int  //index of current character
	readPosition int  //current reading position -> next char to read
	ch           byte //current char to evaluate
}

func NewLexer(inp string) *Lexer {
	l := &Lexer{input: inp}
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
