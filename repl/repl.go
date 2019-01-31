package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/okdmm/monkey/evaluator"
	"github.com/okdmm/monkey/lexer"
	"github.com/okdmm/monkey/object"
	"github.com/okdmm/monkey/parser"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		env := object.NewEnvironment()
		macroEnv := object.NewEnvironment()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacro(program, macroEnv)

		evaluated := evaluator.Eval(expand, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
