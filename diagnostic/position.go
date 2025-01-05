package diagnostic

type Position struct {
	Start int
	End   int
}

func NewPosition(start, end int) Position {
	return Position{Start: start, End: end}
}

func NewPositionAt(n int) Position {
	return Position{Start: n, End: n}
}

type lineBasedPosition struct {
	Line  int
	Start int
	End   int
}

func (p Position) transform(s *Source) []lineBasedPosition {
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
