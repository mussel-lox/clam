package diagnostic

import "strings"

// Source is a shallow encapsulation of the source code.
//
// This struct contains all information needed to display syntax errors friendly.
type Source struct {
	name             string
	text             []rune
	lines            []string
	prefixSumLengths []int
}

// NewSource creates a [Source] with a name (conventionally filename) and its content string.
func NewSource(name, str string) *Source {
	cleanedText := strings.ReplaceAll(str, "\r", "")
	text := []rune(cleanedText)
	lines := strings.Split(cleanedText, "\n")

	accumulatedLength := 0
	prefixSumLengths := make([]int, len(lines))
	for index, line := range lines {
		accumulatedLength += len(line)
		prefixSumLengths[index] = accumulatedLength
	}

	return &Source{
		name:             name,
		text:             text,
		lines:            lines,
		prefixSumLengths: prefixSumLengths,
	}
}

// At returns the rune of the [Source] text at the specific offset.
func (s *Source) At(n int) rune { return s.text[n] }

// Len returns the length of source code (in []rune).
func (s *Source) Len() int { return len(s.text) }

// Slice returns a source string from a range.
func (s *Source) Slice(start, end int) string {
	return string(s.text[start:end])
}

func (s *Source) lengthOfLine(n int) int {
	if n == 0 {
		return s.prefixSumLengths[0]
	}
	return s.prefixSumLengths[n] - s.prefixSumLengths[n-1]
}

func (s *Source) offsetOnLine(index, line int) int {
	if line == 0 {
		return index
	}
	return index - s.prefixSumLengths[line-1]
}
