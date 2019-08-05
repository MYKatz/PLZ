package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MYKatz/PLZ/evaluator"
	"github.com/MYKatz/PLZ/lexer"
	"github.com/MYKatz/PLZ/parser"
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
				continue
			}

			evaluated := evaluator.Eval(prog)
			if evaluated != nil {
				io.WriteString(w, evaluated.Inspect())
				io.WriteString(w, "\n")
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
