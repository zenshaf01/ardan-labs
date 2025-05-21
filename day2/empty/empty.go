package main

import "fmt"

func main() {
	// This is an empty interface. Any means i can become any type.
	// DONT USE empty interfaces or any
	var i any

	i = 7
	fmt.Println("i can be anything 1: ", i)

	i = "Seven"
	fmt.Println("i can be anything 2: ", i)
}
