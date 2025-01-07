// Package scanner provides a simple API to convert source text into tokens.
package scanner

import (
	"fmt"

	"github.com/mussel-lox/clam/internal/diagnostic"
)

var (
	doubleOperatorMap = map[rune]doubleOperatorRule{
		'!': newDoubleOperatorRule('=', TokBangEqual, TokBang),
		'=': newDoubleOperatorRule('=', TokEqualEqual, TokEqual),
		'>': newDoubleOperatorRule('=', TokGreaterEqual, TokGreater),
		'<': newDoubleOperatorRule('=', TokLessEqual, TokLess),
	}

	singleOperatorMap = map[rune]TokenKind{
		'(': TokLeftParenthesis,
		')': TokRightParenthesis,
		'{': TokLeftBrace,
		'}': TokRightBrace,
		',': TokComma,
		'.': TokDot,
		'-': TokMinus,
		'+': TokPlus,
		';': TokSemicolon,
		'/': TokSlash,
		'*': TokStar,
	}

	keywordMap = map[string]TokenKind{
		"and":    TokAnd,
		"class":  TokClass,
		"else":   TokElse,
		"false":  TokFalse,
		"for":    TokFor,
		"fun":    TokFun,
		"if":     TokIf,
		"nil":    TokNil,
		"or":     TokOr,
		"print":  TokPrint,
		"return": TokReturn,
		"super":  TokSuper,
		"this":   TokThis,
		"true":   TokTrue,
		"var":    TokVar,
		"while":  TokWhile,
	}
)

type doubleOperatorRule struct {
	expect    rune
	then      TokenKind
	otherwise TokenKind
}

func newDoubleOperatorRule(expect rune, then, otherwise TokenKind) doubleOperatorRule {
	return doubleOperatorRule{
		expect:    expect,
		then:      then,
		otherwise: otherwise,
	}
}

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
	s.skipWhitespaceAndComment()

	first, exists := s.advance()
	if !exists {
		return nil, nil
	}

	if rule, exists := doubleOperatorMap[first]; exists {
		if peek, exists := s.peek(); exists {
			if peek == rule.expect {
				s.advance()
				return s.submit(rule.then)
			}
		}
		return s.submit(rule.otherwise)
	}

	if operator, exists := singleOperatorMap[first]; exists {
		return s.submit(operator)
	}

	switch {
	case isAlphaOrUnderscore(first):
		return s.scanIdentifierOrKeyword()
	case isDigit(first):
		return s.scanNumber()
	case first == '"':
		return s.scanString()
	default:
		return nil, s.diagnose("unrecognized token %v", first)
	}
}

func (s *scanner) scanIdentifierOrKeyword() (*Token, error) {
	for {
		peek, exists := s.peek()
		if !exists || !isAlphaOrUnderscore(peek) && !isDigit(peek) {
			break
		}
		s.advance()
	}
	lexeme := s.source.Slice(s.start, s.current)
	if keywordKind, exists := keywordMap[lexeme]; exists {
		return s.submit(keywordKind)
	}
	return s.submit(TokIdentifier)
}

func (s *scanner) scanNumber() (*Token, error) {
	for {
		peek, exists := s.peek()
		if !exists {
			break
		}

		if !isDigit(peek) {
			if peek == '.' {
				s.advance()
				for {
					frac, exists := s.peek()
					if !exists || !isDigit(frac) {
						break
					}
					s.advance()
				}
			} else {
				break
			}
		}
		s.advance()
	}
	return s.submit(TokNumber)
}

func (s *scanner) scanString() (*Token, error) {
	terminated := false
	for {
		char, exists := s.advance()
		if !exists {
			break
		}
		if char == '"' {
			terminated = true
			break
		}
	}
	if !terminated {
		return nil, s.diagnose("unterminated string")
	}
	return s.submit(TokString)
}

func (s *scanner) diagnose(format string, a ...any) error {
	return diagnostic.NewDiagnostic(fmt.Sprintf(format, a...)).
		Attach(s.source).
		At(s.start, s.current-1)
}

func (s *scanner) submit(kind TokenKind) (*Token, error) {
	token := &Token{
		Kind:   kind,
		Lexeme: s.source.Slice(s.start, s.current),
		Position: diagnostic.Position{
			Start: s.start,
			End:   s.current - 1,
		},
	}
	s.start = s.current
	return token, nil
}

func (s *scanner) skipWhitespaceAndComment() {
outer:
	for {
		peek, exists := s.peek()
		if !exists {
			break
		}

		switch peek {
		case ' ', '\t', '\r', '\n':
			s.advance()
		case '/':
			if next, exists := s.peek(1); exists && next == '/' {
				for {
					char, exists := s.advance()
					if !exists || char == '\n' {
						break
					}
				}
			} else {
				break outer
			}
		default:
			break outer
		}
	}
	s.start = s.current
}

func (s *scanner) advance() (char rune, exists bool) {
	if s.current < s.source.Len() {
		s.current++
		return s.source.At(s.current - 1), true
	}
	return 0, false
}

func (s *scanner) peek(offset ...int) (char rune, exists bool) {
	if len(offset) > 1 {
		panic("peek with too many arguments")
	}
	position := s.current
	if len(offset) == 1 {
		if offset[0] < 0 {
			panic("cannot peek with negative offset")
		}
		position += offset[0]
	}

	if position > 0 && position < s.source.Len() {
		return s.source.At(position), true
	}
	return 0, false
}
