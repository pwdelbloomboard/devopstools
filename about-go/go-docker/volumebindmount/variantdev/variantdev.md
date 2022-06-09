## Getting Started - Attempting Stock Code

When attempting to first use variantdev, we get the following error when attempting to use the [stock code](https://github.com/variantdev/vals) provided by variantdev:

```
root@e876b9ee99bc:/home/volumebindmount/variantdev# go run govariantdev.go
# github.com/variantdev/vals/pkg/providers/awskms
/go/pkg/mod/github.com/variantdev/vals@v0.17.0/pkg/providers/awskms/awskms.go:5:2: imported and not used: "strings"
```

This in part may be because we're not even using an interface with a service that we're utilizing at all...e.g. vault. Instead we are using AWS.


```
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
        "foo": "ref+vault://127.0.0.1:8200/mykv/foo?proto=http#/mykey",
        "bar": map[string]interface{}{
            "baz": "ref+vault://127.0.0.1:8200/mykv/foo?proto=http#/mykey",
        },
    },
})
```

Looking further into the variantdev documentation:

```
There are four providers for AWS:

SSM Parameter Store
Secrets Manager
S3
KMS
```

* We are actually using SSM
* The example for how to use SSM shows: 

```
ref+awsssm://PATH/TO/PARAM[?region=REGION]
```

