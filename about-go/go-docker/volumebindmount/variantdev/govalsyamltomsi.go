package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/sirupsen/logrus"

	// "gopkg.in/yaml.v3"
	// using 	"github.com/ghodss/yaml" for
	// "gopkg.in/yaml.v2"
	// "github.com/ghodss/yaml"
	"github.com/variantdev/vals"
)

func main() {

	// build a yaml file using kustomize from the dex folder.
	installString := "./dex"
	logrus.Info("instalString set to: ", installString)
	// build the yaml file as a command

	yamlcmd, err := exec.Command("kustomize", "build", installString).Output()
	if err != nil {
		logrus.Fatalf("yamlcmd failed with  %s\n", fmt.Sprint(err))
	}

	// convert the yaml object into a yaml byte for use in yaml.YAMLToJSON
	yamlStr := string(yamlcmd)

	logrus.Info("Turning yamlStr into temp file.")
	tempFile, err := os.Create("tmp/temp.yaml")
	logrus.Info("tempFile created: ", tempFile)

	if err != nil {
		logrus.Info(err)
	}

	// close file when done
	defer tempFile.Close()

	_, err2 := tempFile.WriteString(yamlStr)
	if err2 != nil {
		logrus.Info(err2)
	}

	// take in temp.yaml
	yamlMapStrIntf, err := vals.Input("tmp/temp.yaml")

	// delete the temp.yaml file after done
	// os.Remove("tmp/temp.yaml")

	// now the yamlMapStrIntf should have been created
	logrus.Info("yamlMapStrIntf: ", yamlMapStrIntf)
	logrus.Info("yamlMapStrIntf type is: ", reflect.TypeOf(yamlMapStrIntf).String())

	// secretsToCache := 256 // how many secrets to keep in LRU cache
	opts := vals.Options{}
	opts.CacheSize = 256
	opts.ExcludeSecret = false

	runtime, err := vals.New(opts)
	if err != nil {
		fmt.Println("Error: opts not working.")
	}

	// spit out map[string]interface{} replacing refs with values
	valsRendered, err := runtime.Eval(yamlMapStrIntf)

	// Convert map ( or rather, nested map[string]interface{} ) to json string
	jsonByte, err := json.Marshal(valsRendered)
	if err != nil {
		logrus.Info(err)
	}

	logrus.Info("the value of jsonStr is: ", string(jsonByte))

}
