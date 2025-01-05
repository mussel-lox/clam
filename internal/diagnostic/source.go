package diagnostic

import "strings"

type Source struct {
	name             string
	lines            []string
	prefixSumLengths []int
}

func NewSource(name, text string) *Source {
	lines := strings.Split(strings.ReplaceAll(text, "\r", ""), "\n")

	accumulatedLength := 0
	prefixSumLengths := make([]int, len(lines))
	for index, line := range lines {
		accumulatedLength += len(line)
		prefixSumLengths[index] = accumulatedLength
	}

	return &Source{
		name:             name,
		lines:            lines,
		prefixSumLengths: prefixSumLengths,
	}
}

func (s Source) lengthOfLine(n int) int {
	if n == 0 {
		return s.prefixSumLengths[0]
	}
	return s.prefixSumLengths[n] - s.prefixSumLengths[n-1]
}

func (s Source) offsetOnLine(index, line int) int {
	if line == 0 {
		return index
	}
	return index - s.prefixSumLengths[line-1]
}
