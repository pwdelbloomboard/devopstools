# Using env is GoLang

> To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.

```
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}
```

However, in order to work with the environmental variable, we have to keep in mind that by default the environmental variable is not exported.

Also, we need to take care of error handling, if it's Go, we can use LookupEnv() and use error handling, rather than GetEnv().

```
package main

import (
	"fmt"
	"os"
)

func main() {
	// get the env variable, set it to a go variable
	govar, ok := os.LookupEnv("WHATEVER")
	if !ok {
		fmt.Println("WHATEVER is not an env variable.")
	} else {
		fmt.Printf("WHATEVER is: ", govar)
	}
}
```

After building and running the above, and after exporting the following env:

```
export WHATEVER=hiya
```

We get as a result:

```
WHATEVER is: %!(EXTRA string=hiya)
```

When in reality we just wanted the end string, "hiya".

Basically, according to the docs on [fmt](https://pkg.go.dev/fmt#hdr-Format_errors) this is a type of format error.  

Go fmt has something called, ["verbs"](https://gobyexample.com/string-formatting) which are designed to format general Go values. These are sort of like, "flags," within the command line, they are added within the apostraphes of a string previous to a value to help format the value in a way that Golang expects.

So for example:

```
p := point{1, 2}

fmt.Printf("struct2: %+v\n", p)

fmt.Printf("type: %T\n", p)

```

For example, if using %+v\n above, it would include the struct's field names.  If using, %T\n, it would print outt the type of the value p.

So if we wanted to do a string (with escape characters included):

```
"\"string\"
```

We could use, %s\n as the verb.

So putting this into our code example:

```
func main() {
	// get the env variable, set it to a go variable
	govar, ok := os.LookupEnv("WHATEVER")
	if !ok {
		fmt.Println("WHATEVER is not an env variable.")
	} else {
		fmt.Printf("WHATEVER is: %s\n ", govar)
	}
```

In summary, the entirety of the code is:

```
package main

import (
	"fmt"
	"os"
)

func main() {
	// get the env variable, set it to a go variable
	govar, ok := os.LookupEnv("WHATEVER")
	if !ok {
		fmt.Println("WHATEVER is not an env variable.")
	} else {
		fmt.Printf("WHATEVER is: %s\n ", govar)
	}
}
```
