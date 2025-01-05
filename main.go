package main

import "github.com/mussel-lox/clam/diagnostic"

func main() {
	diagnostic.NewDiagnostic("don't say this!").
		At(19, 25).
		Attach(diagnostic.NewSource("<script>", "Several\r\nLines\nBefore\r\nHell\nWorld\r\nEndline")).
		Display()
}
