package diagnostic

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const contextLines = 2

var (
	printErrorTag       = color.New(color.FgRed).FprintFunc()
	printErrorMessage   = color.New(color.FgHiWhite).FprintlnFunc()
	printErrorUnderline = color.New(color.FgHiRed).FprintFunc()
	printSource         = color.New(color.FgHiBlack).FprintfFunc()

	filenameIndent   = "  "
	sourceLineIndent = "    "
)

type Diagnostic struct {
	message  string
	source   *Source
	position *Position
}

func NewDiagnostic(message string) *Diagnostic {
	return &Diagnostic{message: message}
}

func (d *Diagnostic) At(start, end int) *Diagnostic {
	d.position = NewPosition(start, end)
	return d
}

func (d *Diagnostic) AtChar(offset int) *Diagnostic {
	d.position = NewPositionAt(offset)
	return d
}

func (d *Diagnostic) Attach(source *Source) *Diagnostic {
	d.source = source
	return d
}

func (d *Diagnostic) Display() {
	writer := bufio.NewWriter(os.Stderr)
	defer writer.Flush()

	printErrorTag(writer, "error: ")
	printErrorMessage(writer, d.message)
	if d.source == nil || d.position == nil {
		return
	}

	linepos := d.position.transform(d.source)
	for _, pos := range linepos {
		fmt.Fprintf(writer, "%sat %s (%s)\n", filenameIndent, d.source.name, pos.String())
		startLine := max(0, pos.Line-contextLines)
		lineNumberFormat := fmt.Sprintf("%%%dv ", digitsOf(pos.Line+1))
		for i := startLine; i <= pos.Line; i++ {
			fmt.Fprint(writer, sourceLineIndent)
			printSource(writer, lineNumberFormat, i)
			printSource(writer, "%s\n", d.source.lines[i])
		}
		fmt.Fprint(writer, sourceLineIndent)
		fmt.Fprintf(writer, lineNumberFormat, "")
		for range pos.Start {
			fmt.Fprint(writer, " ")
		}
		for range pos.End - pos.Start + 1 {
			printErrorUnderline(writer, "^")
		}
		printErrorUnderline(writer, " at here\n")
	}
}

func digitsOf(x int) int {
	digits := 0
	for x > 0 {
		digits++
		x /= 10
	}
	return digits
}
