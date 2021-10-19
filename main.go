package main

import (
	"os"
	"compiler/repl"
)

func main()  {
	repl.Start(os.Stdin, os.Stdout)
}

