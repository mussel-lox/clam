package diagnostic

import "strings"

type Source interface {
	Name() string
	LineAt(n int) string
}

type SimpleSource struct {
	name  string
	lines []string
}

func NewSimpleSource(name, text string) *SimpleSource {
	return &SimpleSource{
		name:  name,
		lines: strings.Split(strings.ReplaceAll(text, "\r", ""), "\n"),
	}
}

func (s SimpleSource) Name() string        { return s.name }
func (s SimpleSource) LineAt(n int) string { return s.lines[n] }
