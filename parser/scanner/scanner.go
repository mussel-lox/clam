package scanner

import "github.com/mussel-lox/clam/internal/diagnostic"

func Scan(source *diagnostic.Source) ([]*Token, *diagnostic.Diagnostic) {
	s := newScanner(source)
	tokens := make([]*Token, 0)
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

func (s *scanner) Scan() (*Token, *diagnostic.Diagnostic) {
	panic("unimplemented")
}
