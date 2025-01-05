// Package scanner provides a simple API to convert source text into tokens.
package scanner

import "github.com/mussel-lox/clam/internal/diagnostic"

// Scan turns source text into tokens.
func Scan(source *diagnostic.Source) ([]*Token, error) {
	var tokens []*Token
	s := newScanner(source)
	for {
		token, err := s.Scan()
		if err != nil {
			return nil, err
		}
		if token == nil {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

type scanner struct {
	source  *diagnostic.Source
	start   int
	current int
}

func newScanner(source *diagnostic.Source) *scanner {
	return &scanner{
		source:  source,
		start:   0,
		current: 0,
	}
}

func (s *scanner) Scan() (*Token, error) {
	panic("unimplemented")
}
