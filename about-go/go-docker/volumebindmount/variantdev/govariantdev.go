package main

import (
	"encoding/json"
	"fmt"

	"github.com/variantdev/vals"
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

	valsRendered, err := runtime.Eval(map[string]interface{}{
		"inline": map[string]interface{}{
			"foo": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			"bar": map[string]interface{}{
				"baz": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			},
		},
	})

	// Convert map to json string
	jsonStr, err := json.Marshal(valsRendered)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonStr))

}
