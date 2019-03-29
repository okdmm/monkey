package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/okdmm/monkey/compiler"
	"github.com/okdmm/monkey/lexer"
	"github.com/okdmm/monkey/parser"
	"github.com/okdmm/monkey/vm"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		//env := object.NewEnvironment()
		//macroEnv := object.NewEnvironment()
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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation faild:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.ByteCode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")

		//evaluator.DefineMacros(program, macroEnv)
		//expanded := evaluator.ExpandMacros(program, macroEnv)

		//evaluated := evaluator.Eval(expanded, env)

		//if evaluated != nil {
		//	io.WriteString(out, evaluated.Inspect())
		//	io.WriteString(out, "\n")
		//}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
