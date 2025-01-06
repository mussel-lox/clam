package diagnostic

import (
	"testing"

	"github.com/mussel-lox/clam/internal/represent"
)

// Test snippets.
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

// transformExample is a combination of [Source], [Position] and an expected JSON string. It is used to store the test
// snippets.
type transformExample struct {
	source   *Source
	position *Position
	expect   []lineBasedPosition
}

// newTransformExample is a helper function to create a [transformExample].
func newTransformExample(source string, start, end int, expect ...lineBasedPosition) transformExample {
	return transformExample{
		source:   NewSource("", source),
		position: NewPosition(start, end),
		expect:   expect,
	}
}

func TestPositionTransform(t *testing.T) {
	for _, example := range transformExamples {
		linepos := example.position.transform(example.source)
		if !represent.Eq(example.expect, linepos) {
			t.Errorf("expect %v, got %v", example.expect, &linepos)
		}
	}
}
