// Package parser provides a simple API to turn source code into AST.
package parser

import (
	"github.com/mussel-lox/clam/ast"
)

// Parse turns the source text into a cluster of [ast.Declaration]s.
//
// Lox programs consist of [ast.Declaration]s. Internally, the parser will split source text into tokens, and match the
// syntax rules on the token stream.
//
//revive:disable
func Parse(filename, source string) ([]ast.Declaration, error) {
	panic("unimplemented")
}
