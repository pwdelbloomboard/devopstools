package main

import (
    "utils"
)

func main() {
    defer func() {
        util.Handle(func(err error) {
            // Handle errors in a generic way, 
            // for example using println, or writing to http 
        })
    }()

    var result, err := someFragileFunction()
    Check(err)
}

package utils

func Check(err error) {
    if err != nil {
        panic(err)
    }
}

func Handle(handler func(err error))  {
    if r := recover(); r != nil {
        if err, ok := r.(error); ok {
            handler(err)
        } else {
            panic(r)
        }
    }
}