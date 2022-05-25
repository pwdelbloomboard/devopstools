package main

import "fmt"

// create a custom type
// What we did here is to define a new type of name cities which are of type []string.
type cities []string

func main() {

	// What this does is initializes a variable “Nepal” of type cities which translates
	// to a string-based slice, and adds the provided value to the slice.
	Nepal := cities{"Kathmandu", "Pokhara", "Lumbini"}
	USA := cities{"New York", "Chicago", "Los Angeles", "Houston"}

	// call the reciever function (method) on the instance of the custom type
	Nepal.printcity()

	fmt.Println("")

	// call the reciever function (method) on the instance of the custom type
	USA.printcity()
}

// Reciever function: in these functions
// a value and type are added in parenthesis before the function name is written.
// These could be comparable to “methods” in Matlab classes.

func (c cities) printcity() {

	for i, city := range c {

		fmt.Println(i, city)

	}

}
