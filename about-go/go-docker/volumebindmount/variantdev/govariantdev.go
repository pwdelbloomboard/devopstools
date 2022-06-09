package main

import (
	"github.com/variantdev/vals"
)

type Options struct {
	CacheSize     int
	ExcludeSecret bool
}

func main() {

	// secretsToCache := 256 // how many secrets to keep in LRU cache
	opts := Options{}
	opts.CacheSize = 256
	opts.ExcludeSecret = false

	runtime, err := vals.New(opts)
	if err != nil {
		return nil, err
	}

	valsRendered, err := runtime.Eval(map[string]interface{}{
		"inline": map[string]interface{}{
			"foo": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID[?region=us-west-1]",
			"bar": map[string]interface{}{
				"baz": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID[?region=us-west-1]",
			},
		},
	})
}
