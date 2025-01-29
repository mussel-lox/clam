// Package parser contains the public API to parse a source code string into a slice of [ast.Declaration].
// The internal implementation may be changed any time.
package parser

import (
	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/parser/peg"
)

// Parse is the only API that is stable. The internal implementation (including package peg) may be changed any time.
func Parse(filename, source string) ([]ast.Declaration, error) {
	return peg.ParseWithDiagnostic(filename, source)
}
