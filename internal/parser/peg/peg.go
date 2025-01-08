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

//revive:disable

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

//revive:enable

// ParseWithDiagnostic turns internal parserError into Diagnostic, which is more friendly to read.
func ParseWithDiagnostic(filename, source string) ([]ast.Declaration, error) {
	src := diagnostic.NewSource(filename, source)

	var builder strings.Builder
	expr, err := ParseReader(filename, strings.NewReader(source), Entrypoint("Program"))
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
	return expr.([]ast.Declaration), nil
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
