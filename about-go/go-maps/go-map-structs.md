## Golang program that uses structs as map keys

* A map in golang is like a dict.
* You can create a map dynamically, setting up map keys with a struct.

```
map[Key_Type]Value_Type{}
```

* So first you can set up the struct:

```
//declaring a struct
type Person struct {
    Name    string
    City    string
    Pincode int
}
```

Then with that Person struct, having Names, Cities and Pincodes, you can create several entries.
```
a1 = Person{Name: "Ram", City: "Delhi", Pincode: 2400}
```
Declaring the map as a variable:

```
var m map[Person]int
```
Or alternatively, assigning the map dynamically on the fly:

```
example_map := map[Person]int{a1: 1, a2: 2, a3: 3}
```

This essentially assigns each entry, a1, a2, a3 at different points in the map, 1, 2, 3.

```
map[{{Ram Delhi 2400}:1 Pam Dehradun 2200}:2 {Sam Lucknow 1070}:3]
```

* For an even more clear example, we can create a mapstringstruct and assign values:

https://stackoverflow.com/questions/31761391/create-mapstringstruct-and-assign-a-value

These are also called, "map composite literals."