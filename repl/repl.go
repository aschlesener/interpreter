package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

const PROMPT = ">> "

// Start is a REPL that parses input containing "Monkey" code into a lexer, and prints the parsed tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// Read - scan input
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Evaluate - create lexer from input
		line := scanner.Text()
		l := lexer.New(line)

		// Print - print each parsed token
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
		// Loop
	}
}
