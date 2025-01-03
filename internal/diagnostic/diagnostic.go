package diagnostic

import (
	"github.com/fatih/color"
	"github.com/mussel-lox/clam/internal/console"
)

const (
	positionIndent             = 2
	sourceLineIndent           = 4
	contextLines               = 2
	spacesBetweenLineAndNumber = 1
)

var (
	errorTag   = color.New(color.FgRed).SprintFunc()
	errorLabel = color.New(color.FgHiRed).SprintFunc()
	sourceText = color.New(color.FgHiBlack).SprintFunc()
)

type Diagnostic struct {
	source   Source
	position *Position
	message  string
}

func NewDiagnostic(message string) *Diagnostic {
	return &Diagnostic{message: message}
}

func (d *Diagnostic) At(position *Position) *Diagnostic { d.position = position; return d }
func (d *Diagnostic) Attach(source Source) *Diagnostic  { d.source = source; return d }

func (d *Diagnostic) Show() {
	defer console.Flush()

	console.WriteLine(errorTag("error:"), d.message)
	if d.source == nil || d.position == nil {
		return
	}
	console.WriteRepeated(" ", positionIndent)
	console.WriteLine("at", d.source.Name(), d.position)

	startLine := max(d.position.Line-contextLines, 0)
	maxLineNumberDigits := digitsOf(d.position.Line + 1)
	for i := startLine; i <= d.position.Line; i++ {
		console.WriteRepeated(" ", sourceLineIndent)
		console.WriteRepeated(" ", maxLineNumberDigits-digitsOf(i+1))
		console.Write(sourceText(i + 1))
		console.WriteRepeated(" ", spacesBetweenLineAndNumber)
		console.WriteLine(sourceText(d.source.LineAt(i)))
	}

	if d.position.Start == nil {
		return
	}
	labelLength := len(d.source.LineAt(d.position.Line))
	if d.position.End != nil {
		labelLength = *d.position.End - *d.position.Start
	}
	console.WriteRepeated(" ", sourceLineIndent)
	console.WriteRepeated(" ", maxLineNumberDigits)
	console.WriteRepeated(" ", spacesBetweenLineAndNumber)
	console.WriteRepeated(errorLabel("^"), labelLength)
	console.WriteLine(errorLabel(" at here"))
}

func digitsOf(n int) int {
	digits := 0
	for n > 0 {
		digits++
		n /= 10
	}
	return digits
}
