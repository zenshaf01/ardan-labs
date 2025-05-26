package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func SHA1Sig(fileName string) (string, error) {
	//Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		// unzip the file
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
}
