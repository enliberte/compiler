package main

import (
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main()  {
	content, e := ioutil.ReadFile(os.Args[1])
	if e != nil {
		io.WriteString(os.Stdout, fmt.Sprintf("Error: %q", e.Error()) )
	}

	sourceCode := string(content)
	l := lexer.New(sourceCode)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
	}
}
