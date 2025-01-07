package diagnostic

import "fmt"

// Position is the type used to highlight the code segments in a diagnostic. After all, we need to print the source code
// line by line.
type Position struct {
	Line   int
	Column int
}

// NewPosition creates a [Position] conveniently.
func NewPosition(line, column int) *Position {
	return &Position{
		Line:   line,
		Column: column,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("line %d, column %d", p.Line+1, p.Column+1)
}
