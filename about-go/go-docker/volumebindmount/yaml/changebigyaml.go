package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

type repo struct {
	Charts struct {
		ClientID     string `yaml:"clientID"`
		ClientSecret string `yaml:"clientSecret"`
	} `yaml:"Charts"`
}

func main() {

	// the problem with this method is that *only*
	// the exact values specified in the struct end up in the final yaml

	// load the yaml file
	viper.SetConfigName("bigyaml")
	// from specified directory
	viper.AddConfigPath(currentdir())
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal(err)
	}

	// create a main.repo struct C
	C := repo{}

	// Unmarshal viper data into new struct C
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// print
	logrus.Info("C := repo{} previous to update: ", C)

	// use viper to set value
	// viper.Set("data.config")

	// Change value in map and marshal back into yaml
	C.Charts.ClientID = "replacement_client_id"
	C.Charts.ClientSecret = "replacement_client_secret"

	logrus.Info("C := repo{} after update: ", C)

	// Marshal the data from repo C to byte[] d
	outputByte, err := yaml.Marshal(&C)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// print out the bytestring, stringified
	logrus.Info("string(d): ", string(outputByte))

	// create temporary output file
	f, err := os.Create("/tmp/dat2")
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

}

// custom function, get the current directory
func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
