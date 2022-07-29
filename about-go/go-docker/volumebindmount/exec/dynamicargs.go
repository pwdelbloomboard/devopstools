package main

import (
	"os/exec"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	// original file

	// setting up the arguments
	var args []string
	args = []string{"-e", "s/unix/linux/g", "-e", "s/free/superfree/g", "sedfile.txt"}
	logrus.Info(reflect.TypeOf(args))
	// creating the command
	cmd, cmdStdout := exec.Command("sed", args...), new(strings.Builder)

	logrus.Info("the cmd is: ", cmd)

	cmd.Stdout = cmdStdout
	// run the command
	err := cmd.Run()
	if err != nil {
		logrus.Info(err)
	}
	var cmdStdoutStr string
	cmdStdoutStr = cmdStdout.String()
	logrus.Info(cmdStdoutStr)

}

/*
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

*/
