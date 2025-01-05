package diagnostic

import "fmt"

// Position is the range of a specific code segment. It will be used to highlight it in a diagnostic.
type Position struct {
	Start int
	End   int
}

// NewPosition creates a [Position] with the start and end offset.
func NewPosition(start, end int) *Position {
	return &Position{Start: start, End: end}
}

// NewPositionAt creates a [Position] of a specific character in the source.
func NewPositionAt(n int) *Position {
	return &Position{Start: n, End: n}
}

// lineBasedPosition is the internal type used to highlight the code segments in a diagnostic. After all, we need to
// print the source code line by line.
type lineBasedPosition struct {
	Line  int
	Start int
	End   int
}

func (p *lineBasedPosition) String() string {
	if p.Start < p.End {
		return fmt.Sprintf("line %d, column %d:%d", p.Line+1, p.Start+1, p.End+1)
	}
	return fmt.Sprintf("line %d, column %d", p.Line+1, p.Start+1)
}

// transform turns a [Position] into a slice of [lineBasedPosition]s. If the [Position] itself is on a single line, the
// returned slice contains only one element; otherwise, multiple [lineBasedPosition] is returned.
func (p *Position) transform(s *Source) []lineBasedPosition {
	startLine := binarySearchLineOf(p.Start, s.prefixSumLengths)
	endLine := binarySearchLineOf(p.End, s.prefixSumLengths)
	if startLine == endLine {
		return []lineBasedPosition{{
			Line:  startLine,
			Start: s.offsetOnLine(p.Start, startLine),
			End:   s.offsetOnLine(p.End, endLine),
		}}
	}
	positions := make([]lineBasedPosition, endLine-startLine+1)
	positions[0] = lineBasedPosition{
		Line:  startLine,
		Start: s.offsetOnLine(p.Start, startLine),
		End:   s.lengthOfLine(startLine) - 1,
	}
	for i := 1; i < endLine-startLine; i++ {
		positions[i] = lineBasedPosition{
			Line:  startLine + i,
			Start: 0,
			End:   s.lengthOfLine(startLine+i) - 1,
		}
	}
	positions[endLine-startLine] = lineBasedPosition{
		Line:  endLine,
		Start: 0,
		End:   s.offsetOnLine(p.End, endLine),
	}
	return positions
}

// binarySearchLineOf searches the line containing the target offset. In essence, this is a task searching the minimum
// value of a number that's greater than the target.
func binarySearchLineOf(target int, prefixSum []int) int {
	left := 0
	right := len(prefixSum) - 1
	candidate := len(prefixSum)
	for left <= right {
		middle := (left + right) / 2
		if prefixSum[middle] > target {
			candidate = min(candidate, middle)
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return candidate
}
