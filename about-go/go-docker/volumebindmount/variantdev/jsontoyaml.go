package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type vars struct {
	Data1 string  `json:"data1" yaml:"data1"`
	Data2 float64 `json:"data2" yaml:"data2"`
}

func main() {
	inp := []byte(`{"data1":"meow","data2":0.1234}`)

	data := &vars{}
	if err := json.Unmarshal(inp, data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output of input []byte converted to string: %s", inp)

	// Do some validation on data?

	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(".")
	fmt.Printf("output of Marshalled Yaml:\n%s\n", out)
}
