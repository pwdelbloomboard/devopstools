package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

/*
// top level, yaml file representing entire struct, mirroring entire structure
type ConfigYaml struct {
	Issuer           string                     `yaml:"issuer"`
	Storage          string                     `yaml:"storage"`
	Web              string                     `yaml:"web"`
	Oauth2           string                     `yaml:"oauth2"`
	EnablePasswordDB bool                       `yaml:"enablePasswordDB"`
	Connectors       map[string]NestedConnector `yaml:"connectors"`
	StaticClients    string                     `yaml:"staticClients"`
}
*/

/*
// top level, yaml file representing entire struct, mirroring entire structure
type ConfigYaml struct {
	Issuer           string `mapstructure:"issuer"`
	Storage          string `yaml:"storage"`
	Web              string `yaml:"web"`
	Oauth2           string `yaml:"oauth2"`
	EnablePasswordDB bool   `yaml:"enablePasswordDB"`
	StaticClients    string `yaml:"staticClients"`
	// the "connectors" portion of the yaml file array
	Connectors struct {
		Type string `yaml:"type"`
		Id   string `yaml:"id"`
		Name string `yaml:"name"`
		// the "config" portion of the yaml file which holds the actual clientID and clientSecret
		Config struct {
			ClientID     string `yaml:"clientID"`
			ClientSecret string `yaml:"clientSecret"`
		} `yaml:"config"`
	} `yaml:"connectors"`
}
*/

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
		log.Fatal(err)
	}

	// Using Viper.Get Method to get connectors which includes an array of keys
	connectors := viper.Get("connectors")
	logrus.Info("viper.Get(connectors) result: ", connectors)
	logrus.Info("connectors type: ", reflect.TypeOf(connectors))

	// create a NestedConnector struct into which to unmarshal data from connectors with Viper
	// var typeidnameconfig NestedConnector
	var unmarshalinterface interface{}
	err = viper.UnmarshalKey("connectors", &unmarshalinterface)
	if err != nil {
		logrus.Info("unable to decode into struct, %v", err)
	}
	// now use viper get method on array of keys
	configClientID := viper.Get("config.clientID")
	logrus.Info("configClientID: ", configClientID)

	/*
		// Unmarshal viper data into newly created struct
		// this means finding whatever matches in struct C and populating it
		conf := &config{}
		conf, err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	*/

	/*
		// create a struct following our entire top level pattern above
		C := ConfigYaml{}

		// print empty struct
		logrus.Info("C := repo{} previous to update: ", C)

		// use viper to set value
		// viper.Set("data.config")

		// Change value in map and marshal back into yaml
		C.Connectors.Config.ClientID = "replacement_client_id"
		C.Connectors.Config.ClientSecret = "replacement_client_secret"

		logrus.Info("C := repo{} after update: ", C)

		// Marshal the data from repo C to byte[] d
		outputByte, err := yaml.Marshal(&C)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// print out the bytestring, stringified
		logrus.Info("string(d): ", string(outputByte))

		// create temporary output file
		f, err := os.Create("/tmp/dat")
		if err != nil {
			log.Fatal(err)
		}

		// name of file to write to
		writeFile := "changed.yaml"

		// write to file writeFile from outputByte
		// with FileMode https://pkg.go.dev/io/fs#FileMode
		// 0644 Unix Permission Bits
		err = ioutil.WriteFile(writeFile, outputByte, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// close the file

		f.Close()
	*/

}

// custom function, get the current directory
func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
