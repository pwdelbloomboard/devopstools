package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "admin",
		"password": "password",
		"host":     "localhost",
		"port":     3306,
		"database": "test",
	}
	configName  = "config"
	configPaths = []string{
		".",
	}
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func main() {
	for k, v := range defaults {
		// set all defaults without having to write out multiple lines of code.
		viper.SetDefault(k, v)
		// read out a file
		// set the configName as set above
		viper.SetConfigName(configName)
		for _, p := range configPaths {
			viper.AddConfigPath(p)
		}
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("could not read config file: %v", err)
		}
		// viper.SetConfigType("yaml")
		// set paths to look for configuration files
		// you can add multiple config paths in the order they are added
		// you could add $HOME/.appname, or other ways of adding paths.
		// viper.AddConfigPath(".")
	}
	// grab the variables from viper
	fmt.Printf("Hostname from viper: %s\n", viper.GetString("username"))
	fmt.Printf("Password from viper: %s\n", viper.GetString("password"))
	fmt.Printf("Host from viper: %s\n", viper.GetString("host"))
	fmt.Printf("Port from viper: %d\n", viper.GetInt("port"))
	fmt.Printf("Database from viper: %s\n", viper.GetString("database"))

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}
	// print from the struct / unmarshalled result
	fmt.Printf("Hostname from struct: %s\n", config.Username)
	fmt.Printf("Password from struct: %s\n", config.Password)
	fmt.Printf("Host from struct: %s\n", config.Host)
	fmt.Printf("Port from struct: %d\n", config.Port)
	fmt.Printf("Database from struct: %s\n", config.Database)

	// watch when the file changes.
	changed := false
	viper.WatchConfig()
	// this will update when the config file is changed.
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&config) // unmarshall the data back into our struct
		if err != nil {
			log.Printf("could not decode config after change: %v", err)
		}
		changed = true
	},
	)
	for !changed {
		time.Sleep(time.Second)
		fmt.Printf("Config struct: %v\n", config)
	}

}
