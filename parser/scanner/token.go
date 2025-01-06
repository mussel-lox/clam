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

var tokenKindNames = []string{
	"LeftParenthesis",
	"RightParenthesis",
	"LeftBrace",
	"RightBrace",
	"Comma",
	"Dot",
	"Minus",
	"Plus",
	"Semicolon",
	"Slash",
	"Star",
	"Bang",
	"BangEqual",
	"Equal",
	"EqualEqual",
	"Greater",
	"GreaterEqual",
	"Less",
	"LessEqual",
	"Identifier",
	"String",
	"Number",
	"And",
	"Class",
	"Else",
	"False",
	"For",
	"Fun",
	"If",
	"Nil",
	"Or",
	"Print",
	"Return",
	"Super",
	"This",
	"True",
	"Var",
	"While",
}

//revive:enable

// TokenKind is the category of token (e.g. Identifier).
type TokenKind int

func (k TokenKind) String() string {
	return tokenKindNames[k]
}

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
