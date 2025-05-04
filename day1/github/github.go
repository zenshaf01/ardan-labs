/*
	We can use go to perform http requests. Please do inspect http.cat
*/

package main

// Import the net/http package.
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// This is making a http request to the said url.
	// Note that the fucntion is returning 2 values. Go function can return multiple values.
	// resp is a pointer to the http.Response
	// err is an error returned in case of an http protocol error
	resp, err := http.Get("https://api.github.com/users/tebeka")
	// you should close the response body after you are done with it.
	// You can use defer to defer the closing of the response body

	// error are to be check immidiately
	// Go forces you to think about and handle errors in the main flow of the program
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.Fatalf is doing the following 2 things
			1. log.PrintF("error: %s", err)
			2. os.Exit(1)

			You can either use Fatalf or manually add the two lines. Up to you.
		*/
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
	defer resp.Body.Close()
	// Check content type
	/*
		The below line uses the Header package instead of maps because:
		1. HTTP headers can repeat
		2. Headers are case insensitive. Map keys are case sensitive
	*/
	fmt.Printf("Content-type: %s", resp.Header.Get("Content-Type"))

	// resp.Body

	// _ is a special variable in go. Since you cannot have unused variabvles in go, it is used as a placeholder to not having to use the value of the variable. It can be of any type.
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: cant copy %s", err)
	}

	/*
		Serialization: The process of converting your local data structure from any language (a struct in go, an object in java) into a sequence of bytes. This also called marshalling.
		Serialization is always used for sending and receiving information between computers. This is not required insiode your program. JSON is one format that is used for this.
		JSON is not the most efficient but is the most widly used.

		Deserialization: The process of converting the sequence of bytes into your languages data structure. Also referred to as unmarshalling.
	*/
}
