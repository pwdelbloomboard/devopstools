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
	logrus.Info("args: ",args)
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
