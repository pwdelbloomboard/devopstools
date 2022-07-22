package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func main() {

	// read the file
	yfile, err := ioutil.ReadFile("test.yaml")
	if err != nil {

		log.Fatal(err)
	}

	// make the map interface to hold the yaml
	data := make(map[interface{}]interface{})

	// unmarshal yfile test.yaml into the map interface location
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {

		log.Fatal(err2)
	}

	// range through data and display contents
	for k, v := range data {
		fmt.Printf("%s : %d\n", k, v)
	}

	// show type of data
	logrus.Info("TypeOf data: ", reflect.TypeOf(data))

	// marshal the extracted data,
	outByte, err3 := yaml.Marshal(data)
	if err3 != nil {
		logrus.Fatal(err3)
	}

	outString := string(outByte)
	logrus.Info("outString: ", outString)
	logrus.Info("outStringType: ", reflect.TypeOf(outString))

	// write string to yaml file
	f, err4 := os.Create("./tmp/dat2.yaml")
	if err4 != nil {
		logrus.Info(err4)
	}

	defer f.Close()

	// write string
	n3, err5 := f.WriteString(outString)
	if err5 != nil {
		logrus.Info(err5)
	}
	// success message
	logrus.Info("wrote ", n3, " bytes")

}
