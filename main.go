//revive:disable-line
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mussel-lox/clam/internal/diagnostic"
	"github.com/mussel-lox/clam/parser/scanner"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(line) == "" {
			break
		}

		tokens, err := scanner.Scan(diagnostic.NewSource("<script>", line))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, token := range tokens {
			fmt.Printf("%2d `%s`\n", token.Kind, token.Lexeme)
		}
	}
}
