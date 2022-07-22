package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	// "gopkg.in/yaml.v3"
	// using 	"github.com/ghodss/yaml" for
	// "gopkg.in/yaml.v2"
	// "github.com/ghodss/yaml"
)

func main() {

	// build a yaml file using kustomize from the dex folder.
	templateString := "./dex/config-map-template.yaml"
	logrus.Info("instalString set to: ", templateString)

	// the length of our replaceStringList
	const replaceStringListLen = 2

	// establish replaceStringList with ref's for each item.
	var replaceStringList [replaceStringListLen]string
	replaceStringList[0] = "clientID: ref+awsssm:///V1/kubeflow/local/GITLAB_APPLICATION_ID?region=us-west-1"
	replaceStringList[1] = "clientSecret: ref+awsssm:///V1/kubeflow/local/GITLAB_CLIENT_SECRET?region=us-west-1"

	// the length of our configcmdStdoutList
	const configcmdStdoutListLen = 2

	// the length of newly created *strings.Builder array to hold *strings.Builder items
	var configcmdStdoutList [configcmdStdoutListLen]*strings.Builder

	// cylce through replaceStringList
	for i := 0; i < replaceStringListLen; i++ {
		logrus.Info(replaceStringList[i])
		// go through replaceStringList
		// use vals as a command line tool to create file
		builtConfigByte, err := exec.Command("echo", replaceStringList[i]).Output()

		builtConfigString := string(builtConfigByte)

		// declare configcmd type

		// create a command object with any parameters we need to pass
		// "kubectl apply -f -" indicates the process standard input
		configcmd, configcmdStdout := exec.Command("vals", "eval", "-f", "-"), new(strings.Builder)
		// set the Stdin setup, connect to string file builtyamlstring
		configcmd.Stdin = strings.NewReader(builtConfigString)
		// log the applycmd
		logrus.Info("the applycmd is: ", configcmd)
		// set up the error so we get more error information on this commmand
		// var configout bytes.Buffer --> set configcmd.Stdout to &configout if pushing Stdout to buffer
		var configstderr bytes.Buffer
		configcmd.Stdout = configcmdStdout
		configcmd.Stderr = &configstderr
		// apply the apply command with .Run to apply the application to the cluster
		err = configcmd.Run()
		if err != nil {
			logrus.Fatalf("applycmd.Run() failed with  %s\n", fmt.Sprint(err)+": "+configstderr.String())
		}

		// output the result standard output stored in a string
		logrus.Info("configcmdStdout is: ", configcmdStdout)
		logrus.Info(reflect.TypeOf(configcmdStdout))
		logrus.Info(reflect.TypeOf(configcmd))

		configcmdStdoutList[i] = configcmdStdout
	}

	logrus.Info("configcmdStdoutList: ", configcmdStdoutList)
	/*

		yamlByte, err := ioutil.ReadFile(templateString)
		if err != nil {
			logrus.Fatal(err)
		}

			// convert the yaml object into a yaml byte for use in yaml.YAMLToJSON
			yamlStr := string(yamlByte)

				logrus.Info("Turning yamlStr into temp file.")
				tempFile, err := os.Create("tmp/temp.yaml")
				if err != nil {
					logrus.Info(err)
				}
				logrus.Info("tempFile created: ", tempFile)

				// close file when done
				defer tempFile.Close()

					// write yamlStr, the actual contents of the yaml, to the tempFile, tmp/temp.yaml
					_, err2 := tempFile.WriteString(yamlStr)
					if err2 != nil {
						logrus.Info(err2)
					}
	*/

}
