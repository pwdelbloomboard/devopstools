package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func main() {

	// read the file
	yfile, err := ioutil.ReadFile("bigyaml.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	// get the type of yfile for confirmation
	logrus.Info("reflect.TypeOf(yfile): ", reflect.TypeOf(yfile))

	// make the map interface to hold the yaml
	dataMap := make(map[interface{}]interface{})

	// unmarshal yfile test.yaml into the map interface location
	err2 := yaml.Unmarshal(yfile, &dataMap)
	if err2 != nil {
		logrus.Fatal(err2)
	}

	// show type of dataMap
	logrus.Info("TypeOf dataMap: ", reflect.TypeOf(dataMap))

	// access one particular part of the map
	logrus.Info("---")
	logrus.Info("---")
	dataMapData := dataMap["data"]
	logrus.Info("dataMap[data]: ", dataMapData)

	// drill down into dataMapData using type assertion
	dataMapDataConfig := dataMapData.(map[string]interface{})["config.yaml"]
	logrus.Info("dataMapDataConfig: ", dataMapDataConfig)

	// drill down into dataMapDataConfig, after which contents are a string
	logrus.Info("type of dataMapDataConfig: ", reflect.TypeOf(dataMapDataConfig))
	logrus.Info("reflect.ValueOf(dataMapDataConfig).Kind(): ", reflect.ValueOf(dataMapDataConfig).Kind())

	// use fmt to simply print the value of the interface{} to a string
	dataMapDataConfigStr := fmt.Sprintf("%v", dataMapDataConfig)

	// write contents of string containing contents of interface{} to temporary file
	/*

		// create yaml file to restart unmarshalling process again.
		configStrFile, err := os.Create("./tmp/config.yaml")
		if err != nil {
			logrus.Info(err)
		}

			defer configStrFile.Close()
			whatever, err := configStrFile.WriteString(dataMapDataConfigStr)
			if err != nil {
				logrus.Info(err)
			}
			logrus.Info("whateever: ", whatever)
	*/

	// make the map interface to hold the config yaml
	configMap := make(map[interface{}]interface{})

	// unmarshal yfile test.yaml into the map interface location
	// after converting to []byte
	err = yaml.Unmarshal([]byte(dataMapDataConfigStr), &configMap)
	if err != nil {
		logrus.Fatal(err2)
	}

	// enter key accessed data from new configMap into configMapConnectors
	configMapConnectors := configMap["connectors"]
	logrus.Info("configMapConnectors := configMap[connectors]: ", configMapConnectors)
	logrus.Info("reflect.TypeOf(configMapConnectors): ", reflect.TypeOf(configMapConnectors))

	// dial down into next level, which is an array interface, []interface {}
	configMapConnectorsConfig := configMapConnectors.([]interface{})[0]
	logrus.Info("configMapConnectors.([]interface{})[0]: ", configMapConnectorsConfig)
	logrus.Info("configMapConnectorsConfig type: ", reflect.TypeOf(configMapConnectorsConfig))

	// we're back to map[string]interface so we can use type assertion again
	configMapConnectorsConfigClientId := configMapConnectorsConfig.(map[string]interface{})["clientID"]
	logrus.Info("configMapConnectorsConfigClientId: ", configMapConnectorsConfigClientId)

	//

	// marshal the extracted dataMap,
	outByte, err3 := yaml.Marshal(dataMapData)
	if err3 != nil {
		logrus.Fatal(err3)
	}

	outString := string(outByte)
	// logrus.Info("outString: ", outString)
	// logrus.Info("outStringType: ", reflect.TypeOf(outString))

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
