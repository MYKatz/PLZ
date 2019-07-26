package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MYKatz/PLZ/lexer"
	"github.com/MYKatz/PLZ/token"
)

const prompt = ">>>"

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	for { //input while(true) loop
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		} else {
			line := scanner.Text()
			l := lexer.NewLexer(line)

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { //process input token-by-token until end-of-file
				fmt.Printf("%+v \n", tok)
			}
		}
	}
}
