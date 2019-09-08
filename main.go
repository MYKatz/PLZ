package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/MYKatz/PLZ/interpreter"
	"github.com/MYKatz/PLZ/repl"
)

func main() {
	code := flag.String("code", "", "code string to run")
	flag.Parse()

	if *code != "" {
		interpreter.Interpret(*code)
	} else {
		openrepl()
	}
}

func openrepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the PLZ language REPL, %s \n", user.Username)
	fmt.Printf("Enter commands: \n")

	repl.Start(os.Stdin, os.Stdout)
}
