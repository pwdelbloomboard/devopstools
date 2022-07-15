package main

import (
	"fmt"
)

// using 	"github.com/ghodss/yaml" for

func main() {

	// creating a simple map[string]interface{}
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	fmt.Println(foods)

}
