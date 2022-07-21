package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println("value of foods: ", foods)
	fmt.Println("type of foods", reflect.TypeOf(foods))

}
