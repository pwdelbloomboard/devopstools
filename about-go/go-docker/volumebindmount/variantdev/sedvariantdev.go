package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	// ------------ retrieving environment variables for each replaceable field using variantdev

	// build a yaml file using kustomize from the dex folder.
	templateString := "./tmp/config-map.yaml"
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
	var configcmdStdoutList [configcmdStdoutListLen]string

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

		// turn *strings.Builder into string
		configcmdStdoutStr := configcmdStdout.String()

		// do regex processing to grab the actual value of the string output
		//check for the colon
		configcmdStdoutStrOut := strings.TrimSpace(configcmdStdoutStr[strings.IndexByte(configcmdStdoutStr, ' '):])

		// output the result standard output stored in a string
		logrus.Info("configcmdStdoutStrOut is: ", configcmdStdoutStrOut)
		logrus.Info(reflect.TypeOf(configcmdStdoutStrOut))
		logrus.Info(reflect.TypeOf(configcmd))

		configcmdStdoutList[i] = configcmdStdoutStrOut

	}

	logrus.Info("configcmdStdoutList: ", configcmdStdoutList)

	// ------------ replacing the string in the config.yaml file

	// create the sed command
	newGitLabAppId := configcmdStdoutList[0]
	newGitlabSecret := configcmdStdoutList[1]
	commandStrAppId := "s/$GITLAB_APPLICATION_ID/" + newGitLabAppId + "/g"
	commandStrSecret := "s/$GITLAB_CLIENT_SECRET/" + newGitlabSecret + "/g"

	// establish command to run sed on file.text
	// establish a strings buffer to hold the stdout
	cmd, cmdStdout := exec.Command("sed", "-e", commandStrAppId, "-e", commandStrSecret, "./tmp/config-template.yaml"), new(strings.Builder)
	// reroute stdout to the strings builder we established, cmdStdout
	cmd.Stdout = cmdStdout
	// run the command
	err := cmd.Run()
	if err != nil {
		logrus.Info(err)
	}
	// show the result. Note, type will be *strings.Builder accessible with cmdStdout.String()
	logrus.Info("cmdStdout is: ", cmdStdout)
	logrus.Info("cmdStdout type is: ", reflect.TypeOf(cmdStdout))
	// convert strings builder to string via print
	var cmdStdoutStr string
	cmdStdoutStr = cmdStdout.String()
	logrus.Info("cmdStdoutStr type: ", reflect.TypeOf(cmdStdoutStr))

	// write string to new yaml config file
	finalfile, err := os.Create("./tmp/config-map.yaml")
	if err != nil {
		logrus.Info(err)
	}

	// close file when done
	defer finalfile.Close()

	// write yamlStr, the actual contents of the yaml, to the tempFile, tmp/temp.yaml
	_, err = finalfile.WriteString(cmdStdoutStr)
	if err != nil {
		logrus.Info(err)
	}

}
