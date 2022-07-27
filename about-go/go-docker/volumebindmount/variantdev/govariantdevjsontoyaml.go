package main

import (
	"encoding/json"
	"fmt"

	"github.com/variantdev/vals"

	// using 	"github.com/ghodss/yaml" for
	"github.com/ghodss/yaml"
)

func main() {

	// secretsToCache := 256 // how many secrets to keep in LRU cache
	opts := vals.Options{}
	opts.CacheSize = 256
	opts.ExcludeSecret = false

	runtime, err := vals.New(opts)
	if err != nil {
		fmt.Println("Error: opts not working.")
	}

	// spit out map[string]interface{} replacing refs with values
	valsRendered, err := runtime.Eval(map[string]interface{}{
		"inline": map[string]interface{}{
			"foo": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			"bar": map[string]interface{}{
				"baz": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			},
		},
	})

	// Convert map ( or rather, nested map[string]interface{} ) to json string
	jsonStr, err := json.Marshal(valsRendered)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("the value of jsonStr is: ", jsonStr)

	fmt.Println("the value of string(jsonStr) is: ", string(jsonStr))

	// convert valsRendered, json object to yaml object.
	y, err := yaml.JSONToYAML(jsonStr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// show yaml which was converted using JSONToYAML
	fmt.Println("yaml version of json []byte{}, converted by JSONToYAML: ")

	// creating a convertable string
	yamlStr := string(y)

	fmt.Println(yamlStr)

	/* the below method does not work for what we're trying to do, Input appears to be more of an internal function
	// attempting to do this the output method
	// inputing the yaml as a map[string]interface{}
	// outputting a string
	outputString, err := Output(valsRendered)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(outputString)
	*/

}
