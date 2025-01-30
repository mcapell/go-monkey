package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mcapell/go-monkey/evaluator"
	"github.com/mcapell/go-monkey/lexer"
	"github.com/mcapell/go-monkey/object"
	"github.com/mcapell/go-monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect()) // nolint: errcheck
			io.WriteString(out, "\n")                // nolint: errcheck
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n") // nolint: errcheck
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n") // nolint: errcheck
	}
}
