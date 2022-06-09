package main

import (
	"github.com/variantdev/vals"
)

secretsToCache := 256 // how many secrets to keep in LRU cache
runtime, err := vals.New(secretsToCache)
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