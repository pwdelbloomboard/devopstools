package main

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// using 	"github.com/ghodss/yaml" for

func main() {

	// creating a simple map[string]interface{}
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	fmt.Println(foods)

	// build a yaml file using kustomize from the dex folder.
	installString := "./dex"
	logrus.Info("instalString set to: ", installString)
	// build the yaml file as a command

	yamlcmd, err := exec.Command("kustomize", "build", installString).Output()
	if err != nil {
		logrus.Fatalf("yamlcmd failed with  %s\n", fmt.Sprint(err))
	}

	// build the yaml as a string
	yamlstring = string(yamlcmd)
	logrus.Info(yamlstring)

}
