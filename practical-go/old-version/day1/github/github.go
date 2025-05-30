/*
	We can use go to perform http requests. Please do inspect http.cat
*/

package main

// Import the net/http package.
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	name, numRepos, err := getGithubInfo("tebeka")
	if err != nil {
		log.Fatalf("error: unable to get github info %s", err)
	}
	fmt.Println(name, " Has ", numRepos, " repositories")
}

/*
To decode JSON to Go struct we need a struct we need the JSON fields to be decoded to.

Each field thjat need to be mapped ion to this struct from the JSON object needs to be exported.
Meaning they need to start with a capital letter.
*/
type Reply struct {
	Name string
	// The type of int here is telling go that whenever you see a number in a json cast it to an int.
	// Similarly if you had put float64 here, Go would cast it to a float 64 when the JSON fields get mapped from JSON to Go.
	Public_Repos int
	// This is called a field tags, this can be useful if the struct field name is differnt from the json field
	NumRepos int `json:"public_repos"`
}

/*
The function should return the user name and the num of repos
*/
func getGithubInfo(login string) (string, int, error) {
	// This is making a http request to the said url.
	// Note that the fucntion is returning 2 values. Go function can return multiple values.
	// resp is a pointer to the http.Response
	// err is an error returned in case of an http protocol error
	// PathEscape makes sure it escapes invalid characters in the url path like + sign or other characters which might break the url
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	// you should close the response body after you are done with it.
	// You can use defer to defer the closing of the response body

	// error are to be check immidiately
	// Go forces you to think about and handle errors in the main flow of the program
	if err != nil {
		return "", 0, err
		// log.Fatalf("error: %s", err)
		/*
			log.Fatalf is doing the following 2 things
			1. log.PrintF("error: %s", err)
			2. os.Exit(1)

			You can either use Fatalf or manually add the two lines. Up to you.
		*/
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		// log.Fatalf("error: %s", resp.Status)
		return "", 0, fmt.Errorf("%v - %s", url, resp.Status)
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
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: cant copy %s", err)
	// }

	/*
		Serialization: The process of converting your local data structure from any language (a struct in go, an object in java) into a sequence of bytes. This also called marshalling.
		Serialization is always used for sending and receiving information between computers. This is not required insiode your program. JSON is one format that is used for this.
		JSON is not the most efficient but is the most widly used.

		Deserialization: The process of converting the sequence of bytes into your languages data structure. Also referred to as unmarshalling.

		type mapping from json to go:
		JSON: 					Go:
		----					--
		true / false 	<->		true / false
		string			<->		string
		null			<->		nil
		number 			<-> 	float64. but you can also map to float32, int8, int16, int32, int64, uint8 ...
		array 			<-> 	[]T (slice of type T) ([]any) ([]interface{}) slice of []any can hold an array of mixed types.
		object			<->		map[string]any, struct

		any is the empty interface. But what is the empty interface.
		any is 1.18+

		Encoding / Decoding JSON <-> GO:
		--------------------------------
		JSON -> io.Reader -> Go: json.Decoder
		JSON -> []byte -> Go: json.Unmashal
		Go -> io.Writer -> JSON: json.Encoder
		Go -> []byte -> JSON: json.Marshal

	*/

	// Intialize a variable with type of reply
	// var r Reply

	// This is an anonymous struct. anonymous structs dont have a type.
	// The anonymus struct does not need a type. Use it in places where you need a one off struct which wont be used anywhere else
	var r struct {
		Name string
		// The type of int here is telling go that whenever you see a number in a json cast it to an int.
		// Similarly if you had put float64 here, Go would cast it to a float 64 when the JSON fields get mapped from JSON to Go.
		Public_Repos int
		// This is called a field tags, this can be useful if the struct field name is differnt from the json field
		NumRepos int `json:"public_repos"`
	}
	// Create a new decoder
	dec := json.NewDecoder(resp.Body)
	// Try to map the json fields in the reply struct. We must pass a pointer to the reply variable
	// We pass a pointer because we want our r variable to get updated with the json fields.
	// Since go is pass by value, if we passed the value of r instead of the pointer to it, Go would make a copy of r and the json
	// map will not get reflected to us in our main function here.
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}
