package scanner

import "github.com/mussel-lox/clam/internal/diagnostic"

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

	TokIdentifier
	TokString
	TokNumber

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

//revive:enable

// TokenKind is the category of token (e.g. Identifier).
type TokenKind int

// Token is the smallest unit that a parser can recognize. Scanner is the tool to split source text into tokens.
type Token struct {
	Kind     TokenKind
	Lexeme   string
	Position diagnostic.Position
}

func newToken(kind TokenKind, lexeme []rune, position diagnostic.Position) *Token {
	return &Token{
		Kind:     kind,
		Lexeme:   string(lexeme),
		Position: position,
	}
}
