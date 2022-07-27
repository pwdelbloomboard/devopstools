package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

/*
// Config level object, representing items under config:
type VarConfig struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
}
*/

// NestedConnector level object, representing items in array under "connectors"
type NestedConnector struct {
	Type   string `mapstructure:"type"`
	Id     string `mapstructure:"id"`
	Name   string `mapstructure:"name"`
	Config string `mapstructure:"config"`
}

func main() {

	// read the file
	yfile, err := ioutil.ReadFile("bigyaml.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	// make the map interface to hold the yaml
	dataMap := make(map[interface{}]interface{})

	// unmarshal yfile test.yaml into the map interface location
	err = yaml.Unmarshal(yfile, &dataMap)
	if err != nil {
		logrus.Fatal(err)
	}

	// access the multi-line string contained within the dataMap overall
	dataMapDataStr := fmt.Sprintf("%v", dataMap["data"].(map[interface{}]interface{})["config.yaml"])

	// write contents of string containing contents of interface{} to temporary file
	// create yaml file to restart unmarshalling process again.
	configStrFile, err := os.Create("./tmp/config.yaml")
	if err != nil {
		logrus.Info(err)
	}

	// defer closing the file, write string to config.yaml file created above
	defer configStrFile.Close()
	bytesAmount, err := configStrFile.WriteString(dataMapDataStr)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("bytesAmount written to dataMapDataStr: ", bytesAmount)

	// from specified directory
	// viper.AddConfigPath(currentdir())
	viper.AddConfigPath("./tmp")
	// load the newly created yaml config file into viper as an object
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		logrus.Fatal(err)
	}

	// drill down using type declarations to find the clientID
	clientID := viper.Get("connectors").([]interface{})[0].(map[string]interface{})["config"].(map[string]interface{})["clientID"]
	// configRegister := viper.Get("connectors").([]interface{})[0].(map[string]interface{})["config"].(map[string]interface{})
	logrus.Info("clientID is: ", clientID)
	logrus.Info("clientID is type: ", reflect.TypeOf(clientID))

	connectors := viper.Get("connectors.0.config")
	logrus.Info(connectors)

	// read the smaller, config file
	nfile, err := ioutil.ReadFile("./tmp/config.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	// make the map interface to hold the yaml
	miniConfig := make(map[interface{}]interface{})

	// unmarshal yfile *.yaml into the map interface location
	err = yaml.Unmarshal(nfile, &miniConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	// extract miniConfig as a string
	miniConfigStr := viper.GetString("connectors")
	// miniConfigStr := fmt.Sprintf("%v", miniConfig["connectors"].([]interface{})[0].(map[string]interface{})["config"])

	// write contents of string containing contents of interface{} to temporary file
	// create yaml file to restart unmarshalling process again.
	connectorsStrFile, err := os.Create("./tmp/connectors.yaml")
	if err != nil {
		logrus.Info(err)
	}

	// defer closing the file, write string to config.yaml file created above
	defer configStrFile.Close()
	bytesAmount, err = connectorsStrFile.WriteString(miniConfigStr)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("bytesAmount written to dataMapDataStr: ", bytesAmount)

}
