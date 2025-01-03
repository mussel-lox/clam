package console

import (
	"bufio"
	"fmt"
	"os"
)

var bufferedConsoleWriter = bufio.NewWriter(os.Stdout)

func Write(a ...any) {
	_, err := bufferedConsoleWriter.WriteString(fmt.Sprint(a[0]))
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(a); i++ {
		_, err = bufferedConsoleWriter.WriteString(" ")
		if err != nil {
			panic(err)
		}

		_, err = bufferedConsoleWriter.WriteString(fmt.Sprint(a[i]))
		if err != nil {
			panic(err)
		}
	}
}

func WriteLine(a ...any) { Write(a...); Write("\n") }

func Writef(format string, a ...any) {
	_, err := bufferedConsoleWriter.WriteString(fmt.Sprintf(format, a...))
	if err != nil {
		panic(err)
	}
}

func WriteRepeated(a any, repeat int) {
	for range repeat {
		Write(a)
	}
}

func Flush() {
	err := bufferedConsoleWriter.Flush()
	if err != nil {
		panic(err)
	}
}

func Print(a ...any)                  { Write(a...); Flush() }
func Println(a ...any)                { WriteLine(a...); Flush() }
func Printf(format string, a ...any)  { Writef(format, a...); Flush() }
func PrintRepeated(a any, repeat int) { WriteRepeated(a, repeat); Flush() }
