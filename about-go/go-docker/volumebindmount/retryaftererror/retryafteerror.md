# Implementing the Capability to Retry a Function After an Exit Error Code in Golang

* There is an interesting [Stackoverflow Article on the Topic](https://stackoverflow.com/questions/67069723/keep-retrying-a-function-in-golang). Basically, in a situation where a job needs to be continuous, or perhaps the situation when installing, or rather, "applying," (e.g. using kubectl apply) an application to a pod in Kubernetes, prior to a pod being ready, necessitating re-installation in order for the, "kubectl apply" to work, the need arises to be able to utilize Go Run such that it does not fail and exit the Go runtime upon recieving an exit error, at least in this exact instance, or at least up toa critical count of a particular number of times.

### Packages

#### errors - https://pkg.go.dev/errors

> Package errors implements functions to manipulate errors.  The New function creates errors whose only content is a text message.

> The New function creates errors whose only content is a text message.

##### Further Background on Errors

###### Unwrapping Errors

* There is a concept of, "unwrapping and wrapping" errors.
* This means adding extra context information to the returned error, like the name of the function where the error occured, the cause, the type, etc.  This can be used to create more clear error messages.
* So for example, you ca set up an error with a struct that contains, "When" and "What" with the timestamp and the error string respectively.
* If you create a function:

```
func (e MyError) Error() string {
    return fmt.Sprintf("%v: %v", e.When, e.What)
}
```

* This function may be called as a return, e.g.:

```
func oops() error {
    return MyError{
        time.Date(whatever),
        "Error Message: The file system had gone away",
    }
}
```

Then, when you call the function which calls the error function above, the e.When and e.What can be called respectively as a way to create structure or, "fields" of the error which can be called upon or not, with the outside function, oops() filling in those details respectively ... kind of like an error class or structured error with particular detailed fields.


```
package main

import (
	"errors"
	"fmt"
)

func main() {
    // create a enew error, error1
	err1 := errors.New("error1")
    // set err2 to fmt.Errorf, error format, which formats according to a 
    // format specifier and returns the string as a value that satisfies an error
    // basically uses formatting features to create descriptive error message
    // https://pkg.go.dev/fmt#Errorf
	err2 := fmt.Errorf("error2: [%w]", err1)
    // print err2
	fmt.Println(err2)
    // Unwrap returns the result of calling the Unwrap method or err, otherwise returns nil.
	fmt.Println(errors.Unwrap(err2))
	// Output
	// error2: [error1]
	// error1
}
```

###### As

* As is used to match a target within an error, with the format:

```
func As(err, target error) bool
```

###### Is

* Is is used to match a target within an error, with the format:

```
func Is(err, target error) bool
```


###### New

* There's also, "New"


#### log - https://pkg.go.dev/log

### Definitions in GoLang

* err error

* f func()

* func() error
