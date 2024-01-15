package main

import (
	"fmt"
	"os"

	"github.com/Kreonn/monkey/internal/repl"
)

func main() {
	fmt.Println("Welcome in the Monkey REPL")
	repl.Start(os.Stdin, os.Stdout)
}
