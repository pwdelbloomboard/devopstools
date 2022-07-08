package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// attempt to open non-existing file
	// err != nil
	if _, err := os.Open("non-existing"); err != nil {
		// if err matches target fs.ErrNotExist
		// Is.() is matching the target of err to the second argument, Is(err,target error) bool
		// has a bool output
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}
