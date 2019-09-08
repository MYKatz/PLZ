package interpreter

import (
	"fmt"

	"github.com/MYKatz/PLZ/evaluator"
	"github.com/MYKatz/PLZ/lexer"
	"github.com/MYKatz/PLZ/object"
	"github.com/MYKatz/PLZ/parser"
)

func Interpret(in string) {
	out := ""
	l := lexer.NewLexer(in)
	p := parser.NewParser(l)

	prog := p.ParseProgram()

	env := object.NewEnvironment()

	if len(p.Errors()) != 0 {
		out += "\tOops, there were some errors: \n"
		for _, e := range p.Errors() {
			out += "\t " + e + "\n"
		}
	}

	evaluated := evaluator.Eval(prog, env)
	if evaluated != nil {
		out += evaluated.Inspect() + "\n"
	}

	fmt.Print(out)
}
