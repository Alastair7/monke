package repl

import (
	"Alastair7/monke/lexer"
	"Alastair7/monke/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Fprintf(output, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// Starts from the first token and loops until EOF is reached.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}
	}
}