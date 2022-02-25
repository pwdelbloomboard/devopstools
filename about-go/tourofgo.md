
### Introductory Code

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
#### Packages

* Programs start running in the package ```main```
* fmt, math/rand

####  Imports

Packages can be used in a Printf - factorized, import statment with %g:

```
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```
vs. a regular package, Printline statement:

```
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

#### Exported Names

Names are exported if they begin with a Capital letter.

```
func main() {
	fmt.Println(math.pi)
}
```
...results in an error, "/prog.go:9:14: cannot refer to unexported name math.pi" vs.:

```
func main() {
	fmt.Println(math.Pi)
}
```

#### Function Declarations

type comes AFTER the variable name:

```
func add(x int, y int) int {
	return x + y
}
```

#### Omit Types if Same and Last Contains

* When two or more consecutive named functions parameters share a type, you can omit the type from all but the last.

like so:

```
func add(x, y int) int {
	return x + y
}
```

Note that only "y int" is declared, not both "x int" and "y int" .


#### Functions can Return Any Number of Results

* Example: the swap function returns two strings

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

The above program will return, "world hello" - because we literally swaped the strings, "hello", "world" and the ouput of the swap included "return x, y" (two strings)

.

#### Named Return Values

Go's return values may be named.

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

Note the values calculated, the x and y are returned as values with a space in between, without having named any value after, "return."  This is known as a naked return.

#### Variables

The var statement declares a list of variables, as with a function argument list.

```
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

The above will return an int i as 0, then c, python, java as booleans showing that they are false.

#### Variables with initializers


```
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```
#### Short Variable Declarations



https://go.dev/tour/basics/9


