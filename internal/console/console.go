package console

import (
	"bufio"
	"fmt"
	"os"
)

var bufferedConsoleWriter = bufio.NewWriter(os.Stdout)

func Write(a ...any) {
	if _, err := fmt.Fprint(bufferedConsoleWriter, a[0]); err != nil {
		panic(err)
	}

	for i := 1; i < len(a); i++ {
		if _, err := fmt.Fprint(bufferedConsoleWriter, " "); err != nil {
			panic(err)
		}
		if _, err := fmt.Fprint(bufferedConsoleWriter, a[i]); err != nil {
			panic(err)
		}
	}
}

func WriteLine(a ...any) { Write(a...); Write("\n") }

func Writef(format string, a ...any) {
	if _, err := fmt.Fprintf(bufferedConsoleWriter, format, a...); err != nil {
		panic(err)
	}
}

func WriteRepeated(a any, repeat int) {
	for range repeat {
		Write(a)
	}
}

func Flush() {
	if err := bufferedConsoleWriter.Flush(); err != nil {
		panic(err)
	}
}

func Print(a ...any)                  { Write(a...); Flush() }
func Println(a ...any)                { WriteLine(a...); Flush() }
func Printf(format string, a ...any)  { Writef(format, a...); Flush() }
func PrintRepeated(a any, repeat int) { WriteRepeated(a, repeat); Flush() }
