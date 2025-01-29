//go:generate pigeon -nolint -o generated.go grammar.peg

// Package peg contains grammar definitions using PEG and generated parser from pigeon.
package peg

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/diagnostic"
)

const (
	TokLeftParenthesis TokenKind = iota
	TokRightParenthesis
	TokLeftBrace
	TokRightBrace
	TokComma
	TokDot
	TokMinus
	TokPlus
	TokSemicolon
	TokSlash
	TokStar
	TokBang
	TokBangEqual
	TokEqual
	TokEqualEqual
	TokGreater
	TokGreaterEqual
	TokLess
	TokLessEqual
	TokAnd
	TokClass
	TokElse
	TokFalse
	TokFor
	TokFun
	TokIf
	TokNil
	TokOr
	TokPrint
	TokReturn
	TokSuper
	TokThis
	TokTrue
	TokVar
	TokWhile
)

var operatorMapping = map[TokenKind]ast.BinaryOperator{
	TokSlash:        ast.BinopDivide,
	TokStar:         ast.BinopMultiply,
	TokMinus:        ast.BinopSubtract,
	TokPlus:         ast.BinopAdd,
	TokGreaterEqual: ast.BinopGreaterEqual,
	TokLessEqual:    ast.BinopLessEqual,
	TokGreater:      ast.BinopGreater,
	TokLess:         ast.BinopLess,
	TokBangEqual:    ast.BinopNotEqual,
	TokEqualEqual:   ast.BinopEqual,
	TokAnd:          ast.BinopLogicalAnd,
	TokOr:           ast.BinopLogicalOr,
}

type TokenKind int

// ParseWithDiagnostic turns internal parserError into Diagnostic, which is more friendly to read.
func ParseWithDiagnostic(filename, source string) ([]ast.Declaration, error) {
	var builder strings.Builder

	src := diagnostic.NewSource(filename, source)
	program, err := ParseReader(filename, strings.NewReader(source), Entrypoint("Program"))
	if err != nil {
		var errorList errList
		if !errors.As(err, &errorList) {
			panic("unreachable case in asserting errList")
		}
		for _, err := range errorList {
			var parserErr *parserError
			if !errors.As(err, &parserErr) {
				panic("unreachable case in asserting *parserError")
			}
			var locatedErr locatedError
			if !errors.As(parserErr.Inner, &locatedErr) {
				panic("unreachable case in asserting locatedError")
			}
			diag := diagnostic.NewDiagnostic(fmt.Sprint(locatedErr.Error())).
				At(locatedErr.line-1, locatedErr.column-1).
				Attach(src)
			_, _ = fmt.Fprintln(&builder, diag)
		}
		return nil, errors.New(builder.String())
	}
	return program.([]ast.Declaration), nil
}

func parseBinary(l, pat any) ast.Expression {
	left := l.(ast.Expression)
	for _, p := range pat.([]any) {
		pattern := p.([]any)
		operator, exists := operatorMapping[pattern[0].(TokenKind)]
		if !exists {
			panic(fmt.Sprint("uncovered operator ", pattern[0].(TokenKind)))
		}
		right := pattern[1].(ast.Expression)

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

type locatedError struct {
	line    int
	column  int
	message string
}

func newLocatedError(c *current, message string) locatedError {
	text := strings.TrimRightFunc(string(c.text), unicode.IsSpace)
	return locatedError{
		line:    c.pos.line,
		column:  c.pos.col + len(text),
		message: message,
	}
}

func (l locatedError) Error() string {
	return l.message
}
