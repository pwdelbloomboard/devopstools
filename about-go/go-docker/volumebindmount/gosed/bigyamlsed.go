package main

import (
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	// create the sed command
	newGitLabAppId := "123456789_NEW"
	newGitlabSecret := "ABCDEFGHI_NEW"
	commandStrAppId := "s/$GITLAB_APPLICATION_ID/" + newGitLabAppId + "/g"
	commandStrSecret := "s/$GITLAB_CLIENT_SECRET/" + newGitlabSecret + "/g"

	// establish command to run sed on file.text
	// establish a strings buffer to hold the stdout
	cmd, cmdStdout := exec.Command("sed", "-e", commandStrAppId, "-e", commandStrSecret, "./tmp/bigyaml-template.yaml"), new(strings.Builder)
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
	finalfile, err := os.Create("tmp/bigyaml.yaml")
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
