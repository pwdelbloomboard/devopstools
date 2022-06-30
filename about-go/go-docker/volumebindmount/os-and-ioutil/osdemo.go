// Golang program to tour through os
// use structs as map keys
package main

// importing required packages
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// the output of stdout will be a pointer such as &{0xc0000b6060}
	the_Stdout := os.Stdout
	fmt.Println("the_Stdout is: ", the_Stdout)

	//
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	fmt.Println("r is: ", r)
	fmt.Println("w is: ", w)

	//
	os.Stdout = w

	fmt.Println("Hello, playground") // this gets captured

	fmt.Println("os.Stdout is: ", os.Stdout) // this gets captured

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out)
}
