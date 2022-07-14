# encoding/json module

### Refresher on Slices

* Arrays have fixed sizes. Slices however are dynamically-sized, flexible view into the elements of an array.

```
package main

import "fmt"

func main() {
	// create an array of 6 primes
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice the array "primes" from 1 to 4
	var slice []int = primes[1:4]

	// using array notation 0, 1, 2 ... print out the sliced selection
	fmt.Println(slice)
}

```
* The above program will print out [3 5 7]

### Built-In Types vs Custom Types

* GoLang has the following built-in types:

```
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
bool
string, []string
byte, []byte
rune
float32, float64
```
### Defining Custom Types

* Within structured data (e.g. JSON and using marshalling/unmarshalling), it can be helpful to define custom types to make funtion coupling easier.
* Basically you define custom types when you want to return structured data.
* Custom types are sort of like, "Methods/Classes" in Matlab, where a class essentially holds an object type with a pre-defined set of fields, and the pre-defined Methods are defined to operate on those objects and fill in the fields in a particular, idempotent way.

We may define a custom type with the below example:

```
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
```

* Basically, the custom type gets leveraged within the reciever function, so that the structure of the data is strict on the output, ensuring consistency.

### Defining Custom Types with Structs

```
type contactinfo struct {
	name string
	age int
	city,phone string
}
```

* Note that in our first example, we used ```type cities []string``` to define a type with one basic constraint, that there would be a string element.
* Multiple strings could be added to the type, as long as they were strings. ```Nepal := cities{"Kathmandu", "Pokhara", "Lumbini"}``` - 
* Now, we are requiring several elements (fields) within the type, it is a more complex custom type.
* In order to create instances of a struct type, we have to fill in all of the proper fields with values, which can be done by filling the values one by one.

```
var shiju Person
shiju.name="Shiju"
shiju.age=35
shiju.city="Longmont"
shiju.phone="+1-940 033 72xx"
```
* There is another way to fill the values, using a struct from the beginning (actually there are a couple ways):

```
	// assign values to the fields in a struct
	Pingman := Contactinfo{
		name:  "Pingman",
		age:   92,
		city:  "San Antonio",
		phone: "+1-940 033 72xx",
	}
```
However now in order to operate on this custom type struct, we need to create a method:

```
func (c Contactinfo) printcontact() {

	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", c.name, c.age, c.city, c.phone)

}
```
* Within that method, we specify the original object / type that we are referring to, "Contactinfo" as well as a placeholder or variable, c which is used to reference the fields of that object/ custom type.

* In this instance, we use Printf to refer to the various fields and print them out as shown.


### Defining Methods within Methods

* Methods can also be put within methods hierarchically to organize them further.




# Resources

* [Custom Types, Reciever Functions](https://medium.com/wesionary-team/slices-custom-types-and-receiver-functions-methods-in-golang-cdce4c01a5e8)
* [GoLang Type System](https://thenewstack.io/understanding-golang-type-system)
