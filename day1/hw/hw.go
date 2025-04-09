/*
Every file should start with a package declaration
Each file lives inside a package
For this file we put it inside the main package
The main package is a special package is the main entry point package.
Every file in a folder should have the same package clause
*/
package main

/*
	Import necessary standard or third party packages
	Unused imports and variabled will be removed by go fmt
	Unsued imports and variables are a compile error
*/
import "fmt" // package for string formatting

// Main entry point for the program
func main() {
	/*
		Strings are unicode in Go
		Always precede the package name with the function name
		The package name is added here coz you can have a function name
		Println in your own package and it wont collide with teh fucntion in the fmt package.
		Functions and variables having a capital first letter are exported and can be imported into other modules.
		Small letter ones cannot.
	*/
	fmt.Println("Hello Gophers!") // Semicolons are optional

	/*
		You can pass command line arguments using the flags package
	*/
}

/*
You can run the above code from the terminal or from the IDE.
Commands:
	go build:
		This builds the executable. The executable contains not just the code you wrote
		but also contains the go runtime, the go scheduler etc.
		The executable is statically linked. This means it does not depend on any shared library
		on your local machine.

		You can also cross compile for other OS's with the `GOOS=darwin go build`
	go run:
		This runs your program without building an exe. This is faster.
	go env GOARCH GOOS:
		This returns the architecture of the computer and the OS of the local machine
*/
