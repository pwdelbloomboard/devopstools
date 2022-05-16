# General Go Tutorial

https://www.golangprograms.com/go-language.html

## Variables

## Constants

## Data Types

* {} Curly brackets on their own define lexical scope in Go. Meaning, whatever happens in that scope, stays in that scope. Variable declarations, calculations, etc.

### Struct

* [Go Structs - Go by Example](https://gobyexample.com/structs)

* You can set up a person struct type with fields for name, age, for example:

```
type person struct {
    name string
    age  int
}
```
* In order to populate the struct, you may need a function with a pointer:

```
func newPerson(name string) *person {
# You can safely return a pointer to local variable as a 
# local variable will survive the scope of the function.

    p := person{name: name}
    p.age = 42
    return &p
}
```
* Recall that with pointers, we're doing it for memory efficiency, using the stack rather than the heap.
* A simpler way to think about it is just that *person and &p means we're pointing memory addresses at each other in order to achieve this stack/heap step explicitly, rather than allow an interpretive language to figure that out for us.

We can then create structs with that function, or optionally just create them the non-memory efficient way:

```
// non-memory efficient way
person{"Bob", 20}

// memory efficient way
newPerson{"Jon"}

// printed out
{Bob 20}
&{Jon 42}

```
The fields within structs can be accessed with a dot.  E.g.:

```
s := person{name: "Sean", age: 50}
fmt.Println(s.name)
```

### Map Composite Literals

* [Map Composite Literals Documentation](https://go.dev/ref/spec#Composite_literals)

#### Subcategory of Map Composite Literals - map[string]struct



#### Subcategory of Map Composite Literals - map[string]interface{}


## Convert Types

## Operators

## If Else

## Switch Case

## For Loops

## Functions

## Variadic Functions

## Deferred Functions Calls

## Panic and Recover

## Arrays

## Slices

## Maps

## Struct

## Interface

Go Interfaces:

https://gobyexample.com/interfaces

## Goroutines

## Channels

## Concurrency Problems

## Logs

## Files and Directories

## Reading and Writing Files

## Regular Expression

## Find DNS records

## Cryptography

## Gotchas in Golang

## Import and Export

## Best Golang Packages

## Web Application