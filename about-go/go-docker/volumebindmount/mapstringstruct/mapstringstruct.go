// Golang program to show how to
// use structs as map keys
package main

// importing required packages
import "fmt"

//declaring a struct
type Person struct {
	Name    string
	City    string
	Pincode string
}

func main() {

	// Creating struct instances
	a1 := Person{"Pam", "Dehradun", "HiThere"}
	a2 := Person{Name: "Ram", City: "Delhi", Pincode: "OneTwo"}
	a3 := Person{Name: "Sam", City: "Lucknow", Pincode: "TwoOne"}

	// Declaring a map
	// var mp map[Person]int

	// Checking if the map is empty or not
	// if mp == nil {
	//	fmt.Println("True")
	//} else {
	//	fmt.Println("False")
	//}

	// Declaring and initialising
	// using map literals
	sample := map[string]Person
	fmt.Println(sample)
}
