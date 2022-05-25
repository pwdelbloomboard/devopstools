package main

import "fmt"

// create a custom type using a struct full of individual primatives
// What we did here is to define a new type of name cities which are of type []string.
type Contactinfo struct {
	name        string
	age         int
	city, phone string
}

func main() {

	// Assign values to the fields, one by one
	var Shiju Contactinfo
	Shiju.name = "Shiju"
	Shiju.age = 35
	Shiju.city = "Longmont"
	Shiju.phone = "+1-940 033 72xx"

	// assign values to the fields in a struct
	Pingman := Contactinfo{
		name:  "Pingman",
		age:   92,
		city:  "San Antonio",
		phone: "+1-940 033 72xx",
	}

	// call the reciever function (method) on the instance of the custom type
	Shiju.printcontact
	Pingman.printcontact()

}

// Reciever function: in these functions
// a value and type are added in parenthesis before the function name is written.
// These could be comparable to “methods” in Matlab classes.
// note the input is the actual custom type
func (c Contactinfo) printcontact() {

	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", c.name, c.age, c.city, c.phone)

}
