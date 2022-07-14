## Getting Started - Attempting Stock Code

When attempting to first use variantdev, we get the following error when attempting to use the [stock code](https://github.com/variantdev/vals) provided by variantdev:

```
root@e876b9ee99bc:/home/volumebindmount/variantdev# go run govariantdev.go
# github.com/variantdev/vals/pkg/providers/awskms
/go/pkg/mod/github.com/variantdev/vals@v0.17.0/pkg/providers/awskms/awskms.go:5:2: imported and not used: "strings"
```

* This error is likely because of some kind of dependency on variantdev's side, they are importing a module which is not being used.
* This error can be cleared by back-dating the version of variatndev being used within the go.mod file:

```
$ cat go.mod
module example.com/m

go 1.17

require github.com/variantdev/vals v0.16.2
```

After this was completed and variantdev was run again, we found another error:

```
$ go run govariantdev.go
# command-line-arguments
./govariantdev.go:7:1: syntax error: non-declaration statement outside function body
note: module requires Go 1.17
```

Essentially, we need a function body and can't have some random declaration sitting there on its own, a-la:

```
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
```

So basically, a couple things:

* I wrapped the entity of everything above in a main(){} function, for that which followed import().
* After doing so, we have a few errors:

```
$ go run govariantdev.go
# command-line-arguments
./govariantdev.go:11:26: cannot use secretsToCache (type int) as type vals.Options in argument to vals.New
./govariantdev.go:13:3: too many arguments to return
	have (nil, error)
	want ()
note: module requires Go 1.17
```

* So we can't use this without using Go 1.17
* On line 11, we can't use an int under vals.New(), which is referring to: https://github.com/variantdev/vals/blob/d1a86060746caf40416b3bf96c7d27008d40e8ce/vals.go#L91
* This appears to take in opts, which must be of type Options:

```
// New returns an instance of Runtime
func New(opts Options) (*Runtime, error) {
```

Looking at type Options:

```
type Options struct {
	CacheSize     int
	ExcludeSecret bool
}
```
So we can create a new variable of type Options:

```
opts := Options(
    CacheSize   256
    ExcludeSecret   False
)
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

## Expansion on the Vals Package

* We have already been using the vals package, as shown above and within govariantdev.go within this directory.
* Documentation for any package can be found by appending the package's github after pkg.go.dev:

https://pkg.go.dev/github.com/variantdev/vals

## Sources

https://pkg.go.dev/github.com/variantdev/vals#readme-go

https://pkg.go.dev/github.com/variantdev/vals#section-readme

https://pkg.go.dev/github.com/variantdev/vals#Eval

https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L412

https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L113

