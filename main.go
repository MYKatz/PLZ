package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MYKatz/PLZ/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the PLZ language REPL, %s \n", user.Username)
	fmt.Printf("Enter commands: \n")

	repl.Start(os.Stdin, os.Stdout)
}
