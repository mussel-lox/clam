// Package parser provides an abstract API to parse source into ASTs.
package parser

import (
	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/parser/peg"
)

// Parse is the only API public to outside. Inner implementation may change any time.
func Parse(source string) (ast.Expression, error) {
	expr, err := peg.ParseWithDiagnostic("<script>", source)
	if err != nil {
		return nil, err
	}
	return expr.(ast.Expression), nil
}
