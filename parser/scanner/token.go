package scanner

import "github.com/mussel-lox/clam/diagnostic"

const (
	TokLeftParenthesis TokenKind = iota
	TokRightParenthesis
	TokLeftBrace
	TokRightBrace
)

type TokenKind int

type Token struct {
	Kind     TokenKind
	Lexeme   string
	Position diagnostic.Position
}
