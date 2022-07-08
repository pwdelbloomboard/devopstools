package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := map[string]string{
		"name": "Olivia",
	}

	// marshal map
	b, err := json.Marshal(m)
	if err != nil {
		panic("never to happen")
	}

	fmt.Println("the Marshalled value of b is: ", b)

}
