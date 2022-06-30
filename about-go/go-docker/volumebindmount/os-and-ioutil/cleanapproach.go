package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Your function
// some_function(as a go-routine) will write a string into pipe writer
func some_function(w *io.PipeWriter) {
	defer w.Close()
	// Fill pipe writer
	fmt.Fprintln(w, "Hello World")
}

// main function
func main() {
	// create a pipe reader and writer
	// if you write something into the pipe writer, it will be copied to pipe reader by go
	pr, pw := io.Pipe()

	// pass writer to function, even though nothing has been created yet
	go some_function(pw)

	// Use buffer b
	// Create a MultiWriter with os.Stdout and custom buffer b
	// custom buffer to get standard output of function
	var b bytes.Buffer

	// Multiwriter is similar to the linux Tee command, so as os.Stdout recieves the input it gets duplicated to &b
	// os.Stdout and variable byte buffer `b`
	// os.Stdout will receive the output as well as your custom buffer b
	mw := io.MultiWriter(os.Stdout, &b)

	// io.Copy will then copy content from pipe-reader (pr) into multi-writer (mw)
	// copies pipe reader content to standard output & custom buffer
	_, err := io.Copy(mw, pr)

	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	// use variable, which was stored in the buffer, b
	fmt.Println(b.String())
}
