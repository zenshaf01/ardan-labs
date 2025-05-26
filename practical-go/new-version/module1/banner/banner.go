package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
}

/*
A string in go is a struct with 2 fields.
1. A length
2. Pointer to the underlying bytes
*/
func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2

	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))
}
