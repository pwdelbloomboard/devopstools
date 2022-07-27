package main

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func main() {

	// establish command to run sed on file.text
	cmd := exec.Command("sed", "-e", "s/hello/goodbye/g", "./tmp/file.txt")
	// run the command
	resultStr, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Info(err)
	}
	// show the result
	logrus.Info(string(resultStr))

	// write the string to file as a yaml
	file, err := os.Create("./tmp/bigyaml.yaml")
	if err != nil {
		logrus.Info(err)
	}
	defer file.Close()
	file.WriteString(resultStr)

}
