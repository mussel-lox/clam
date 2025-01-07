//go:generate pigeon -nolint -o generated.go grammar.peg

// Package peg contains grammar definitions using PEG and generated parser from pigeon.
package peg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/diagnostic"
)

// ParseWithDiagnostic turns internal parserError into Diagnostic, which is more friendly to read.
func ParseWithDiagnostic(filename, source string) (ast.Expression, error) {
	src := diagnostic.NewSource(filename, source)

	var builder strings.Builder
	expr, err := ParseReader(filename, strings.NewReader(source), Entrypoint("Expression"))
	if err != nil {
		for _, err := range err.(errList) {
			e := err.(*parserError)
			diag := diagnostic.NewDiagnostic(fmt.Sprint(e.Inner.Error())).
				At(e.pos.line-1, e.pos.col-1).
				Attach(src)
			fmt.Fprintln(&builder, diag)
		}
		return nil, errors.New(builder.String())
	}
	return expr.(ast.Expression), nil
}
