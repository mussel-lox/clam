//revive:disable
package pratt

import (
	"strconv"

	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/parsers/pratt/scanner"
)

const (
	PrecAssignment Precedence = iota
	PrecLogicalOr
	PrecLogicalAnd
	PrecEquality
	PrecComparison
	PrecAdditive
	PrecMultiplicative
	PrecUnary
	PrecInvocation
	PrecPrimary
)

// Precedence is the infix operator precedence defined in binary expressions.
type Precedence int

func (p *parser) parseExpression() (ast.Expression, error) {
	panic("unimplemented")
}

func (p *parser) parsePrecedence(precedence Precedence) (ast.Expression, error) {
	panic("unimplemented")
}

func (p *parser) parseUnary() (ast.Expression, error) {
	op, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	var operator ast.UnaryOperator
	switch op.Kind {
	case scanner.TokMinus:
		operator = ast.UopNegate
	case scanner.TokBang:
		operator = ast.UopLogicalNot
	default:
		return nil, p.diagnose("invalid unary operator")
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	return &ast.UnaryExpression{
		Operand:  expr,
		Operator: operator,
	}, nil
}

func (p *parser) parseInvocation() (ast.Expression, error) {
	panic("")
}

func (p *parser) parsePrimary() (ast.Expression, error) {
	token, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	switch token.Kind {
	case scanner.TokTrue:
		return ast.BooleanLiteral(true), nil
	case scanner.TokFalse:
		return ast.BooleanLiteral(false), nil
	case scanner.TokNil:
		return ast.Nil{}, nil
	case scanner.TokThis:
		return ast.This{}, nil
	case scanner.TokSuper:
		return ast.Super{}, nil
	case scanner.TokNumber:
		n, err := strconv.ParseFloat(token.Lexeme, 64)
		if err != nil {
			return nil, err
		}
		return ast.NumberLiteral(n), nil
	case scanner.TokString:
		return ast.StringLiteral(token.Lexeme), nil
	case scanner.TokIdentifier:
		return ast.Identifier(token.Lexeme), nil
	case scanner.TokLeftParenthesis:
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
			return nil, err
		}
		return expr, nil
	default:
		return nil, p.diagnose("invalid primary expression")
	}
}
