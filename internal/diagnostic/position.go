package diagnostic

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	Line  int
	Start *int
	End   *int
}

func NewPositionAtLine(line int) *Position {
	return &Position{Line: line}
}

func NewPositionAtChar(line, char int) *Position {
	return &Position{Line: line, Start: &char}
}

func NewPosition(line, start, end int) *Position {
	return &Position{Line: line, Start: &start, End: &end}
}

func (p Position) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(p.Line + 1))
	if p.Start != nil {
		builder.WriteString(fmt.Sprintf(":%d", *p.Start+1))
	}
	if p.End != nil {
		builder.WriteString(fmt.Sprintf(":%d", *p.End))
	}
	return builder.String()
}
