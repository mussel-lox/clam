//revive:disable
package pratt

import (
	"fmt"

	"github.com/mussel-lox/clam/ast"
	"github.com/mussel-lox/clam/internal/diagnostic"
	"github.com/mussel-lox/clam/internal/parsers/pratt/scanner"
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

type parser struct {
	source *diagnostic.Source
	tokens []*scanner.Token
	offset int
}

func newParser(source *diagnostic.Source, tokens []*scanner.Token) *parser {
	return &parser{
		source: source,
		tokens: tokens,
		offset: 0,
	}
}

func (p *parser) advance() (token *scanner.Token, exists bool) {
	if p.offset < len(p.tokens) {
		return p.tokens[p.offset], true
	}
	return nil, false
}

func (p *parser) mustAdvance() (*scanner.Token, error) {
	token, exists := p.advance()
	if !exists {
		return nil, p.diagnose("expected token, got EOF")
	}
	return token, nil
}

func (p *parser) mustConsume(kind scanner.TokenKind) (*scanner.Token, error) {
	token, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}
	if token.Kind == kind {
		return token, nil
	}
	return nil, p.diagnose("expected %s, got %s", kind.String(), token.Kind.String())
}

func (p *parser) diagnose(format string, a ...any) error {
	var position *diagnostic.Position
	if p.offset == 0 {
		if len(p.tokens) == 0 {
			position = diagnostic.NewPositionAt(0)
		} else {
			position = &p.tokens[0].Position
		}
	} else {
		position = &p.tokens[p.offset-1].Position
	}
	return diagnostic.NewDiagnostic(fmt.Sprintf(format, a...)).
		At(position.Start, position.End).
		Attach(p.source)
}
