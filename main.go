package main

import "github.com/mussel-lox/clam/internal/diagnostic"

func main() {
	diagnostic.NewDiagnostic("Hello, World").
		At(diagnostic.NewPosition(3, 0, 4)).
		Attach(diagnostic.NewSimpleSource("<script>", "Several\r\nLines\nBefore\r\nHello\nWorld")).
		Show()
}
