# Go Maps

## Go Maps Overview

* Maps are basically what dicts are for python - they are associative arrays, meaning a key associates with a value.

## Map Breakdown

You can make an empty map with, "make"

```
m := make(map[string]int)
```

* If the name of the map is m
* The keys of the map are of type string
* The values of the map are of type int

## Creating an Entry in a Given Map

```
value = m["key"]
```

## Deleting an Entry in a Given Map

```
delete(m,"key")
```

## Checking if an Entry Exists

* Optional second return value to disambiguate between missing keys and keys with null or 0 value, or "".

```
	_, present := m["key1"]
	fmt.Println("present_or_not:", present)
```

## Declaring New Map in Same Line

* New maps with keys and values pre-created can be created in one-liners by adding 

```
newmap := map[string]int{"foo": 1, "bar": 2}
```
* Note: Curly brackets on their own define lexical scope in Go. Meaning, whatever happens in that scope, stays in that scope. Variable declarations, calculations, etc.



## Resources

* [Go by Example - Maps](https://gobyexample.com/maps)