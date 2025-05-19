/*
Working with files
*/
package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println("Signature: ", sig)
}

func sha1Sum(fileName string) (string, error) {
	/*
		idiom: Acquire a resource, check for error, defer the release
		correct for files and mutexes for example
	*/
	// file here is a compressed version of the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close() // defers in a function are called in LIFO order.

	/*
		you need to decompress the file first before you can read it.
		NewReader takes an IO reader but the os.File implements the interface
		io.Reader interface is implemented by the os.File package since it has a method named
		Read.
	*/
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// Create a new writer from the sha1 package
	w := sha1.New()
	// copy the content from reader to writer
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	// Calculate the signature with Sum
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
