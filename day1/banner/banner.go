package main

import (
	"fmt"
	"unicode/utf8"
)

/*
* Go has full support for Unicode. All strings are unicode characters
* A single unicode code point / character is called a rune in go.
* A rune is created by ''
 */

func main() {
	banner("Go", 6)
	banner("G☺", 6) // With a smiley unicode emoji
	s := "G☺"
	/*
		* In Unicode we have a number for each character in every language in the world
		* The below line prints 4 coz of the unicode character used in the string.
		* The reason being the number which represents the smiley emoji cannot fit in 1 byte.
		* UTF-8 is an encoding scheme which has a map of what symbol has what unicode number
		* UTF-8 character, rune, unicode code point can be from 1 byte to 4 bytes.
		* So below line is 1 byte for the letter G and 3 bytes for the smiley face emoji.


		* The len function when passed a string returns the number of bytes the string is.
		* If you want to find the number of characters in a string you would have call

		*  In Go, You can look at string in two ways
			One being as a collection of bytes,
			Second being a collection of runes and characters.


		* Datatype:
			byte = uint8
			rune = int32
	*/
	fmt.Println("len: ", len(s)) // this prints 4. Why is that ?
	// range for loop
	for i, r := range s {
		// here r is giving you the rune / UTF8 code point
		fmt.Println(i, r)
		if i == 0 {
			// while using range, the character will be of rune type = int32
			fmt.Printf("%c of type %T\n", r, r)
		}
	}
	b := s[0] // The square bracket will give you 0th character as byte type = uint8
	fmt.Printf("%c of type %T\n", b, b)

	x, y := 1, "1"
	// example of printing variable values
	//
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Println("G: ", isPalindrome("G"))
}

/*
Function parameters:

  - All function parameters are passed by value.

  - This means that every function gets a copy of the arguments passed to that function.

  - Go types are concrete types.

  - In terms of strings go creates a copy of the pointer of the string and then passes that in to the function.

  - Strings are read only and are immutable.
*/
func banner(text string, width int) {
	// := means declare a variable and assign the value of the expression on the right
	// This variable is using type inference
	// padding := (width - len(text)) / 2 // This line has a bug
	padding := (width - utf8.RuneCountInString(text)) / 2

	//counting for loop
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(s string) bool {
	rs := []rune(s) // get slice of runes out of s. Get a slice of runs from the string
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}
