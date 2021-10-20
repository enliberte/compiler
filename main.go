package main

import (
	"compiler/repl"
	"os"
)

func main()  {
	repl.Start(os.Stdin, os.Stdout)
}

