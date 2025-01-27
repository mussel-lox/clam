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
func (d *Diagnostic) At(line, column int) *Diagnostic {
	d.position = NewPosition(line, column)
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

	_, _ = fmt.Fprintf(builder, "%sin %s (%s)\n", filenameIndent, d.source.name, d.position.String())
	startLine := max(0, d.position.Line-contextLines)
	lineNumberFormat := fmt.Sprintf("%%%dv ", digitsOf(d.position.Line+1))
	for i := startLine; i <= d.position.Line; i++ {
		_, _ = fmt.Fprint(builder, sourceLineIndent)
		printSource(builder, lineNumberFormat, i+1)
		printSource(builder, "%s\n", d.source.lines[i])
	}
	_, _ = fmt.Fprint(builder, sourceLineIndent)
	_, _ = fmt.Fprintf(builder, lineNumberFormat, "")
	for range d.position.Column {
		_, _ = fmt.Fprint(builder, " ")
	}
	printErrorUnderline(builder, "^ around here\n")
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
