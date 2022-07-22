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

* In the case of vals, there are different libraries:

* CLI
* Helm
* Go

### Vals CLI

* vals is a helm-like configuration, "values" loader with support for various ssources and merge strategies.
* The command line can be used interactively, 

```
  eval		Evaluate a JSON/YAML document and replace any template expressions in it and prints the result
  exec		Populates the environment variables and executes the command
  env		Renders environment variables to be consumed by eval or a tool like direnv
  ksdecode	Decode YAML document(s) by converting Secret resources' "data" to "stringData" for use with "vals eval"
```

To install the executable, you would do:

```
git clone git@github.com:variantdev/vals.git
go build ./cmd/vals
```

* This gives you access to the actual command line interface.

### Vals Helm - Transforming Refs to Secrets

* Use value references as Helm Chart values, so that you can feed the helm template output to vals -f - for transforming the refs to secrets.
* So for example, for vault secrets, you can replace the string for your environmental variable with, "refs+vault://secret/data/foo#/mykey"

```
stringData:
  mysql-password: refs+vault://secret/data/foo#/mykey
  mysql-root-password: vZQmqdGw3z
```
* When you finally deploy the manifests, run vals eval to replace all the refs to actual secrets:

```
stringData:
    mysql-password: myvalue
    mysql-root-password: 0A8V1SER9t
```

* Finally run kubectl apply to apply manifests:

```
kubectl apply -f all.yaml
```
* So hypthetically, we could write a few different pre-written yaml files, which include references in them to the relevant secrets (depending upon whether local, staging, production), and then run:

```
vals eval
```
* ... to replace the refs with actual secrets.  Here is the command fully written out as a shell command:

```
cat manifests.yaml | ~/p/values/bin/vals eval -f - | tee all.yaml
```
* This all.yaml could then be applied to the cluster with, "kubectl apply".

### Vals Go

* We have already accomplished the use of Vals within Go within [this exercise](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/variantdev/govariantdev.go) which essentially extracts the values as strings within go that are json formatted.

* The key thing that we had established is that we can create a map[string]interface{} representation with the following:

```
valsRendered, err := runtime.Eval(map[string]interface{}{
    "inline": map[string]interface{}{
        "foo": "ref+vault://127.0.0.1:8200/mykv/foo?proto=http#/mykey",
        "bar": map[string]interface{}{
            "baz": "ref+vault://127.0.0.1:8200/mykv/foo?proto=http#/mykey",
        },
    },
})
```

* So hypothetically if we convert a static yaml file into a map[string]interface{} and use runtime.Eval(), the result will be valsRendered as a map[string]interface{} which includes the proper fields, (foo, bar) with the references converted to values. We could then unmarshall these values back into a json object (or hypothetically yaml object).

* Hypothetically, we could even Unmarhsal JSON to YAML with [this library](https://pkg.go.dev/github.com/ghodss/yaml#JSONToYAML)

* If we were working with a json string, as given in this file here, within our exercise cited above:

```
	// Convert map to json string
	jsonStr, err := json.Marshal(valsRendered)
	if err != nil {
		fmt.Println(err)
	}
```

* Then we would need to Unmarshal the JSON string, turn it into a YAML, then Marshal it back into a string.
* Or, we might Unmarshal json to a struct, and then yaml marshal it, as with:

[This repo's jsontoyaml.go we prooved out here](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/variantdev/jsontoyaml.go).

* In Summary:

* Replace references to desired ENV values with ref's in the desired input.yaml files.
* Turn this input.yaml into a map[string]interface{} input to runtime.Eval()
* Run runtime.Eval() to get valsRendered
* Do yaml.Marshal to on this a valsRendered.
* Turn that into a string output, then feed into kustomize and apply.

### Upshot of Above - Using Refs and Secrets

* If using Kustomize, it may be possible to just go in and do ref/secret substitution, but we have to evaluate how to do that by reading through the documentation.

#### Env Function - https://pkg.go.dev/github.com/variantdev/vals#Env

The code for Env is found [here](https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L379):

* The input is a map[string]interface{} called template
* the output is a []string string slice
* Env uses Eval() on the template 


#### Eval Function and runtime.Eval - https://pkg.go.dev/github.com/variantdev/vals#Eval

```
func Eval(template map[string]interface{}, o ...Options) (map[string]interface{}, error)
```
* Evaluate the any template expressions in a YAML document and print out the result (E.g. including the secret).
* the template input, a map[string]interface{} appears to mean template expressions, as mentioned in the Readme file, that Eval does: "Evaluate a JSON/YAML document and replace any template expressions in it and prints the result."
* Templates appear to be analogous to helm templates, e.g. locally rendered templtes, where you render a chart template and display the output as discussed [here](https://helm.sh/docs/helm/helm_template/#helm).


* So Eval appears to take in a map[string]interface{} and then output another map[string]interface{}

* Looking at the Eval coe, it appears to initiate the Options struct (from the Runtime), then if there are options available, e.g. o > 0, starts a New(opts) instance and then recursively returns Eval again on that template.

```
func Eval(template map[string]interface{}, o ...Options) (map[string]interface{}, error) {
	opts := Options{}
	if len(o) > 0 {
		opts = o[0]
	}
	runtime, err := New(opts)
	if err != nil {
		return nil, err
	}
	return runtime.Eval(template)
}
```

##### runtime.Eval

Looking further into runtime.Eval, we see:

```
// Eval replaces 'ref+<provider>://xxxxx' entries by their actual values
```

* ...with a much more complicated, longer function which can be found [here](https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L112).

* This is different than the (*Runtime) Eval function described below, which appears to set up the Runtime struct and Options struct.

#### func (*Runtime) Eval - https://pkg.go.dev/github.com/variantdev/vals#Runtime.Eval

* Runtime is a struct with various Options.  The purpose of Runtime appears to be to just set up the Runtime environment for the Vals library to be able to operate.

```
type Runtime struct {
	Options Options
	// contains filtered or unexported fields
}
```
* Eval replaces 'ref+<provider>://xxxxx' entries by their actual values

* The options are another struct:

```
type Options struct {
	CacheSize     int
	ExcludeSecret bool
}
```
* Within the Go framework, these options can be set as: 

```
opts := vals.Options{}
opts.CacheSize = 256/whatever
opts.ExcludeSecret = false/true
```

#### Exec Function - https://pkg.go.dev/github.com/variantdev/vals#Exec

* from the readme, it says: "Populates the environment variables and executes the command"

```
func Exec(template map[string]interface{}, args []string) error
```

* Takes in a map[string]interface{}, outputs a string list []string
* It also takes in args as a string list []string and gives no output.
* Looking at the code, it looks like it literally just executes the command (in the sense that a shell would execute a command).

```
func Exec(template map[string]interface{}, args []string) error {
	if len(args) == 0 {
		return errors.New("missing args")
	}
	env, err := Env(template)
	if err != nil {
		return err
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
```

* So what does this mean in this instance, to execute the command?
* Basically, this would seem to mean, to feed in whatever map[string]interface{} was fed in from runtime.Eval(), and to output values rendered within the specified place.  So using our example from [here](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/variantdev/govariantdev.go#L10):

```
	valsRendered, err := runtime.Eval(map[string]interface{}{
		"inline": map[string]interface{}{
			"foo": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			"bar": map[string]interface{}{
				"baz": "ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1",
			},
		},
	})
```
* If you think of, "inline" as the yaml file, having been converted to a map[string]interface{}, the references above will have been replaced with their values, and then output to another map[string]interface{}, ready for Exec() to execute on.
* Overall, variantdev/vals appears to be designed to work with a configuration file, which already exists in a cluster, with the following order of operations:

```
Use vals eval -f refs.yaml to replace all the refs in the file to actual values and secrets.

Use vals exec -f env.yaml -- <COMMAND> to populate envvars and execute the command.

Use vals env -f env.yaml to render envvars that are consumable by eval or a tool like direnv
```
* Basically, you can replace a shell script which is designed to work on an already operating deployment, config file with something like:

```
$ VAULT_TOKEN=yourtoken VAULT_ADDR=http://127.0.0.1:8200/ \
  echo "foo: ref+vault://secret/data/foo?proto=http#/mykey" | vals eval -f -
```

* Which replaces the item at the key in the yaml file with myvalue, having previously been "foo: ref+vault://secret/data/foo?proto=http#/mykey":

```
foo: myvalue
```
* Which would be like replacing the following script:

```
VAULT_TOKEN=yourtoken  VAULT_ADDR=http://127.0.0.1:8200/ cat <<EOF
foo: $(vault kv get -format json secret/foo | jq -r .data.data.mykey)
EOF
```
#### New Function - https://pkg.go.dev/github.com/variantdev/vals#New

* Returns an instance of the Runtime struct.

#### Input Function - https://pkg.go.dev/github.com/variantdev/vals#Input

* [Source Code](https://github.com/variantdev/vals/blob/v0.18.0/io.go#L65)

* Takes in a string, outputs a map[string]interface{}
* Does an unmarshal function on a yaml.
* This might potentially be a way to input a string from a yaml and output to the proper internal format used by variantdev.

There is a section in this function which makes it look like it's meant to deal with yaml, to unmarshal the yaml:

```
	if err := yaml.Unmarshal(input, &m); err != nil {
		return nil, err
	}

```

#### Output Function - 

* [Source Code](https://github.com/variantdev/vals/blob/ead3b58e118841ec33c33a29483ba3539b632aa4/io.go#L65)

* Appears to take a *string and an interface as an input and output a string.
* Likely this is taking the input either JSON or YAML and is marshalling it (to be able to display or print).
* Again, outputs a string, likely a string version of JSON or YAML.

## Pulling It All Together

So the pathway to pull everything together in the way we need it to be done is as follows:

0. Create a yaml file with the desired ref's pointing to where secrets are kept.  We did this with [the "dex" folder here](https://github.com/pwdelbloomboard/devopstools/tree/main/about-go/go-docker/volumebindmount/variantdev/dex).

1. Pull a yaml file as a string from an arbitrary URL/location.  This can be done with kustomize and [this example code on these lines](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/variantdev/govalsyamltomsi.go#L26).

* A) Better yet, build the config-map-template.yaml first, replacing the appropriate values, then writing the results of that to config-map.yaml with the replaced values.

2. Convert said yaml file string into a map[string]interface{}.  This means: 

* A) First building the yaml file using Kustomize with ```yamlcmd, err := exec.Command("kustomize", "build", installString).Output()```
* B) Next converting that command to a string, ```yamlStr := string(yamlcmd)```
* C) The tricky part, create a temp.yaml file into which you write the yamlStr with tempFile.WriteString(yamlStr).
* D) Using vals.Input("tmp/temp.yaml"), write to a yamlMapStrIntf to finally have the finalized map string interface.
* E) Range over the temp.yaml to ensure every line is copied.
* F) Very important, remove the temporary file for security purposes.

* 2A. Since yaml files unmarshalling only marshals one section seperated by "---" at one time, we have to decode the yaml file, as we have done [here in this example](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/yaml/yamltodoc.go).  This iterative method is necessary because if we attempt to unmarshal a yaml into a struct, straight up without any iteration, only the first section of the yaml file gets unmarshalled, as with [this example converting a yaml to a map](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/yaml/yamltomap.go).

* 


3. Use the [example here](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/variantdev/govariantdevjsontoyaml.go) to pull said rendered values, and to then convert that to JSON, and then into YAML, and a YAML string.

4. Said YAML string can be used to apply to create a configmap on a deployment.


## Appendix

### Alternative Methods of "Pulling It All Together" Which Don't Work

#### Directly Converting a Map String Interface from Yaml - Tutorial Methods

* Allegedly, in order to convert a yaml to a map[string]interface{}, there is a pathway to do that using json as described in this tutorial here by [Bitfield Consulting](https://bitfieldconsulting.com/golang/map-string-interface).

* So in order to achieve step 2) above, we could have attempted to do the following:

1. Convert YAML to JSON using the "github.com/ghodss/yaml" library and the YAMLtoJSON function.
2. Use json.Unmarshal to unmarshal the data from json to map[string]interface{} -- in the below example, p should have the data given in the variable, "data" which was originally in JSON.

```
var p map[string]interface{}
err = json.Unmarshal(data, &p)
```
* However, doing this does not accomplish what we want, because it simply creates a map[string]interface{} without having a structure for which to put the yaml fields in. A struct must be built which is custom to this yaml in order to inject values into the map[string]interface{}.

#### Turning []byte into a JSON Object

* Note ... An alternative method for step 2C) of, "Pulling it All Together" above, which sounds reasonable, but does not work would be to do the following:

* C) Turning that []byte into a JSON Object, which is ultimately a map[string]interface{}.


2A. We can attempt to do this by using the ```jsonObj, err := yaml.YAMLToJSON(yamlcmd)``` which should take the input yaml command.  This comes from the library [ghodss/yaml](https://pkg.go.dev/github.com/ghodss/yaml).  Specifically, this [YAMLToJSON command documented here](https://pkg.go.dev/github.com/ghodss/yaml#YAMLToJSON) takes an input of a []byte (byte array), and outputs a []byte (byte array), so we have to convert from string to []byte.

```
func YAMLToJSON(y []byte) ([]byte, error)
```

2B. When we attempt to convert the yaml file as a JSON object back into a string, we only see one line as an output: 

```
{"apiVersion":"v1","kind":"Namespace","metadata":{"name":"auth"}}
```

Also of note, that above function, YAMLToJSON outputs a []byte, when really we need a map[string]interface{}.


#### Turning the Output of Kustomize into a String, Then Byte Array

According to [this blog post](https://www.ribice.ba/golang-yaml-string-map/), the default yaml library marshals maps into map[string]interface{}.  Looking at the Marshal/Unmarshal function from the Yaml Library:

```
func Unmarshal(in []byte, out interface{}) (err error)     
func Marshal(in interface{}) (out []byte, err error)
```
* Note, there is also an, "Unmarshalstrict" option.
* So when we attempt to Unmarshal the yaml byte, we get:

```
	err = yaml.Unmarshal(yamlByte, &t)
	if err != nil {
		logrus.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
```

* Which results in essentially an empty unmarsahlled :

```
--- t:
{ {0 []}}
```
* This is likely because our struct is not formatted correctly, e.g. that the struct is pre-defined toward fields that do not match up with the fields coming from the YAML, e.g.:

```
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}
```
* So at this point, it's good to ask the question ... is there anything in variantdev/vals that might already do what we're trying to do?

#### Attempting to Feed a yaml as a String Directly Into vals.Input()

* There's the Input Function:

```
func Input(f string) (map[string]interface{}, error)
```

* This evidently inputs a string and outputs a map[string]interface{}. Does it work with our yamlStr?
* Evidently, we hit an input error upon attempting the conversion, telling us that the, "file name is too long."  Does this really mean that the file name is too long?

```
	yamlMapStrIntf, err := vals.Input(yamlStr)

	logrus.Info("Banana!  vals.Input command complete.")
	if err != nil {
		logrus.Info("Double Banana! We hit an error!")
		logrus.Info(err)
	}

```
* Looking at the actual function for Input, we see:

```
func Input(f string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	var input []byte
	var err error
	if f == "-" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else if f != "" {
		input, err = ioutil.ReadFile(f)
	} else {
		return nil, fmt.Errorf("Nothing to eval: No file specified")
	}
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(input, &m); err != nil {
		return nil, err
	}

	return m, nil
}
```
* It appears that Input uses the utility, ioutil.ReadAll() on the os.Stdin with the input string as "-" or if no input string, uses ioutil.ReadFile(f).
* What seems to be happening is that ioutil.Readfile() is being activated on yamlStr, which is leading to an error.
* Looking at the documentation for ioutil.ReadFile() it appears that as of Go version 1.16, ReadFile() simply calls a filename from os.ReadFile, e.g.:

```
func main() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}
```
* So, we have to figure out how to feed in the string as a Stdin.
* We know that using ```vals.Input("-")``` will basically make the vals.Input listen for the Stdin. However if we run this, and then never run the command itself, we never get an output.
* Looking at the actual [ioutil source code](https://pkg.go.dev/io/ioutil#ReadAll):

```
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
```
* Basically, it appears that we just have to feed the results of strings.NewReader() into ioutil.ReadAll() in order to get it to read and create a result.
* However, even this method results in an empty map[string]interface{}

#### Creating a Temp File Using os.CreateTemp

* Basically, if we attempt to use os.CreateTemp to create an actual file which gets fed into the vals.Input function, "othertemp.yaml", it seems that it would work, but the result is an empty map[string]interface{}

```
	// create temp file using os.CreateTemp
	// map[string]interface{} result is empty if using this method
	/*
		otherTempFile, err := os.CreateTemp("", "othertemp.yaml")
		if err != nil {
			logrus.Fatal("Cannot create temporary file", err)
		}

		defer otherTempFile.Close()
		defer os.Remove(otherTempFile.Name())

		yamlStrByteArray := []byte(yamlStr)
		if _, err = otherTempFile.Write(yamlStrByteArray); err != nil {
			log.Fatal("Failed to write to temporary file", err)
		}
	*/
```

#### Similar to Above, Creating a Temp File using ioutil.TempFile

* This also results in an empty map[string]interface{} output from the vals.Input() function.

```
	/*
		// Create our Temp File:  This will create a filename like /tmp/prefix-123456
		// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
		tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
		if err != nil {
			log.Fatal("Cannot create temporary file", err)
		}

		// Remember to clean up the file afterwards
		defer os.Remove(tmpFile.Name())

		logrus.Info("Created File: " + tmpFile.Name())

		// Example writing to the file
		yamlStrByteArray := []byte(yamlStr)
		if _, err = tmpFile.Write(yamlStrByteArray); err != nil {
			log.Fatal("Failed to write to temporary file", err)
		}
	*/
```

#### Attempting to Pass a String as an Input to exec.Command

* Basically, we can try to use some fancy methodology to pass stdin to a command, and then run vals.Input("-") given the vals.Input() source code looking for a "-" input to signify that a stdin input is coming in:

```
	// pass a string as input to exec.Command
	// yamlStdin := strings.NewReader(yamlStr)

	/*
		// use a pipe
		pipeInput, pipeOutput, err := os.Pipe()
		if err != nil {
			logrus.Info(err)
		}

		logrus.Info("pipeOutput", pipeOutput)

		cmd := exec.Command("echo", yamlStr)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// run the command, put into output as cmd.Run()
		cmderr := cmd.Run()
		if cmderr != nil {
			logrus.Info(cmderr)
		}
	*/
```



## Sources

https://pkg.go.dev/github.com/variantdev/vals#readme-go

https://pkg.go.dev/github.com/variantdev/vals#section-readme

https://pkg.go.dev/github.com/variantdev/vals#Eval

https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L412

https://github.com/variantdev/vals/blob/v0.18.0/vals.go#L113

