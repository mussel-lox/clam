package parser

import (
	"encoding/json"
	"math/rand/v2"
	"strings"
)

var (
	spaces           = []byte{' ', '\t', '\r', '\n'}
	maxRandomSpacing = len(spaces)
)

func randomSpaced(str string) string {
	var builder strings.Builder
	for i := range rand.Perm(rand.Int() % maxRandomSpacing) {
		builder.WriteByte(spaces[i])
	}
	builder.WriteString(str)
	for i := range rand.Perm(rand.Int() % maxRandomSpacing) {
		builder.WriteByte(spaces[i])
	}
	return builder.String()
}

func repr(value any) string {
	serialized, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(serialized)
}
