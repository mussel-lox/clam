package parser

import (
	"strings"
	"testing"

	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/parser/grammar"
)

var (
	primaryExamples = map[string]string{
		"true":            repr(ast.BooleanLiteral(true)),
		"false":           repr(ast.BooleanLiteral(false)),
		"nil":             repr(ast.Nil{}),
		"this":            repr(ast.This{}),
		"super":           repr(ast.Super{}),
		"12.34":           repr(ast.NumberLiteral(12.34)),
		"\"some string\"": repr(ast.StringLiteral("\"some string\"")),
		"snake_camelCase": repr(ast.Identifier("snake_camelCase")),
	}

	advancedExamples = map[string]string{
		"-1":    repr(ast.UnaryExpression{Operator: ast.UopNegate, Operand: ast.NumberLiteral(1)}),
		"!true": repr(ast.UnaryExpression{Operator: ast.UopLogicalNot, Operand: ast.BooleanLiteral(true)}),
		"2 * 3": repr(ast.BinaryExpression{
			Left:     ast.NumberLiteral(2),
			Right:    ast.NumberLiteral(3),
			Operator: ast.BinopMultiply,
		}),
	}
)

func TestPrimary(t *testing.T) {
	for source, expect := range primaryExamples {
		source = randomSpaced(source)
		expr, err := grammar.ParseReader(
			"expression_test.go",
			strings.NewReader(source),
			grammar.Entrypoint("Primary"),
		)
		if err != nil {
			t.Error(err)
		}

		result := repr(expr)
		if result != expect {
			t.Errorf("For %s, expect %s, got %s\n", source, expect, result)
		}
	}
}

func TestAdvanced(t *testing.T) {
	for source, expect := range advancedExamples {
		source = randomSpaced(source)
		expr, err := grammar.ParseReader(
			"expression_test.go",
			strings.NewReader(source),
			grammar.Entrypoint("Factor"),
		)
		if err != nil {
			t.Error(err)
		}

		result := repr(expr)
		if result != expect {
			t.Errorf("For %s, expect %s, got %s\n", source, expect, result)
		}
	}
}
