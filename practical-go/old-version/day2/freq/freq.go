package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
The questiuon we want to answer is:

What is the most common word in sherlock.txt.
Figure out the word frequency


Problems:
1. extract the text line by line
2. count words
*/

/*
Note that the function accepts an interface of io.Reader
This means that now this function can accept a file, a socket or anything that
implements the Reader interface. This makes the function much more flexible.

you should always process files line by line
*/

/*
This is how you create a regular expression
Note that the below line runs before main

The MustCompile function causes a panic if it fails.
Many other pakcages do this. This is because these are compiled once befoire the main
They just panic instead of returning an error

the “ used to define a regular expression is a raw string. in raw strings \ is just a character.
*/
var wordRe = regexp.MustCompile(`[a-zA-Z]`)
var rawString = `THIS IS A RAW STRING AND IT CAN HELP YOU DEFINE MULTI LINE STRINGS.
IN RAW STRINGS \ IS JUST A \.`

// This will also run before main
// func init() {

// }

func mapDemo() {
	var stocks map[string]float64 // symbol -> price
	sym := "TTWO"
	// Since map key values could be 0 by intent or by go's way of initializing it to its nil value
	// We should use comma ok , ok to determine if the value actually exists or not.
	if price, ok := stocks[sym]; ok {
		if ok {
			fmt.Println("%s -> $%.2f\n", sym, price)
		} else {
			fmt.Println("%s not found \n", sym)
		}
	}

	// You can initialize a map in two ways
	// stocks = make(map[string]float64)
	stocks = map[string]float64{
		sym:    137.73,
		"AAPL": 172.35,
	}
	// stocks[sym] = 136.73

	// You can interate over the map like so
	for key, val := range stocks {
		fmt.Println(key, " -> ", val)
	}

	if price, ok := stocks[sym]; ok {
		if ok {
			fmt.Printf("%s -> $%.2f\n", sym, price)
		} else {
			fmt.Printf("%s not found \n", sym)
		}
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "AAPL") // no panic on dleeteing a non existent key
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	// create a new scanner and pass it the file
	// bufio scanner also implements the same read method so you can pass the file as is
	s := bufio.NewScanner(r)
	lnum := 0
	freq := make(map[string]int) // word -> count
	// Read file line by line
	for s.Scan() { // Start scanning the file
		lnum++
		// line := s.Text() // this returns the current line
		/*
			We will use regular expression to find words ion text

			regular expressions allow us to find patterns in text

			-1 means all matches

			even numbers being part of the line are strings you would have to convert them to numbers
		*/
		// Extract text per line
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freq, nil
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

func main() {
	// Open file
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer file.Close()

	// file here is os.File and it can be passed into the function since it implements the Read function
	// And satisfy's the io.Reader interface
	word, err := mostCommon(file)
	if err != nil {
		fmt.Println("error encountered: ", err)
	}
	fmt.Println("Most common word is: ", word)
	// mapDemo()
}
