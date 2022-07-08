package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// attempt to open non-existing file
	if _, err := os.Open("non-existing"); err != nil {
		// declare a variable of type pathError
		// https://pkg.go.dev/os#PathError
		// this is from os package, it records an error and the operation and file path that caused it.
		var pathError *fs.PathError
		// the errors.As() function looks at error and matches the pattern, in this case, &pathError
		if errors.As(err, &pathError) {
			// then gives the following message, "failed at path" appended rather than just the pathError.Path signal
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
