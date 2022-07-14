// from https://stackoverflow.com/questions/17452722/how-to-get-the-key-value-from-a-json-string-in-go

package main

import (
	"encoding/json"
	"fmt"
)

// your JSON structure as a byte slice
var j = []byte(`{"foo":1,"bar":2,"baz":[3,4]}`)

func main() {

	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarshal JSON
	e := json.Unmarshal(j, &c)
	// print the unmarshaled json
	fmt.Println("e is: ", e)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		fmt.Println("s is: ", s)
		fmt.Println("c is: ", c)
		k[i] = s
		i++
	}
	// print the keys
	fmt.Println(k)

}
