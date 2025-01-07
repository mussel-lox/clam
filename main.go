//revive:disable-line
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/mussel-lox/clam/internal/parser"
)

func main() {
	printer := repr.New(os.Stdout, repr.OmitEmpty(false))
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(line) == "" {
			break
		}

		expr, err := parser.Parse(line)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			break
		}
		printer.Println(expr)
	}
}
