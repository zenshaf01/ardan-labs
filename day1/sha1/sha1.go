/*
Working with files
*/
package main

import "os"

func sha1Sum(fileName string) (string, error) {
	/*
	*/
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	return "", nil
}
