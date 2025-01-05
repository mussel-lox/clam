//revive:disable-line
package main

import (
	"fmt"
	"os"

	"github.com/mussel-lox/clam/internal/diagnostic"
)

func main() {
	var err error = diagnostic.NewDiagnostic("don't say this").
		Attach(diagnostic.NewSource("<script>", "Hello, World")).
		At(0, 3)
	fmt.Fprintln(os.Stderr, err)
}
