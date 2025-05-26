package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

/*
	This program should kill the process if given an active process id
*/

func KillServer(pidFile string) error {
	// Open the file
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	/*
		- Defer will always run when the function exits. Will also run if there is a panic
		- Defer works at the function level. Dont put it inside a loop
		- Defer's are executed in reverse order if you have more than one defer. (LIFO, stack)
		- Idiom: Acquire resource, check for error and defer release
	*/
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", pidFile, err)
	}

	slog.Info("killing", "pid", pid)
	if err := os.Remove(pidFile); err != nil {
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}

func main() {
	err := KillServer("server.pid")
	if err != nil {
		fmt.Println("Error: ", err)

		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}

		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s \n", e)
		}
	}
}
