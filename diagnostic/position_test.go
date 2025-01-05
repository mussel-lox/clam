package diagnostic

import (
	"github.com/mussel-lox/clam/internal"
	"testing"
)

var (
	transformExamples = []transformExample{
		newTransformExample("Hello, World", 0, 3, lineBasedPosition{
			Line:  0,
			Start: 0,
			End:   3,
		}),
		newTransformExample(
			"Several\nLines\r\nBefore\nHello\r\nWorld",
			18,
			21,
			lineBasedPosition{
				Line:  3,
				Start: 0,
				End:   3,
			},
		),
		newTransformExample(
			"Long\nContent\r\nAcross\nLines",
			4,
			20,
			lineBasedPosition{
				Line:  1,
				Start: 0,
				End:   6,
			},
			lineBasedPosition{
				Line:  2,
				Start: 0,
				End:   5,
			},
			lineBasedPosition{
				Line:  3,
				Start: 0,
				End:   3,
			},
		),
	}
)

type transformExample struct {
	source   *Source
	position Position
	expect   string
}

func newTransformExample(source string, start, end int, expect ...lineBasedPosition) transformExample {
	return transformExample{
		source:   NewSource("", source),
		position: NewPosition(start, end),
		expect:   internal.Repr(expect),
	}
}

func TestPositionTransform(t *testing.T) {
	for _, example := range transformExamples {
		linepos := example.position.transform(example.source)
		repr := internal.Repr(linepos)
		if repr != example.expect {
			t.Errorf("expect %s, got %s", example.expect, repr)
		}
	}
}
