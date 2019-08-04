package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MYKatz/PLZ/lexer"
	"github.com/MYKatz/PLZ/parser"
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
			p := parser.NewParser(l)

			prog := p.ParseProgram()

			if len(p.Errors()) != 0 {
				printParserErrors(w, p.Errors())
				return
			}

			io.WriteString(w, prog.String())
			io.WriteString(w, "\n")

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { //process input token-by-token until end-of-file
				fmt.Printf("%+v \n", tok)
			}
		}
	}
}

func printParserErrors(w io.Writer, errs []string) {
	io.WriteString(w, "\tOops, there were some errors: \n")
	for _, e := range errs {
		io.WriteString(w, "\t  "+e+"\n")
	}
}
