package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	Cluster      string
	Port         int
	DatabaseName string
	UserName     string
	Password     string
	SslMode      string
}

func NewConfiguration(configFile string) *Configuration {
	// let's get some config going!
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(configFile)

	var config = &Configuration{}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to decode config file, %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file into configuration, %v", err)
	}
	return config
}
