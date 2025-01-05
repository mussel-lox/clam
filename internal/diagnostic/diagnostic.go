// Package diagnostic defines some data structures to display syntax errors more friendly.
package diagnostic

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const contextLines = 2

var (
	printErrorTag       = color.New(color.FgRed).FprintFunc()
	printErrorMessage   = color.New(color.FgHiWhite).FprintlnFunc()
	printErrorUnderline = color.New(color.FgHiRed).FprintFunc()
	printSource         = color.New(color.FgHiBlack).FprintfFunc()

	filenameIndent   = strings.Repeat(" ", 2)
	sourceLineIndent = strings.Repeat(" ", 4)
)

// Diagnostic contains all information about a syntax error, and can display the error more friendly.
type Diagnostic struct {
	message  string
	source   *Source
	position *Position
}

// NewDiagnostic creates a [Diagnostic] with only the message.
func NewDiagnostic(message string) *Diagnostic {
	return &Diagnostic{message: message}
}

// At specifies the [d.position] of [d].
func (d *Diagnostic) At(start, end int) *Diagnostic {
	d.position = NewPosition(start, end)
	return d
}

// AtChar sets the [d.position] at a specific character.
func (d *Diagnostic) AtChar(offset int) *Diagnostic {
	d.position = NewPositionAt(offset)
	return d
}

// Attach sets the [d.source] field.
func (d *Diagnostic) Attach(source *Source) *Diagnostic {
	d.source = source
	return d
}

// Error implements the [error] interface, making [Diagnostic] of the [error] type and can be treated as a regular
// [error].
func (d *Diagnostic) Error() string {
	builder := new(strings.Builder)

	printErrorTag(builder, "error: ")
	printErrorMessage(builder, d.message)
	if d.source == nil || d.position == nil {
		return builder.String()
	}

	linepos := d.position.transform(d.source)
	for _, pos := range linepos {
		fmt.Fprintf(builder, "%sin %s (%s)\n", filenameIndent, d.source.name, pos.String())
		startLine := max(0, pos.Line-contextLines)
		lineNumberFormat := fmt.Sprintf("%%%dv ", digitsOf(pos.Line+1))
		for i := startLine; i <= pos.Line; i++ {
			fmt.Fprint(builder, sourceLineIndent)
			printSource(builder, lineNumberFormat, i+1)
			printSource(builder, "%s\n", d.source.lines[i])
		}
		fmt.Fprint(builder, sourceLineIndent)
		fmt.Fprintf(builder, lineNumberFormat, "")
		for range pos.Start {
			fmt.Fprint(builder, " ")
		}
		for range pos.End - pos.Start + 1 {
			printErrorUnderline(builder, "^")
		}
		printErrorUnderline(builder, " at here\n")
	}
	return builder.String()
}

func digitsOf(x int) int {
	digits := 0
	for x > 0 {
		digits++
		x /= 10
	}
	return digits
}
