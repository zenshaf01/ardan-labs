package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(saveDiv(1, 0))
}

/*
Handling panics
division by 0 will create a panic

note named return types. q and err here are local variables just like a and b.
named return values are mainly used win conjucntion with recover

error is also an interface
*/
func saveDiv(a, b int) (q int, err error) {

	/*
		Below is a defer call with an anonymous function. Dont forget to put the parenthesis () at the end.

		recover() is a built in function. Reoover returns a nil value if there is not panic, If there is a panic it will return
		the empty interface. so that means panioc can and will return any type.

		Below is how you recover from panics
	*/
	defer func() {
		// e's type is any. e is not an error type
		if e := recover(); e != nil {
			log.Println("Error: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
