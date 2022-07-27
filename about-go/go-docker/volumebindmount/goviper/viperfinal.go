package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// from specified directory
	// viper.AddConfigPath(currentdir())
	viper.AddConfigPath("./")
	// load the yaml file as our viper object
	viper.SetConfigName("bigyaml-template")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Fatal(err)
	}

	// two step string conversion from interface to viper object
	dataConfig := viper.Get("data.config.yaml")
	dataConfigStr := fmt.Sprintf("%v", dataConfig)

	logrus.Info("dataConfigStr type: ", reflect.TypeOf(dataConfigStr))
	logrus.Info("dataConfigStr: ", dataConfigStr)

	// write the string to file as a yaml, to convert to json
	file, err := os.Create("./tmp/dataconfig.yaml")
	if err != nil {
		logrus.Info(err)
	}
	defer file.Close()
	file.WriteString(dataConfigStr)

	// convert file yaml to json to work better in viper yaml.YAMLToJSON
	filebytearray, err := ioutil.ReadFile("./tmp/dataconfig.yaml")
	jsonbytearray, err := yaml.YAMLToJSON(filebytearray)
	if err != nil {
		logrus.Info(err)
	}

	// set values with viper
	viper.SetConfigType("json")
	var jsondataconfigReader io.Reader
	jsondataconfigReader = strings.NewReader(string(jsonbytearray))
	err = viper.ReadConfig(jsondataconfigReader)
	if err != nil {
		logrus.Fatal(err)
	}

	// setting values that need to stay the same
	viper.Set("connectors.0.type", "gitlab")
	viper.Set("connectors.0.id", "gitlab")
	viper.Set("connectors.0.name", "Gitlab")
	viper.Set("connectors.0.config.redirectURI", "https://kf.ds.bloomboard.com/dex/callback")
	viper.Set("connectors.0.config.groups.0.bloomboard", "bloomboard")
	// setting values that need to change
	// step down conversion of viper object
	var newclientIDjson interface{}
	newclientIDjson = "whatever_clientid_new"
	var newclientSecretjson interface{}
	newclientSecretjson = "whatever_secret_new"
	viper.Set("connectors.0.config.clientID", newclientIDjson)
	viper.Set("connectors.0.config.clientSecret", newclientSecretjson)

	// write back to json file

	viper.AddConfigPath("./tmp/")
	viper.SetConfigName("jsonyamltest")
	viper.WriteConfig()

	// convert back to YAML with yaml.JSONToYAML, maintaining good formatting
	// write the string to file as a yaml, to convert to json

	// convert file yaml to json to work better in viper yaml.YAMLToJSON
	filebytearray, err = ioutil.ReadFile("./tmp/jsonyamltest.json")
	yamlbytearray, err := yaml.JSONToYAML(filebytearray)
	if err != nil {
		logrus.Info(err)
	}

	file, err = os.Create("./tmp/formattedyamltest.yaml")
	if err != nil {
		logrus.Info(err)
	}
	defer file.Close()
	file.WriteString(string(yamlbytearray))

	// alternate method - straight from Viper to YAML, write out with Viper

	// read in the sub-string as an io.Reader type
	var dataConfigReader io.Reader
	dataConfigReader = strings.NewReader(dataConfigStr)

	// get the string back into viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(dataConfigReader)
	if err != nil {
		logrus.Fatal(err)
	}

	// step down conversion of viper object
	var newclientID interface{}
	newclientID = "whatever_clientid_new"
	var newclientSecret interface{}
	newclientSecret = "whatever_secret_new"

	// set the value on the viper object
	// setting values that need to stay the same
	viper.Set("connectors.0.type", "gitlab")
	viper.Set("connectors.0.id", "gitlab")
	viper.Set("connectors.0.name", "Gitlab")
	viper.Set("connectors.0.config.redirectURI", "https://kf.ds.bloomboard.com/dex/callback")
	viper.Set("connectors.0.config.groups.0.bloomboard", "bloomboard")
	// setting values that need to change
	viper.Set("connectors.0.config.clientID", newclientID)
	viper.Set("connectors.0.config.clientSecret", newclientSecret)

	// write stuff
	viper.AddConfigPath("./tmp/")
	viper.SetConfigName("finaltest")
	viper.WriteConfig()

}
