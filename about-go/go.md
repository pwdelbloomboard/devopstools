# GoLang

### Basic Setup

* We can get started with GoLang using Docker and a [GoLang Docker image](https://hub.docker.com/_/golang).
* There are other distributions available found on the [GoLang Github](https://github.com/docker-library/golang/tree/dbdde931579e4a3d446b17167c67f573658d6989/1.17).

Using the dockerfile within this folder, we can set up Golang on a local environment running the Docker runtime with:

First, build the image:

```
docker build -t golang-bullseye .
```
Then, run the image, assigning it a name and port:
```
docker run -it --rm                                   \
-p 8883:8883                                          \
--name playwithgolang                                 \
golang-bullseye
```
You should then have access to a go interpreter / shell:

```
root@a631c2ab0aa8:/go# ls
```

This is the bash shell, starting out in the folder /go.

From here, you can start up a go file by installing nano.

```
apt-get update

apt-get -y install nano
```
Now, you can create a simple program:

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
Then, you can build, install packages and run the program with:

```
root@c2ec3ca8f4de:/go# go build packages.go
root@c2ec3ca8f4de:/go# go install packages.go
root@c2ec3ca8f4de:/go# go run packages.go
My favorite number is 1
```

#### Go Shell


### Setting Up Volume for Bind Mount

Let's say we want to keep some of the files we create within the GoLang shell on our local machine for safe keeping, to have for the next time we sign into practice Go.

To do this, we can set up a bind mount within a docker-compose.yml file (see the example attached).

We can then spin up the image and the service with:

```
docker-compose up -d
docker exec -t -i playwithgolang_container /bin/bash
```
Sure enough, after testing, we have a fully formed Go container and 

### Running Go with Port to View Externally

## Go Playground

* "Go Playground" is a web service that runs on GoLang.org's servers.

https://go.dev/tour/welcome/4
## About Go

[Tour Of Go](/about-go/tourofgo.md)

## Defining Variables

For example:
```
var (
appConfig   map[string]deploy.App
apps        []string
)
```
is the same as:

```
var appConfig map[string]deploy.App
var apps []string
```
And similar to:

```
errorMessage := range
```
Which doesn't care about the variable type, but declares the variable type based upon what the assigned variable type implicitly was.

## Go Build vs. Go Install

### For Packages

> go build   builds your package then discards the results.

> go install builds then installs the package in your $GOPATH/pkg directory

Basically, packages is a directory with some code files which expose different variables from a single point of reference, similar to the concept of a package on Node.

In an imaginary scenario, a single project may have thousands of dependency functions, some of which may have common behavior. The concept of a package allows you to put similar functions in the same directory. E.g. perhaps a function, "upperCase" and another one, "lowerCase" may be kept in a directory called, "case" which would be the package.

Every Go program must be part of some package. A standalone Go program must have package main declaration, even if there is nothing else going on.

If a program is part of the, "main" package, then "go install" will create a binary file, which upon execution calls the main function of the program.

Package names is the name of the directory stored in the src/ directory of packages where the GOPATH points to.  App is the package since app is the child directory of the src directory. Go install, "app" with the src/, e.g. src/app/ would be considered a package.

So basically, "go install <package>" looks for any file with the, "main" package directory and treats this as the entrypoint of the executable program, for which it needs to create a binary. If the file does not have, "main" it will create a package archive, "whatever.a" file inside of the /pkg directory, so you will have /pkg/darwin_amd64/app.a and then you would have /src/app/entry.go.

For packages, it's recommended to use plain and simple names, e.g. as though they were utilities, such as, "stringutils" or, "caseutil."  Underscores and hyphens should be avoided.

So essentially, "go build" is more temporary in nature to test something out while, "go install," actually builds the binaries.

### For Commands (package main)

> go build   builds the command and leaves the result in the current working directory.

> go install builds the command in a temporary directory then moves it to $GOPATH/bin


### More on Structuring Go Programs

Basically, you must always have one file which will be considered the, "main," file while the rest of the source code files (under src) draw into that main file as non-main packages which get imported.  Newly created packages, must be imported with the, ```import "whatever"``` section of the *.go file.

When you install a third party package, Go compiles the package and creates a package archive file. VSCode compiles the package immediately if you have the Go Plugin installed.


## & and * Pointers in Go

### About

* Pointers are a variable that store the memory address of another variable.
* Within Golang, they may also be called, "special variables."
* Memory addresses are hexidecimal in format, 0XFFAAF for example.

### Why are Pointers Needed?

* The concept of a variable in the first place is a name given to a memory location where the actual data is stored. To access the memory stored, we need to access the memory location where that data is stored. Variables in computing are not the same as variables in mathematics, though they appear to function that way to our human monkey-like brains.

* So whereas if we do:

```
var x int = 100
```

* This variable x will point to a memory address, for example 0x0201 which contains, "100" in the memory.

Whereas:

```
var y *int = &x
```

* The &x means that the actual value in the memory will be the memory address of x above, so it would put 0x0201 as the value.
* *int means that we are extracting that memory address 0x0201 out.

So in other words,

```
var x int=5748
var p *int
p=&x
```

* the first line is saying value stored in x is 5748
* the second line is saying address of x is 0x414020 (this is where the value 5748 is stored)
* the last line is showing that p has the value 0x414020, because we grabbed it using &x and put it into p.

### Where to Use Pointers

* Pointers don't make your application faster necessarily.
* Go is a garbage collected language. When you pass a pointer to a function, Go needs to perform, "Escape Analysis," to figure out if that variable should be stored in the heap or the stack.
* You can check what the escape analysis is doing by running, "go build -gcflags="-m"
* If the variable does not escape the heap, it lives on the stack.

Basically, you use pointers when you have structs which contain such large amounts of data, that you never want this data going on the heap because the slowness of working with that struct on the heap would make things slower overall because Golang is constantly doing garbage collection, which happens on the heap, and every time it has to do garbage collection on this big struct it takes a while. Instead, it's just faster to keep it on the nice stack where it doesn't need to be garbage collected.

## Data Types and Working with Data

### Slices

> In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure. Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice. It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array. Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. It is allowed to store duplicate elements in the slice. The first index position in a slice is always 0 and the last one will be (length of slice – 1).

* Slices are declared just like an array, but don't have the same size as arrays.

```
[]T
or 
[]T{}
or 
[]T{value1, value2, value3, ...value n}
```
In the example above, the T is the type of he elements, for example:

```
var my_slice[]int 
```

Slice has three main parts:

* Slice Pointer (pointer) - basically, the point on the original Array where the slice starts.
* Length - the lengh of the Slice.
* Capacity - the length from the Pointer to the end of the Array.

### Byte

* Bytes in Golang is an unsigned 8-bit integer which has type uint8. Bytes have a limit of 0-255 in numerical range and can represent ASCII characters.

Example:

```
var a1 byte = 97
```
So if you do:

```
fmt.Printf("%c\n", a1)
a
```
Basically, "97" represents the character, "a".

What is the reason for this?  Essentially, there is no, "char" character type in Go, instead, it uses the type, "byte" and something called, "rune" to represent character values.

What is rune?  In the past, there was only one character set, ASCII, which used 7 bits to represent 128 characters. Eventually this was expanded to include UTF-8 which is a superset of ASCII, and so on (UTF-16, UTF-32, etc.). In Golang, every single different UTF-8 (or otherwise Unicodely mapped since ASCII) character, including tabs, accented characters, emoji's, carriage return, has a Unicode Code Point, or in the Golang, known as a, "Rune."

The Rune type is an alias of int32.

* Strings are sequences of bytes and not of a Rune. Strings may contain Unicode text encoded in UTF-8, but the Go source code encodes it as UTF-8, there is no need to encode it in UTF-8.

Example:

```
♄
```
This is a Rune witha hexidecimal value ♄.

* This represents a Rune constant, where an integer value recognizes a Unicode code point. In Go language, a Rune Literal is expressed as one or more characters enclosed in single quotes like, 'g', '\t' in between single quotes.

If you did:

```
slc := []rune{♛}
```

then:

```
for i, value:= range slc{
	fmt.Printf("\Character: %c, Unicode:%U, Position:%d)
}
```

The output would be:

```
Character: ♛, Unicode:U+265B, Position:0 
```

Note that:

* %c is the actual character output, the ♛
* %U translates the output into unicode, U+XXXX
* %d is just the position within the rune itself, in this case the first position 0

So in summary, a Rune maps a Unicode set of characters to a character representation.

In python, this would have been accomplished with char. In python chr() converts an integer into a character directly, whereas ord() converts a character to an integer, and str() just represents the string itself. So in analogy:

* %c in Golang is kind of like char() in python
* %U in Golang is behind the scenes in python.


# Resources

* [Tour of Go](https://go.dev/tour/welcome/1)
* [Golang in Juputer Notebook](https://levelup.gitconnected.com/running-golang-on-jupyter-notebook-f7f9fba37812)
* [Everything You Need to Know About Packages in Go](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
* [Bytes in GoLang](https://zetcode.com/golang/byte/)
* [Where to Use Pointers in Go](https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac)
* [Slices in Golang](https://www.geeksforgeeks.org/slices-in-golang/)