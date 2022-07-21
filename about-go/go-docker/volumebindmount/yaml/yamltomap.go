package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {

	// read the file
	yfile, err := ioutil.ReadFile("test.yaml")
	if err != nil {

		log.Fatal(err)
	}

	// make the map interface to hold the yaml
	data := make(map[interface{}]interface{})

	// unmarshal yfile test.yaml into the map interface location
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {

		log.Fatal(err2)
	}

	// range through data and display contents
	for k, v := range data {
		fmt.Printf("%s : %d\n", k, v)
	}
}
